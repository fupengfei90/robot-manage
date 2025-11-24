package service

import (
	"fmt"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

type RBACService struct {
	repo *repository.RBACRepository
}

func NewRBACService(repo *repository.RBACRepository) *RBACService {
	return &RBACService{repo: repo}
}

// User相关方法

func (s *RBACService) GetUsers(query *model.WecomUserQuery) ([]*model.WecomUserResponse, int64, error) {
	users, total, err := s.repo.GetUsers(query)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.WecomUserResponse, len(users))
	for i, user := range users {
		roleNames := make([]string, len(user.Roles))
		for j, role := range user.Roles {
			roleNames[j] = role.Name
		}
		responses[i] = &model.WecomUserResponse{
			WecomUser: user,
			RoleNames: roleNames,
		}
	}

	return responses, total, nil
}

func (s *RBACService) GetUserByID(id uint) (*model.WecomUserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Name
	}

	return &model.WecomUserResponse{
		WecomUser: user,
		RoleNames: roleNames,
	}, nil
}

func (s *RBACService) CreateUser(req *model.WecomUserRequest) (*model.WecomUserResponse, error) {
	// 检查用户ID是否重复
	exists, err := s.repo.CheckUserIDExists(req.UserID, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("用户ID %s 已存在", req.UserID)
	}

	user := &model.WecomUser{
		UserID: req.UserID,
		Name:   req.Name,
		Active: req.Active,
		PID:    req.PID,
		Phone:  req.Phone,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	// 分配角色
	if len(req.RoleIDs) > 0 {
		if err := s.repo.AssignUserRoles(user.ID, req.RoleIDs); err != nil {
			return nil, err
		}
	}

	return s.GetUserByID(user.ID)
}

func (s *RBACService) UpdateUser(id uint, req *model.WecomUserRequest) (*model.WecomUserResponse, error) {
	// 检查用户ID是否重复
	exists, err := s.repo.CheckUserIDExists(req.UserID, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("用户ID %s 已存在", req.UserID)
	}

	user := &model.WecomUser{
		UserID: req.UserID,
		Name:   req.Name,
		Active: req.Active,
		PID:    req.PID,
		Phone:  req.Phone,
	}

	if err := s.repo.UpdateUser(id, user); err != nil {
		return nil, err
	}

	// 更新角色分配
	if err := s.repo.AssignUserRoles(id, req.RoleIDs); err != nil {
		return nil, err
	}

	return s.GetUserByID(id)
}

func (s *RBACService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *RBACService) UpdateUserStatus(id uint, active int) error {
	return s.repo.UpdateUserStatus(id, active)
}

func (s *RBACService) BatchOperateUsers(req *model.BatchWecomUserOperationRequest) error {
	switch req.Operation {
	case "enable":
		return s.repo.BatchUpdateUserStatus(req.UserIDs, 1)
	case "disable":
		return s.repo.BatchUpdateUserStatus(req.UserIDs, 0)
	case "delete":
		return s.repo.BatchDeleteUsers(req.UserIDs)
	default:
		return fmt.Errorf("不支持的操作: %s", req.Operation)
	}
}

func (s *RBACService) AssignUserRoles(userID uint, req *model.WecomUserRoleAssignRequest) error {
	return s.repo.AssignUserRoles(userID, req.RoleIDs)
}

// Role相关方法

func (s *RBACService) GetRoles(query *model.RoleQuery) ([]*model.RoleResponse, int64, error) {
	roles, total, err := s.repo.GetRoles(query)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.RoleResponse, len(roles))
	for i, role := range roles {
		permissionNames := make([]string, len(role.Permissions))
		for j, permission := range role.Permissions {
			if permission.Intent2 != "" {
				permissionNames[j] = fmt.Sprintf("%s:%s", permission.Intent, permission.Intent2)
			} else {
				permissionNames[j] = permission.Intent
			}
		}
		responses[i] = &model.RoleResponse{
			Role:            role,
			UserCount:       len(role.Users),
			PermissionCount: len(role.Permissions),
			PermissionNames: permissionNames,
		}
	}

	return responses, total, nil
}

func (s *RBACService) GetRoleByID(id uint) (*model.RoleResponse, error) {
	role, err := s.repo.GetRoleByID(id)
	if err != nil {
		return nil, err
	}

	permissionNames := make([]string, len(role.Permissions))
	for i, permission := range role.Permissions {
		if permission.Intent2 != "" {
			permissionNames[i] = fmt.Sprintf("%s:%s", permission.Intent, permission.Intent2)
		} else {
			permissionNames[i] = permission.Intent
		}
	}

	return &model.RoleResponse{
		Role:            role,
		UserCount:       len(role.Users),
		PermissionCount: len(role.Permissions),
		PermissionNames: permissionNames,
	}, nil
}

func (s *RBACService) CreateRole(req *model.RoleRequest) (*model.RoleResponse, error) {
	// 检查角色名称是否重复
	exists, err := s.repo.CheckRoleNameExists(req.Name, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("角色名称 %s 已存在", req.Name)
	}

	role := &model.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.CreateRole(role); err != nil {
		return nil, err
	}

	// 分配权限
	if len(req.PermissionIDs) > 0 {
		if err := s.repo.AssignRolePermissions(role.ID, req.PermissionIDs); err != nil {
			return nil, err
		}
	}

	return s.GetRoleByID(role.ID)
}

func (s *RBACService) UpdateRole(id uint, req *model.RoleRequest) (*model.RoleResponse, error) {
	// 检查角色名称是否重复
	exists, err := s.repo.CheckRoleNameExists(req.Name, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("角色名称 %s 已存在", req.Name)
	}

	role := &model.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.UpdateRole(id, role); err != nil {
		return nil, err
	}

	// 更新权限分配
	if err := s.repo.AssignRolePermissions(id, req.PermissionIDs); err != nil {
		return nil, err
	}

	return s.GetRoleByID(id)
}

func (s *RBACService) DeleteRole(id uint) error {
	return s.repo.DeleteRole(id)
}

func (s *RBACService) AssignRolePermissions(roleID uint, req *model.RolePermissionAssignRequest) error {
	return s.repo.AssignRolePermissions(roleID, req.PermissionIDs)
}

// Permission相关方法

func (s *RBACService) GetPermissions(query *model.PermissionQuery) ([]*model.Permission, int64, error) {
	return s.repo.GetPermissions(query)
}

func (s *RBACService) GetPermissionTree() ([]*model.PermissionTreeNode, error) {
	permissions, err := s.repo.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	// 构建权限树
	intentMap := make(map[string]*model.PermissionTreeNode)
	
	for _, permission := range permissions {
		// 确保一级意图节点存在
		if _, exists := intentMap[permission.Intent]; !exists {
			intentMap[permission.Intent] = &model.PermissionTreeNode{
				Intent:      permission.Intent,
				Description: fmt.Sprintf("%s模块", permission.Intent),
				Children:    make([]*model.PermissionTreeNode, 0),
			}
		}
		
		// 如果有二级意图，添加为子节点
		if permission.Intent2 != "" {
			child := &model.PermissionTreeNode{
				ID:          permission.ID,
				Intent:      permission.Intent,
				Intent2:     permission.Intent2,
				Description: permission.Description,
				Children:    nil,
			}
			intentMap[permission.Intent].Children = append(intentMap[permission.Intent].Children, child)
		} else {
			// 如果没有二级意图，更新一级意图的ID和描述
			intentMap[permission.Intent].ID = permission.ID
			if permission.Description != "" {
				intentMap[permission.Intent].Description = permission.Description
			}
		}
	}

	// 转换为数组
	result := make([]*model.PermissionTreeNode, 0, len(intentMap))
	for _, node := range intentMap {
		result = append(result, node)
	}

	return result, nil
}

func (s *RBACService) GetPermissionByID(id uint) (*model.Permission, error) {
	return s.repo.GetPermissionByID(id)
}

func (s *RBACService) CreatePermission(req *model.PermissionRequest) (*model.Permission, error) {
	permission := &model.Permission{
		Intent:      req.Intent,
		Intent2:     req.Intent2,
		Description: req.Description,
	}

	if err := s.repo.CreatePermission(permission); err != nil {
		return nil, err
	}

	return s.GetPermissionByID(permission.ID)
}

func (s *RBACService) UpdatePermission(id uint, req *model.PermissionRequest) (*model.Permission, error) {
	permission := &model.Permission{
		Intent:      req.Intent,
		Intent2:     req.Intent2,
		Description: req.Description,
	}

	if err := s.repo.UpdatePermission(id, permission); err != nil {
		return nil, err
	}

	return s.GetPermissionByID(id)
}

func (s *RBACService) DeletePermission(id uint) error {
	return s.repo.DeletePermission(id)
}