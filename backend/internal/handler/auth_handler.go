package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

// AuthHandler 认证相关的HTTP处理器
type AuthHandler struct {
	svc *service.AuthService
}

// NewAuthHandler 构造函数
func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	token, user, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.OK(c, LoginResponse{
		Token: token,
		User:  user,
	})
}

// GetCurrentUser 获取当前用户信息
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "未找到用户信息")
		return
	}

	// user_id在JWT middleware中存储的是username（sub字段）
	user, err := h.svc.GetUserByID(userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, user)
}

// Logout 退出登录
func (h *AuthHandler) Logout(c *gin.Context) {
	// 这里可以添加token黑名单逻辑
	response.OK(c, gin.H{"message": "退出成功"})
}

