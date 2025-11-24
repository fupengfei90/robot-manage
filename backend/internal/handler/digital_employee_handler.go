package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

type DigitalEmployeeHandler struct {
	service *service.DigitalEmployeeService
}

func NewDigitalEmployeeHandler(service *service.DigitalEmployeeService) *DigitalEmployeeHandler {
	return &DigitalEmployeeHandler{service: service}
}

// ========== 历史会话记录相关 ==========

// GetMessageRecords 获取会话记录列表
func (h *DigitalEmployeeHandler) GetMessageRecords(c *gin.Context) {
	var query model.MessageRecordQuery
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

	records, total, err := h.service.GetMessageRecords(&query)
	if err != nil {
		response.InternalServerError(c, "获取会话记录失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      records,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

// GetMessageRecordByID 根据ID获取会话记录
func (h *DigitalEmployeeHandler) GetMessageRecordByID(c *gin.Context) {
	msgID := c.Param("id")
	if msgID == "" {
		response.BadRequest(c, "消息ID不能为空", "")
		return
	}

	record, err := h.service.GetMessageRecordByID(msgID)
	if err != nil {
		response.NotFound(c, "会话记录不存在", err.Error())
		return
	}

	response.OK(c, record)
}

// GetMessagesByConversationID 根据会话ID获取所有消息
func (h *DigitalEmployeeHandler) GetMessagesByConversationID(c *gin.Context) {
	conversationID := c.Param("conversation_id")
	if conversationID == "" {
		response.BadRequest(c, "会话ID不能为空", "")
		return
	}

	messages, err := h.service.GetMessagesByConversationID(conversationID)
	if err != nil {
		response.InternalServerError(c, "获取会话消息失败", err.Error())
		return
	}

	response.OK(c, messages)
}

// CreateMessageRecord 创建会话记录
func (h *DigitalEmployeeHandler) CreateMessageRecord(c *gin.Context) {
	var req model.MessageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	record, err := h.service.CreateMessageRecord(&req)
	if err != nil {
		response.BadRequest(c, "创建会话记录失败", err.Error())
		return
	}

	response.Created(c, record)
}

// UpdateMessageRecord 更新会话记录
func (h *DigitalEmployeeHandler) UpdateMessageRecord(c *gin.Context) {
	msgID := c.Param("id")
	if msgID == "" {
		response.BadRequest(c, "消息ID不能为空", "")
		return
	}

	var req model.MessageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	record, err := h.service.UpdateMessageRecord(msgID, &req)
	if err != nil {
		response.BadRequest(c, "更新会话记录失败", err.Error())
		return
	}

	response.OK(c, record)
}

// DeleteMessageRecord 删除会话记录
func (h *DigitalEmployeeHandler) DeleteMessageRecord(c *gin.Context) {
	msgID := c.Param("id")
	if msgID == "" {
		response.BadRequest(c, "消息ID不能为空", "")
		return
	}

	if err := h.service.DeleteMessageRecord(msgID); err != nil {
		response.InternalServerError(c, "删除会话记录失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// BatchOperateMessageRecords 批量操作会话记录
func (h *DigitalEmployeeHandler) BatchOperateMessageRecords(c *gin.Context) {
	var req model.BatchMessageOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.BatchOperateMessageRecords(&req); err != nil {
		response.BadRequest(c, "批量操作失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "批量操作成功"})
}

// GetConversationGroups 获取会话分组
func (h *DigitalEmployeeHandler) GetConversationGroups(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	groups, err := h.service.GetConversationGroups(limit)
	if err != nil {
		response.InternalServerError(c, "获取会话分组失败", err.Error())
		return
	}

	response.OK(c, groups)
}

// GetMessageStatistics 获取会话统计
func (h *DigitalEmployeeHandler) GetMessageStatistics(c *gin.Context) {
	stats, err := h.service.GetMessageStatistics()
	if err != nil {
		response.InternalServerError(c, "获取统计数据失败", err.Error())
		return
	}

	response.OK(c, stats)
}

// ========== 服务回传记录相关 ==========

// GetExportRecords 获取导出记录列表
func (h *DigitalEmployeeHandler) GetExportRecords(c *gin.Context) {
	var query model.ExportRecordQuery
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

	records, total, err := h.service.GetExportRecords(&query)
	if err != nil {
		response.InternalServerError(c, "获取导出记录失败", err.Error())
		return
	}

	response.OK(c, gin.H{
		"list":      records,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	})
}

// GetExportRecordByID 根据ID获取导出记录
func (h *DigitalEmployeeHandler) GetExportRecordByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的记录ID", err.Error())
		return
	}

	record, err := h.service.GetExportRecordByID(uint(id))
	if err != nil {
		response.NotFound(c, "导出记录不存在", err.Error())
		return
	}

	response.OK(c, record)
}

// CreateExportRecord 创建导出记录
func (h *DigitalEmployeeHandler) CreateExportRecord(c *gin.Context) {
	var req model.ExportRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	record, err := h.service.CreateExportRecord(&req)
	if err != nil {
		response.BadRequest(c, "创建导出记录失败", err.Error())
		return
	}

	response.Created(c, record)
}

// UpdateExportRecord 更新导出记录
func (h *DigitalEmployeeHandler) UpdateExportRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的记录ID", err.Error())
		return
	}

	var req model.ExportRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	record, err := h.service.UpdateExportRecord(uint(id), &req)
	if err != nil {
		response.BadRequest(c, "更新导出记录失败", err.Error())
		return
	}

	response.OK(c, record)
}

// DeleteExportRecord 删除导出记录
func (h *DigitalEmployeeHandler) DeleteExportRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的记录ID", err.Error())
		return
	}

	if err := h.service.DeleteExportRecord(uint(id)); err != nil {
		response.InternalServerError(c, "删除导出记录失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "删除成功"})
}

// BatchOperateExportRecords 批量操作导出记录
func (h *DigitalEmployeeHandler) BatchOperateExportRecords(c *gin.Context) {
	var req model.BatchExportOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	if err := h.service.BatchOperateExportRecords(&req); err != nil {
		response.BadRequest(c, "批量操作失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "批量操作成功"})
}

// RetryExportEmail 重试发送邮件
func (h *DigitalEmployeeHandler) RetryExportEmail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的记录ID", err.Error())
		return
	}

	if err := h.service.RetryExportEmail(uint(id)); err != nil {
		response.BadRequest(c, "重试发送失败", err.Error())
		return
	}

	response.OK(c, gin.H{"message": "已加入重试队列"})
}

// GetExportStatistics 获取导出统计
func (h *DigitalEmployeeHandler) GetExportStatistics(c *gin.Context) {
	stats, err := h.service.GetExportStatistics()
	if err != nil {
		response.InternalServerError(c, "获取统计数据失败", err.Error())
		return
	}

	response.OK(c, stats)
}
