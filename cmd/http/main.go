package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"

	"github.com/quocbang/distribute-video-processor/pkg/config"

	"github.com/quocbang/distribute-video-processor/adapters/infrastructures/cloud"
	"github.com/quocbang/distribute-video-processor/adapters/infrastructures/cloud/s3"
	"github.com/quocbang/distribute-video-processor/adapters/infrastructures/task"
	"github.com/quocbang/distribute-video-processor/adapters/interfaces/api"
	dTask "github.com/quocbang/distribute-video-processor/domain/task"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config, error: %v", err)
	}

	if err := registerLogger(true); err != nil {
		log.Fatal(err)
	}

	client, err := cloud.NewS3Client()
	if err != nil {
		log.Fatal(err)
	}

	s3 := s3.NewS3CloudService(client.Client, &s3.S3ConfigInfo{
		BucketName:       cfg.AWS.S3.BucketName,
		ShareEndPoint:    cfg.AWS.S3.ShareEndPoint,
		DefaultDirectory: cfg.AWS.S3.Directory,
	})

	taskClient := task.NewTaskClient(fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port), cfg.Redis.Password, 10, s3, cfg.AWS.S3.BucketName)

	a := api.API{
		Task: api.Task{
			Distributor: taskClient,
			Client:      taskClient.Client,
		},
		Cloud: api.Cloud{
			S3: s3,
		},
	}

	server, err := api.NewAPIServer(a)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// asynq worker pool
	go func() {
		zap.L().Info("start asynq pool")
		defer wg.Done()
		srv := asynq.NewServer(asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			Password: cfg.DB.Password,
			DB:       0,
			PoolSize: 20,
		}, asynq.Config{
			Concurrency: 20,
			Logger:      zap.S(),
			Queues: map[string]int{
				"critical": 12,
				"default":  6,
				"low":      2,
			},
		})

		mux := asynq.NewServeMux()
		mux.HandleFunc(dTask.TypeVideoConvertToHLS, taskClient.ConvertHLSVideoProcess)

		if err := srv.Run(mux); err != nil {
			log.Fatal(err)
		}
		zap.L().Info("asynq shutdown")
		os.Exit(1)
	}()

	go func() {
		defer wg.Done()
		zap.L().Info("start api server")
		// api server
		port := "8080"
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server.Handler); err != nil {
			zap.L().Sugar().Fatalf("failed to start api server, error: %v", err)
		}
		zap.L().Info("api server shutdown")
		os.Exit(1)
	}()

	wg.Wait()
}

func registerLogger(isDevMod bool) error {
	logger := &zap.Logger{}
	var err error
	if isDevMod {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			return err
		}
	}

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
	return nil
}
