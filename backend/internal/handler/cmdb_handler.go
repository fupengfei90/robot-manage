package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

// CMDBHandler CMDB HTTP处理层
type CMDBHandler struct {
	svc *service.CMDBService
}

// NewCMDBHandler 创建CMDB Handler
func NewCMDBHandler(svc *service.CMDBService) *CMDBHandler {
	return &CMDBHandler{svc: svc}
}

// ========== 值班排班相关 ==========

// GetDutySchedules 获取值班排班列表
// @Summary 获取值班排班列表
// @Tags CMDB
// @Param startDate query string false "开始日期"
// @Param endDate query string false "结束日期"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(20)
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/duty/schedules [get]
func (h *CMDBHandler) GetDutySchedules(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	staffName := c.Query("staffName")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	schedules, total, err := h.svc.GetDutySchedules(startDate, endDate, staffName, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     schedules,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetDutyScheduleByID 根据ID获取值班排班
// @Summary 根据ID获取值班排班
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/duty/schedules/{id} [get]
func (h *CMDBHandler) GetDutyScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	schedule, err := h.svc.GetDutyScheduleByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, schedule)
}

// CreateDutySchedule 创建值班排班
// @Summary 创建值班排班
// @Tags CMDB
// @Param request body model.DutyScheduleCreateRequest true "创建请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/duty/schedules [post]
func (h *CMDBHandler) CreateDutySchedule(c *gin.Context) {
	var req model.DutyScheduleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	schedule, err := h.svc.CreateDutySchedule(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, schedule)
}

// UpdateDutySchedule 更新值班排班
// @Summary 更新值班排班
// @Tags CMDB
// @Param id path int true "ID"
// @Param request body model.DutyScheduleUpdateRequest true "更新请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/duty/schedules/{id} [put]
func (h *CMDBHandler) UpdateDutySchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.DutyScheduleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	schedule, err := h.svc.UpdateDutySchedule(id, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, schedule)
}

// DeleteDutySchedule 删除值班排班
// @Summary 删除值班排班
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/duty/schedules/{id} [delete]
func (h *CMDBHandler) DeleteDutySchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteDutySchedule(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// ========== 大事记相关 ==========

// GetMilestoneEvents 获取大事记列表
// @Summary 获取大事记列表
// @Tags CMDB
// @Param startDate query string false "开始日期"
// @Param endDate query string false "结束日期"
// @Param isActive query bool false "是否激活"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(20)
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/milestones [get]
func (h *CMDBHandler) GetMilestoneEvents(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	eventContent := c.Query("eventContent")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	var isActive *bool
	if isActiveStr := c.Query("isActive"); isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	events, total, err := h.svc.GetMilestoneEvents(startDate, endDate, eventContent, isActive, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     events,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetMilestoneEventByID 根据ID获取大事记
// @Summary 根据ID获取大事记
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/milestones/{id} [get]
func (h *CMDBHandler) GetMilestoneEventByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	event, err := h.svc.GetMilestoneEventByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, event)
}

// CreateMilestoneEvent 创建大事记
// @Summary 创建大事记
// @Tags CMDB
// @Param request body model.MilestoneEventCreateRequest true "创建请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/milestones [post]
func (h *CMDBHandler) CreateMilestoneEvent(c *gin.Context) {
	var req model.MilestoneEventCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	event, err := h.svc.CreateMilestoneEvent(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, event)
}

// UpdateMilestoneEvent 更新大事记
// @Summary 更新大事记
// @Tags CMDB
// @Param id path int true "ID"
// @Param request body model.MilestoneEventUpdateRequest true "更新请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/milestones/{id} [put]
func (h *CMDBHandler) UpdateMilestoneEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.MilestoneEventUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	event, err := h.svc.UpdateMilestoneEvent(id, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, event)
}

// DeleteMilestoneEvent 删除大事记
// @Summary 删除大事记
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/milestones/{id} [delete]
func (h *CMDBHandler) DeleteMilestoneEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteMilestoneEvent(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}
