package repository

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/model"
)

type RBACRepository struct {
	db *gorm.DB
}

func NewRBACRepository(db *gorm.DB) *RBACRepository {
	return &RBACRepository{db: db}
}

// User相关方法

func (r *RBACRepository) GetUsers(query *model.WecomUserQuery) ([]*model.WecomUser, int64, error) {
	var users []*model.WecomUser
	var total int64

	db := r.db.Model(&model.WecomUser{}).Where("deleted_at IS NULL").Preload("Roles")

	if query.UserID != "" {
		db = db.Where("user_id LIKE ?", "%"+query.UserID+"%")
	}
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Active != nil {
		db = db.Where("active = ?", *query.Active)
	}
	if query.RoleID > 0 {
		db = db.Joins("JOIN wecom_user_roles ur ON wecom_users.id = ur.user_id").
			Where("ur.role_id = ? AND ur.deleted_at IS NULL", query.RoleID)
	}
	if query.StartTime != "" {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("created_at <= ?", query.EndTime)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}
	orderDir := strings.ToUpper(query.OrderDir)
	if orderDir != "ASC" && orderDir != "DESC" {
		orderDir = "DESC"
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", orderBy, orderDir)).
		Offset(offset).Limit(query.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *RBACRepository) GetUserByID(id uint) (*model.WecomUser, error) {
	var user model.WecomUser
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).
		Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RBACRepository) CreateUser(user *model.WecomUser) error {
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	return r.db.Create(user).Error
}

func (r *RBACRepository) UpdateUser(id uint, user *model.WecomUser) error {
	now := time.Now()
	user.UpdatedAt = &now
	return r.db.Where("id = ? AND deleted_at IS NULL", id).Updates(user).Error
}

func (r *RBACRepository) DeleteUser(id uint) error {
	now := time.Now()
	return r.db.Model(&model.WecomUser{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", now).Error
}

func (r *RBACRepository) UpdateUserStatus(id uint, active int) error {
	now := time.Now()
	return r.db.Model(&model.WecomUser{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"active":     active,
			"updated_at": now,
		}).Error
}

func (r *RBACRepository) BatchUpdateUserStatus(ids []uint, active int) error {
	now := time.Now()
	return r.db.Model(&model.WecomUser{}).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Updates(map[string]interface{}{
			"active":     active,
			"updated_at": now,
		}).Error
}

func (r *RBACRepository) BatchDeleteUsers(ids []uint) error {
	now := time.Now()
	return r.db.Model(&model.WecomUser{}).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Update("deleted_at", now).Error
}

func (r *RBACRepository) AssignUserRoles(userID uint, roleIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除现有角色关联
		if err := tx.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		
		// 添加新的角色关联
		for _, roleID := range roleIDs {
			userRole := &model.UserRole{
				UserID: userID,
				RoleID: roleID,
			}
			if err := tx.Create(userRole).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// Role相关方法

func (r *RBACRepository) GetRoles(query *model.RoleQuery) ([]*model.Role, int64, error) {
	var roles []*model.Role
	var total int64

	db := r.db.Model(&model.Role{}).Where("deleted_at IS NULL")

	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}
	orderDir := strings.ToUpper(query.OrderDir)
	if orderDir != "ASC" && orderDir != "DESC" {
		orderDir = "DESC"
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", orderBy, orderDir)).
		Offset(offset).Limit(query.PageSize).
		Preload("Users").Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *RBACRepository) GetRoleByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).
		Preload("Users").Preload("Permissions").First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RBACRepository) CreateRole(role *model.Role) error {
	now := time.Now()
	role.CreatedAt = &now
	role.UpdatedAt = &now
	return r.db.Create(role).Error
}

func (r *RBACRepository) UpdateRole(id uint, role *model.Role) error {
	now := time.Now()
	role.UpdatedAt = &now
	return r.db.Where("id = ? AND deleted_at IS NULL", id).Updates(role).Error
}

func (r *RBACRepository) DeleteRole(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		
		// 检查是否有用户使用该角色
		var count int64
		if err := tx.Model(&model.UserRole{}).
			Where("role_id = ? AND deleted_at IS NULL", id).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("角色正在被 %d 个用户使用，无法删除", count)
		}
		
		// 删除角色
		return tx.Model(&model.Role{}).
			Where("id = ? AND deleted_at IS NULL", id).
			Update("deleted_at", now).Error
	})
}

func (r *RBACRepository) AssignRolePermissions(roleID uint, permissionIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除现有权限关联
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}
		
		// 添加新的权限关联
		for _, permissionID := range permissionIDs {
			rolePermission := &model.RolePermission{
				RoleID:       roleID,
				PermissionID: permissionID,
			}
			if err := tx.Create(rolePermission).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// Permission相关方法

func (r *RBACRepository) GetPermissions(query *model.PermissionQuery) ([]*model.Permission, int64, error) {
	var permissions []*model.Permission
	var total int64

	db := r.db.Model(&model.Permission{}).Where("deleted_at IS NULL")

	if query.Intent != "" {
		db = db.Where("intent LIKE ?", "%"+query.Intent+"%")
	}
	if query.Intent2 != "" {
		db = db.Where("intent2 LIKE ?", "%"+query.Intent2+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order("intent ASC, intent2 ASC").
		Offset(offset).Limit(query.PageSize).Find(&permissions).Error; err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

func (r *RBACRepository) GetAllPermissions() ([]*model.Permission, error) {
	var permissions []*model.Permission
	if err := r.db.Where("deleted_at IS NULL").
		Order("intent ASC, intent2 ASC").Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RBACRepository) GetPermissionByID(id uint) (*model.Permission, error) {
	var permission model.Permission
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *RBACRepository) CreatePermission(permission *model.Permission) error {
	now := time.Now()
	permission.CreatedAt = &now
	permission.UpdatedAt = &now
	return r.db.Create(permission).Error
}

func (r *RBACRepository) UpdatePermission(id uint, permission *model.Permission) error {
	now := time.Now()
	permission.UpdatedAt = &now
	return r.db.Where("id = ? AND deleted_at IS NULL", id).Updates(permission).Error
}

func (r *RBACRepository) DeletePermission(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		
		// 检查是否有角色使用该权限
		var count int64
		if err := tx.Model(&model.RolePermission{}).
			Where("permission_id = ? AND deleted_at IS NULL", id).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("权限正在被 %d 个角色使用，无法删除", count)
		}
		
		// 删除权限
		return tx.Model(&model.Permission{}).
			Where("id = ? AND deleted_at IS NULL", id).
			Update("deleted_at", now).Error
	})
}

func (r *RBACRepository) CheckUserIDExists(userID string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&model.WecomUser{}).
		Where("user_id = ? AND deleted_at IS NULL", userID)
	
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *RBACRepository) CheckRoleNameExists(name string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&model.Role{}).
		Where("name = ? AND deleted_at IS NULL", name)
	
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}