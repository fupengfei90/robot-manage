package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

type RBACHandler struct {
	service *service.RBACService
}

func NewRBACHandler(service *service.RBACService) *RBACHandler {
	return &RBACHandler{service: service}
}

// User相关接口

func (h *RBACHandler) GetUsers(c *gin.Context) {
	var query model.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	users, total, err := h.service.GetUsers(&query)
	if err != nil {
		response.InternalServerError(c, "获取用户列表失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      users,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

func (h *RBACHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID", err.Error())
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		response.NotFound(c, "用户不存在", err.Error())
		return
	}

	response.OK(c, user)
}

func (h *RBACHandler) CreateUser(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		response.BadRequest(c, "创建用户失败", err.Error())
		return
	}

	response.Created(c, user)
}

func (h *RBACHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID", err.Error())
		return
	}

	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	user, err := h.service.UpdateUser(uint(id), &req)
	if err != nil {
		response.BadRequest(c, "更新用户失败", err.Error())
		return
	}

	response.OK(c, user)
}

func (h *RBACHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID", err.Error())
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		response.InternalServerError(c, "删除用户失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

func (h *RBACHandler) UpdateUserStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID", err.Error())
		return
	}

	var req struct {
		Active int `json:"active" binding:"oneof=0 1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.UpdateUserStatus(uint(id), req.Active); err != nil {
		response.InternalServerError(c, "更新状态失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "状态更新成功"})
}

func (h *RBACHandler) BatchOperateUsers(c *gin.Context) {
	var req model.BatchUserOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.BatchOperateUsers(&req); err != nil {
		response.BadRequest(c, "批量操作失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "批量操作成功"})
}

func (h *RBACHandler) AssignUserRoles(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID", err.Error())
		return
	}

	var req model.UserRoleAssignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.AssignUserRoles(uint(id), &req); err != nil {
		response.BadRequest(c, "分配角色失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "角色分配成功"})
}

// Role相关接口

func (h *RBACHandler) GetRoles(c *gin.Context) {
	var query model.RoleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	roles, total, err := h.service.GetRoles(&query)
	if err != nil {
		response.InternalServerError(c, "获取角色列表失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      roles,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

func (h *RBACHandler) GetRoleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的角色ID", err.Error())
		return
	}

	role, err := h.service.GetRoleByID(uint(id))
	if err != nil {
		response.NotFound(c, "角色不存在", err.Error())
		return
	}

	response.OK(c, role)
}

func (h *RBACHandler) CreateRole(c *gin.Context) {
	var req model.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	role, err := h.service.CreateRole(&req)
	if err != nil {
		response.BadRequest(c, "创建角色失败", err.Error())
		return
	}

	response.Created(c, role)
}

func (h *RBACHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的角色ID", err.Error())
		return
	}

	var req model.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	role, err := h.service.UpdateRole(uint(id), &req)
	if err != nil {
		response.BadRequest(c, "更新角色失败", err.Error())
		return
	}

	response.OK(c, role)
}

func (h *RBACHandler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的角色ID", err.Error())
		return
	}

	if err := h.service.DeleteRole(uint(id)); err != nil {
		response.BadRequest(c, "删除角色失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

func (h *RBACHandler) AssignRolePermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的角色ID", err.Error())
		return
	}

	var req model.RolePermissionAssignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.AssignRolePermissions(uint(id), &req); err != nil {
		response.BadRequest(c, "分配权限失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "权限分配成功"})
}

// Permission相关接口

func (h *RBACHandler) GetPermissions(c *gin.Context) {
	var query model.PermissionQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	permissions, total, err := h.service.GetPermissions(&query)
	if err != nil {
		response.InternalServerError(c, "获取权限列表失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      permissions,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

func (h *RBACHandler) GetPermissionTree(c *gin.Context) {
	tree, err := h.service.GetPermissionTree()
	if err != nil {
		response.InternalServerError(c, "获取权限树失败", err.Error())
		return
	}

	response.OK(c, tree)
}

func (h *RBACHandler) GetPermissionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的权限ID", err.Error())
		return
	}

	permission, err := h.service.GetPermissionByID(uint(id))
	if err != nil {
		response.NotFound(c, "权限不存在", err.Error())
		return
	}

	response.OK(c, permission)
}

func (h *RBACHandler) CreatePermission(c *gin.Context) {
	var req model.PermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	permission, err := h.service.CreatePermission(&req)
	if err != nil {
		response.BadRequest(c, "创建权限失败", err.Error())
		return
	}

	response.Created(c, permission)
}

func (h *RBACHandler) UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的权限ID", err.Error())
		return
	}

	var req model.PermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	permission, err := h.service.UpdatePermission(uint(id), &req)
	if err != nil {
		response.BadRequest(c, "更新权限失败", err.Error())
		return
	}

	response.OK(c, permission)
}

func (h *RBACHandler) DeletePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的权限ID", err.Error())
		return
	}

	if err := h.service.DeletePermission(uint(id)); err != nil {
		response.BadRequest(c, "删除权限失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}