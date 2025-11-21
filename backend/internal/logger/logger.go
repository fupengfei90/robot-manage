package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
	once         sync.Once
)

// Init 根据配置初始化全局日志器。
func Init(level, encoding string) (*zap.Logger, error) {
	var initErr error
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Encoding = encoding
		if encoding == "console" {
			cfg = zap.NewDevelopmentConfig()
		}

		if err := cfg.Level.UnmarshalText([]byte(level)); err != nil {
			initErr = err
			return
		}

		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		globalLogger, initErr = cfg.Build()
	})

	return globalLogger, initErr
}

// L 返回全局日志器，若未初始化则使用默认开发配置。
func L() *zap.Logger {
	if globalLogger == nil {
		l, _ := zap.NewDevelopment()
		return l
	}

	return globalLogger
}
