package task

import (
	"github.com/hibiken/asynq"

	"github.com/quocbang/distribute-video-processor/domain/cloud/s3"
)

type Task struct {
	Client     *asynq.Client
	S3         s3.S3Cloud
	bucketName string
}

// NewTaskClient redisAddr is host:port
func NewTaskClient(redisAddr string, redisPass string, poolSize int, s3 s3.S3Cloud, bucketName string) *Task {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
		PoolSize: poolSize,
	})

	return &Task{
		Client:     client,
		S3:         s3,
		bucketName: bucketName,
	}
}
