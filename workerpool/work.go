package workerpool

import (
	"github.com/panjf2000/ants/v2"
)

func NewMultipleWorkerPool(workerSize int, sizePerPoll int, lbs ants.LoadBalancingStrategy, opts ...ants.Option) (*ants.MultiPool, error) {
	return ants.NewMultiPool(workerSize, sizePerPoll, lbs, opts...)
}
