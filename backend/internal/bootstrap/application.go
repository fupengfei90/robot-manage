package bootstrap

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/config"
	"github.com/example/robot-manage/internal/handler"
	"github.com/example/robot-manage/internal/infra/cache"
	"github.com/example/robot-manage/internal/infra/database"
	"github.com/example/robot-manage/internal/infra/queue"
	"github.com/example/robot-manage/internal/infra/vector"
	"github.com/example/robot-manage/internal/logger"
	"github.com/example/robot-manage/internal/repository"
	"github.com/example/robot-manage/internal/server"
	"github.com/example/robot-manage/internal/service"
)

// Application 聚合所有运行时实例。
type Application struct {
	Config    *config.AppConfig
	Server    *server.Server
	DB        *gorm.DB
	Cache     redis.UniversalClient
	TaskQueue *asynq.Client
	Vector    vector.Store
}

// Build 初始化配置、日志、依赖并返回可运行的App。
func Build(configPath string) (*Application, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}

	if _, err := logger.Init(cfg.Logging.Level, cfg.Logging.Encoding); err != nil {
		return nil, fmt.Errorf("初始化日志失败: %w", err)
	}

	app := &Application{Config: cfg}

	if cfg.Database.DSN != "" {
		if db, err := database.New(cfg.Database); err != nil {
			logger.L().Warn("初始化数据库失败", zap.Error(err))
		} else {
			app.DB = db
		}
	}

	if cfg.Redis.Addr != "" {
		if client, err := cache.New(cfg.Redis); err != nil {
			logger.L().Warn("初始化Redis失败", zap.Error(err))
		} else {
			app.Cache = client
		}
	}

	if cfg.TaskQueue.RedisDSN != "" {
		if client, err := queue.NewClient(cfg.TaskQueue); err != nil {
			logger.L().Warn("初始化任务队列失败", zap.Error(err))
		} else {
			app.TaskQueue = client
		}
	}

	if store, err := vector.New(cfg.Vector); err != nil {
		logger.L().Warn("初始化向量数据库失败", zap.Error(err))
	} else {
		app.Vector = store
	}

	dashboardRepo := repository.NewDashboardRepository(app.DB)
	dashboardSvc := service.NewDashboardService(dashboardRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardSvc)

	cmdbRepo := repository.NewCMDBRepository(app.DB)
	cmdbSvc := service.NewCMDBService(cmdbRepo)
	cmdbHandler := handler.NewCMDBHandler(cmdbSvc)

	httpServer := server.New(cfg, dashboardHandler, cmdbHandler)

	app.Server = httpServer

	return app, nil
}

// Shutdown 统一释放资源。
func (a *Application) Shutdown(ctx context.Context) error {
	if a.Server != nil {
		if err := a.Server.Shutdown(ctx); err != nil {
			return err
		}
	}

	if a.Cache != nil {
		_ = a.Cache.Close()
	}

	if a.TaskQueue != nil {
		_ = a.TaskQueue.Close()
	}

	if a.Vector != nil {
		_ = a.Vector.Close(ctx)
	}

	return nil
}
