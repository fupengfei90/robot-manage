package response

import "github.com/gin-gonic/gin"

// StandardResponse 统一响应格式。
type StandardResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// OK 返回成功响应。
func OK[T any](c *gin.Context, data T) {
	c.JSON(200, StandardResponse[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 返回失败响应。
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, StandardResponse[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
