package task

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	ConvertHLSVideoTask(context.Context, ConvertHLSVideo) (*asynq.Task, error)
}

type TaskProcessor interface {
	ConvertHLSVideoProcess(context.Context, *asynq.Task) error
}
