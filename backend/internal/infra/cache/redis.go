package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/example/robot-manage/internal/config"
)

// New 初始化Redis客户端并做一次健康检查。
func New(cfg config.RedisSection) (*redis.Client, error) {
	if cfg.Addr == "" {
		return nil, fmt.Errorf("未配置Redis地址")
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		DB:   cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, err
	}

	return client, nil
}
