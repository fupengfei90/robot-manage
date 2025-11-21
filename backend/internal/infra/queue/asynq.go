package queue

import (
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/example/robot-manage/internal/config"
)

// NewClient 基于Redis的Asynq任务客户端。
func NewClient(cfg config.TaskQueueSection) (*asynq.Client, error) {
	if cfg.RedisDSN == "" {
		return nil, fmt.Errorf("未配置任务队列Redis DSN")
	}

	opt, err := asynq.ParseRedisURI(cfg.RedisDSN)
	if err != nil {
		return nil, err
	}

	return asynq.NewClient(opt), nil
}
