package task

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/quocbang/distribute-video-processor/domain/task"
)

func (t *Task) ConvertHLSVideoTask(ctx context.Context, c task.ConvertHLSVideo) (*asynq.Task, error) {
	payload, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(task.TypeVideoConvertToHLS, payload), nil
}
