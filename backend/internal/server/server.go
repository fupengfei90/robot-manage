package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/example/robot-manage/internal/config"
	"github.com/example/robot-manage/internal/handler"
)

// Server 封装HTTP服务实例。
type Server struct {
	httpServer *http.Server
}

// New 根据配置创建Server。
func New(cfg *config.AppConfig, dashboardHandler *handler.DashboardHandler, cmdbHandler *handler.CMDBHandler) *Server {
	router := NewRouter(cfg, dashboardHandler, cmdbHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.App.HTTP.Host, cfg.App.HTTP.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{httpServer: srv}
}

// Start 启动HTTP服务。
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown 优雅停止。
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
