package api

import (
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/quocbang/distribute-video-processor/domain/cloud/s3"
	"github.com/quocbang/distribute-video-processor/domain/task"
)

type Task struct {
	Distributor task.TaskDistributor
	Client      *asynq.Client
}

type Cloud struct {
	S3 s3.S3Cloud
}

type API struct {
	Task  Task
	Cloud Cloud
}

func (a API) videoGroup(g *echo.Group) {
	a.NewVideoHandlerRoutes(g)
}

func NewAPIServer(a API) (*http.Server, error) {
	e := echo.New()

	// v1 group
	v1 := e.Group("/v1")

	// video
	videoGroup := v1.Group("/video")
	a.videoGroup(videoGroup)

	return e.Server, nil
}
