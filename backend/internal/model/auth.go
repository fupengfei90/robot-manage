package model

import "time"

// User 用户模型
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Password  string    `gorm:"not null;size:255" json:"-"` // 不返回密码
	Email     string    `gorm:"size:100" json:"email"`
	Role      string    `gorm:"default:'user';size:20" json:"role"`
	Avatar    string    `gorm:"size:255" json:"avatar"`
	Status    int       `gorm:"default:1" json:"status"` // 1: 启用, 0: 禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "manage_users"
}

