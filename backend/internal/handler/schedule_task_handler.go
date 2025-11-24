package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

type ScheduleTaskHandler struct {
	service *service.ScheduleTaskService
}

func NewScheduleTaskHandler(service *service.ScheduleTaskService) *ScheduleTaskHandler {
	return &ScheduleTaskHandler{service: service}
}

// GetList 获取任务列表
func (h *ScheduleTaskHandler) GetList(c *gin.Context) {
	var query model.ScheduleTaskQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	// 设置默认值
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	tasks, total, err := h.service.GetList(&query)
	if err != nil {
		response.InternalServerError(c, "获取任务列表失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      tasks,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

// GetByID 根据ID获取任务
func (h *ScheduleTaskHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID", err.Error())
		return
	}

	task, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "任务不存在", err.Error())
		return
	}

	response.OK(c, task)
}

// Create 创建任务
func (h *ScheduleTaskHandler) Create(c *gin.Context) {
	var req model.ScheduleTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	task, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, "创建任务失败", err.Error())
		return
	}

	response.Created(c, task)
}

// Update 更新任务
func (h *ScheduleTaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID", err.Error())
		return
	}

	var req model.ScheduleTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	task, err := h.service.Update(uint(id), &req)
	if err != nil {
		response.BadRequest(c, "更新任务失败", err.Error())
		return
	}

	response.OK(c, task)
}

// Delete 删除任务
func (h *ScheduleTaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID", err.Error())
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.InternalServerError(c, "删除任务失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// UpdateStatus 更新任务状态
func (h *ScheduleTaskHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID", err.Error())
		return
	}

	var req struct {
		Active int `json:"active" binding:"oneof=0 1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.UpdateStatus(uint(id), req.Active); err != nil {
		response.InternalServerError(c, "更新状态失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "状态更新成功"})
}

// BatchOperation 批量操作
func (h *ScheduleTaskHandler) BatchOperation(c *gin.Context) {
	var req model.BatchOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.BatchOperation(&req); err != nil {
		response.BadRequest(c, "批量操作失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "批量操作成功"})
}

// ExecuteTask 立即执行任务
func (h *ScheduleTaskHandler) ExecuteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的任务ID", err.Error())
		return
	}

	if err := h.service.ExecuteTask(uint(id)); err != nil {
		response.BadRequest(c, "执行任务失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "任务执行成功"})
}

// GetCategories 获取所有分类
func (h *ScheduleTaskHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		response.InternalServerError(c, "获取分类失败", err.Error())
		return
	}

	// 添加预设分类
	presetCategories := []string{"rota", "weekly-report", "version-align"}
	categoryMap := make(map[string]bool)
	
	// 合并预设分类和数据库中的分类
	allCategories := make([]string, 0)
	for _, cat := range presetCategories {
		allCategories = append(allCategories, cat)
		categoryMap[cat] = true
	}
	
	for _, cat := range categories {
		if !categoryMap[cat] {
			allCategories = append(allCategories, cat)
		}
	}

	response.OK(c, allCategories)
}

// GetDCNs 获取所有数据中心
func (h *ScheduleTaskHandler) GetDCNs(c *gin.Context) {
	dcns, err := h.service.GetDCNs()
	if err != nil {
		response.InternalServerError(c, "获取数据中心失败", err.Error())
		return
	}

	// 添加预设数据中心
	presetDCNs := []string{"6A0"}
	dcnMap := make(map[string]bool)
	
	// 合并预设数据中心和数据库中的数据中心
	allDCNs := make([]string, 0)
	for _, dcn := range presetDCNs {
		allDCNs = append(allDCNs, dcn)
		dcnMap[dcn] = true
	}
	
	for _, dcn := range dcns {
		if !dcnMap[dcn] {
			allDCNs = append(allDCNs, dcn)
		}
	}

	response.OK(c, allDCNs)
}