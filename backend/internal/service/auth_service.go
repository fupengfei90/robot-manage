package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"go.uber.org/zap"

	"github.com/example/robot-manage/internal/config"
	"github.com/example/robot-manage/internal/logger"
	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

// AuthService 认证服务
type AuthService struct {
	repo   *repository.AuthRepository
	config *config.AppConfig
}

// NewAuthService 构造函数
func NewAuthService(repo *repository.AuthRepository, cfg *config.AppConfig) *AuthService {
	return &AuthService{
		repo:   repo,
		config: cfg,
	}
}

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, *model.User, error) {
	// 查找用户
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		logger.L().Warn("登录失败: 用户不存在", zap.String("username", username), zap.Error(err))
		return "", nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.L().Warn("登录失败: 密码错误", zap.String("username", username), zap.Error(err))
		return "", nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return "", nil, errors.New("用户已被禁用")
	}

	// 生成JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return "", nil, errors.New("生成token失败")
	}

	// 清除密码字段
	user.Password = ""

	return token, user, nil
}

// GetUserByID 根据ID获取用户
func (s *AuthService) GetUserByID(userID interface{}) (*model.User, error) {
	var id uint
	switch v := userID.(type) {
	case string:
		// 如果userID是字符串，尝试从JWT claims中获取username，然后查找用户
		user, err := s.repo.FindByUsername(v)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		return user, nil
	case uint:
		id = v
	default:
		return nil, errors.New("无效的用户ID类型")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

// generateToken 生成JWT token
func (s *AuthService) generateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.Username,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天过期
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Security.JWTSecret))
}

