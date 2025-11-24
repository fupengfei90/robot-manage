package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/model"
)

// AuthRepository 认证相关的数据访问层
type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepository 构造函数
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// FindByUsername 根据用户名查找用户
func (r *AuthRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找用户
func (r *AuthRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (r *AuthRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

