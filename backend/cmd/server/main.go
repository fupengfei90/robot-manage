package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/robot-manage/internal/bootstrap"
)

func main() {
	configPath := os.Getenv("ROBOT_MANAGE_CONFIG")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	app, err := bootstrap.Build(configPath)
	if err != nil {
		log.Fatalf("启动失败: %v", err)
	}

	go func() {
		if err := app.Server.Start(); err != nil {
			log.Printf("HTTP服务结束: %v", err)
		}
	}()

	log.Printf("robot-manage 启动成功，监听 %s:%d", app.Config.App.HTTP.Host, app.Config.App.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Printf("优雅停止失败: %v", err)
	}
}
