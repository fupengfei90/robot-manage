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

// BadRequest 返回400错误响应。
func BadRequest(c *gin.Context, message string, detail string) {
	c.JSON(400, StandardResponse[any]{
		Code:    400,
		Message: message + ": " + detail,
		Data:    nil,
	})
}

// NotFound 返回404错误响应。
func NotFound(c *gin.Context, message string, detail string) {
	c.JSON(404, StandardResponse[any]{
		Code:    404,
		Message: message + ": " + detail,
		Data:    nil,
	})
}

// InternalServerError 返回500错误响应。
func InternalServerError(c *gin.Context, message string, detail string) {
	c.JSON(500, StandardResponse[any]{
		Code:    500,
		Message: message + ": " + detail,
		Data:    nil,
	})
}

// Created 返回201创建成功响应。
func Created[T any](c *gin.Context, data T) {
	c.JSON(201, StandardResponse[T]{
		Code:    0,
		Message: "created",
		Data:    data,
	})
}
