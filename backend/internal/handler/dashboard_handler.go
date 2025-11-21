package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/service"
)

// DashboardHandler HTTP 层，负责参数校验与响应。
type DashboardHandler struct {
	svc *service.DashboardService
}

// NewDashboardHandler 构造函数。
func NewDashboardHandler(svc *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{svc: svc}
}

// Summary 返回概览。
func (h *DashboardHandler) Summary(c *gin.Context) {
	data, err := h.svc.Summary()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, data)
}

// Knowledge 返回知识库动态。
func (h *DashboardHandler) Knowledge(c *gin.Context) {
	data, err := h.svc.Knowledge()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, data)
}

// Timeline 返回日例。
func (h *DashboardHandler) Timeline(c *gin.Context) {
	data, err := h.svc.Timeline()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, data)
}

// AssistantReport 返回助手总结。
func (h *DashboardHandler) AssistantReport(c *gin.Context) {
	rangeType := c.DefaultQuery("range", "daily")
	data, err := h.svc.AssistantReport(rangeType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, data)
}
