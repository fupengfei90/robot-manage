package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/example/robot-manage/internal/logger"
)

// RequestID 中间件注入全链路追踪ID。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := uuid.NewString()
		c.Set("request_id", rid)
		c.Writer.Header().Set("X-Request-ID", rid)
		c.Next()
	}
}

// AccessLog 输出访问日志。
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		logger.L().Info("access",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("request_id", c.GetString("request_id")),
		)
	}
}
