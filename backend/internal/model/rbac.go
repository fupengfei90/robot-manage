package model

import (
	"time"
)

// WecomUser RBAC用户模型 (基于wecom_users表)
type WecomUser struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string     `gorm:"not null;size:100;uniqueIndex;comment:用户ID" json:"user_id"`
	Name      string     `gorm:"not null;size:100;comment:用户姓名" json:"name"`
	CreatedAt *time.Time `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt *time.Time `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt *time.Time `gorm:"comment:删除时间" json:"deleted_at"`
	Active    int        `gorm:"default:0;comment:是否可用" json:"active"`
	PID       string     `gorm:"size:100;comment:P账号" json:"p_id"`
	Phone     string     `gorm:"size:20;comment:手机号" json:"phone"`
	Roles     []Role     `gorm:"many2many:wecom_user_roles;foreignKey:ID;joinForeignKey:user_id;References:ID;joinReferences:role_id" json:"roles"` 
}

func (WecomUser) TableName() string {
	return "wecom_users"
}

// Role 角色模型 (基于wecom_roles表)
type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"not null;size:100;uniqueIndex;comment:角色名称" json:"name"`
	Description string       `gorm:"type:text;comment:角色描述" json:"description"`
	CreatedAt   *time.Time   `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   *time.Time   `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt   *time.Time   `gorm:"comment:删除时间" json:"deleted_at"`
	Users       []WecomUser  `gorm:"many2many:wecom_user_roles;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:user_id" json:"users"`
	Permissions []Permission `gorm:"many2many:wecom_role_permissions;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:permission_id" json:"permissions"`
}

func (Role) TableName() string {
	return "wecom_roles"
}

// Permission 权限模型 (基于wecom_permissions表)
type Permission struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Intent      string     `gorm:"not null;size:100;comment:一级意图" json:"intent"`
	Intent2     string     `gorm:"size:100;comment:二级意图" json:"intent2"`
	Description string     `gorm:"type:text;comment:权限描述" json:"description"`
	CreatedAt   *time.Time `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"comment:删除时间" json:"deleted_at"`
	Roles       []Role     `gorm:"many2many:wecom_role_permissions;foreignKey:ID;joinForeignKey:permission_id;References:ID;joinReferences:role_id" json:"roles"`
}

func (Permission) TableName() string {
	return "wecom_permissions"
}

// UserRole 用户角色关联表
type UserRole struct {
	UserID    uint       `gorm:"primaryKey" json:"user_id"`
	RoleID    uint       `gorm:"primaryKey" json:"role_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (UserRole) TableName() string {
	return "wecom_user_roles"
}

// RolePermission 角色权限关联表
type RolePermission struct {
	RoleID       uint       `gorm:"primaryKey" json:"role_id"`
	PermissionID uint       `gorm:"primaryKey" json:"permission_id"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (RolePermission) TableName() string {
	return "wecom_role_permissions"
}

// 请求和响应模型

// WecomUserRequest 用户请求
type WecomUserRequest struct {
	UserID  string `json:"user_id" binding:"required,max=100"`
	Name    string `json:"name" binding:"required,max=100"`
	Active  int    `json:"active"`
	PID     string `json:"p_id" binding:"max=100"`
	Phone   string `json:"phone" binding:"max=20"`
	RoleIDs []uint `json:"role_ids"`
}

// WecomUserQuery 用户查询参数
type WecomUserQuery struct {
	UserID    string `form:"user_id"`
	Name      string `form:"name"`
	Active    *int   `form:"active"`
	RoleID    uint   `form:"role_id"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	Page      int    `form:"page,default=1"`
	PageSize  int    `form:"page_size,default=10"`
	OrderBy   string `form:"order_by,default=id"`
	OrderDir  string `form:"order_dir,default=desc"`
}

// RoleRequest 角色请求
type RoleRequest struct {
	Name          string `json:"name" binding:"required,max=100"`
	Description   string `json:"description"`
	PermissionIDs []uint `json:"permission_ids"`
}

// RoleQuery 角色查询参数
type RoleQuery struct {
	Name     string `form:"name"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=10"`
	OrderBy  string `form:"order_by,default=id"`
	OrderDir string `form:"order_dir,default=desc"`
}

// PermissionRequest 权限请求
type PermissionRequest struct {
	Intent      string `json:"intent" binding:"required,max=100"`
	Intent2     string `json:"intent2" binding:"max=100"`
	Description string `json:"description"`
}

// PermissionQuery 权限查询参数
type PermissionQuery struct {
	Intent   string `form:"intent"`
	Intent2  string `form:"intent2"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=10"`
}

// 响应模型
type WecomUserResponse struct {
	*WecomUser
	RoleNames []string `json:"role_names"`
}

type RoleResponse struct {
	*Role
	UserCount       int    `json:"user_count"`
	PermissionCount int    `json:"permission_count"`
	PermissionNames []string `json:"permission_names"`
}

type PermissionTreeNode struct {
	Intent      string                `json:"intent"`
	Description string                `json:"description"`
	Children    []*PermissionTreeNode `json:"children"`
	ID          uint                  `json:"id,omitempty"`
	Intent2     string                `json:"intent2,omitempty"`
}

// BatchWecomUserOperationRequest 批量用户操作请求
type BatchWecomUserOperationRequest struct {
	UserIDs   []uint `json:"user_ids" binding:"required"`
	Operation string `json:"operation" binding:"required,oneof=enable disable delete"`
}

// WecomUserRoleAssignRequest 用户角色分配请求
type WecomUserRoleAssignRequest struct {
	RoleIDs []uint `json:"role_ids" binding:"required"`
}

// RolePermissionAssignRequest 角色权限分配请求
type RolePermissionAssignRequest struct {
	PermissionIDs []uint `json:"permission_ids" binding:"required"`
}

// Type aliases for handler compatibility
type UserQuery = WecomUserQuery
type UserRequest = WecomUserRequest
type BatchUserOperationRequest = BatchWecomUserOperationRequest
type UserRoleAssignRequest = WecomUserRoleAssignRequest