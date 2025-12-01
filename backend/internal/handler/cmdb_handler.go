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

// ========== 批量时间相关 ==========

// GetBatchTimes 获取批量时间列表
// @Summary 获取批量时间列表
// @Tags CMDB
// @Param systemName query string false "系统名称"
// @Param subsysName query string false "子系统名称"
// @Param batchTime query string false "批量时间"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(20)
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/batch-times [get]
func (h *CMDBHandler) GetBatchTimes(c *gin.Context) {
	systemName := c.Query("systemName")
	subsysName := c.Query("subsysName")
	batchTime := c.Query("batchTime")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	batchTimes, total, err := h.svc.GetBatchTimes(systemName, subsysName, batchTime, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     batchTimes,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetBatchTimeByID 根据ID获取批量时间
// @Summary 根据ID获取批量时间
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/batch-times/{id} [get]
func (h *CMDBHandler) GetBatchTimeByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	batchTime, err := h.svc.GetBatchTimeByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, batchTime)
}

// CreateBatchTime 创建批量时间
// @Summary 创建批量时间
// @Tags CMDB
// @Param request body model.BatchTimeCreateRequest true "创建请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/batch-times [post]
func (h *CMDBHandler) CreateBatchTime(c *gin.Context) {
	var req model.BatchTimeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	batchTime, err := h.svc.CreateBatchTime(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, batchTime)
}

// UpdateBatchTime 更新批量时间
// @Summary 更新批量时间
// @Tags CMDB
// @Param id path int true "ID"
// @Param request body model.BatchTimeUpdateRequest true "更新请求"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/batch-times/{id} [put]
func (h *CMDBHandler) UpdateBatchTime(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.BatchTimeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	batchTime, err := h.svc.UpdateBatchTime(uint(id), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, batchTime)
}

// DeleteBatchTime 删除批量时间
// @Summary 删除批量时间
// @Tags CMDB
// @Param id path int true "ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/batch-times/{id} [delete]
func (h *CMDBHandler) DeleteBatchTime(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteBatchTime(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// GetSystems 获取系统列表
// @Summary 获取系统列表
// @Tags CMDB
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/systems [get]
func (h *CMDBHandler) GetSystems(c *gin.Context) {
	systems, err := h.svc.GetSystems()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, systems)
}

// GetSubsystemsBySystemID 根据系统ID获取子系统列表
// @Summary 根据系统ID获取子系统列表
// @Tags CMDB
// @Param id path int true "系统ID"
// @Success 200 {object} response.StandardResponse
// @Router /api/v1/info/systems/{id}/subsystems [get]
func (h *CMDBHandler) GetSubsystemsBySystemID(c *gin.Context) {
	systemID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的系统ID")
		return
	}

	subsystems, err := h.svc.GetSubsystemsBySystemID(uint(systemID))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, subsystems)
}

// ========== WB CMDB相关 ==========

func (h *CMDBHandler) GetWBCMDBList(c *gin.Context) {
	systemName := c.Query("systemName")
	environment := c.Query("environment")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	list, total, err := h.svc.GetWBCMDBList(systemName, environment, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (h *CMDBHandler) GetWBCMDBByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	info, err := h.svc.GetWBCMDBByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) CreateWBCMDB(c *gin.Context) {
	var req model.WBCMDBInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.svc.CreateWBCMDB(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) UpdateWBCMDB(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.WBCMDBInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.svc.UpdateWBCMDB(uint(id), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) DeleteWBCMDB(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteWBCMDB(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// ========== VB CMDB相关 ==========

func (h *CMDBHandler) GetVBCMDBList(c *gin.Context) {
	systemName := c.Query("systemName")
	environment := c.Query("environment")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	list, total, err := h.svc.GetVBCMDBList(systemName, environment, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (h *CMDBHandler) GetVBCMDBByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	info, err := h.svc.GetVBCMDBByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) CreateVBCMDB(c *gin.Context) {
	var req model.VBCMDBInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.svc.CreateVBCMDB(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) UpdateVBCMDB(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.VBCMDBInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.svc.UpdateVBCMDB(uint(id), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, info)
}

func (h *CMDBHandler) DeleteVBCMDB(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteVBCMDB(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// ========== ITSM出包记录相关 ==========

func (h *CMDBHandler) GetITSMPackageRecords(c *gin.Context) {
	subsystem := c.Query("subsystem")
	status := c.Query("status")
	environment := c.Query("environment")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	records, total, err := h.svc.GetITSMPackageRecords(subsystem, status, environment, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":     records,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (h *CMDBHandler) GetITSMPackageRecordByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	record, err := h.svc.GetITSMPackageRecordByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.OK(c, record)
}

func (h *CMDBHandler) CreateITSMPackageRecord(c *gin.Context) {
	var req model.ITSMPackageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	record, err := h.svc.CreateITSMPackageRecord(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, record)
}

func (h *CMDBHandler) UpdateITSMPackageRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.ITSMPackageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	record, err := h.svc.UpdateITSMPackageRecord(uint(id), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, record)
}

func (h *CMDBHandler) DeleteITSMPackageRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteITSMPackageRecord(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}
