package server

import (
	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/config"
	"github.com/example/robot-manage/internal/handler"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/server/middleware"
)

// NewRouter 注册所有HTTP路由。
func NewRouter(cfg *config.AppConfig, dh *handler.DashboardHandler, ch *handler.CMDBHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID(), middleware.AccessLog())

	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth(cfg.Security.JWTSecret, cfg.App.Env == "local"))

	dashboard := api.Group("/dashboard")
	{
		dashboard.GET("/summary", dh.Summary)
		dashboard.GET("/knowledge", dh.Knowledge)
		dashboard.GET("/timeline", dh.Timeline)
		dashboard.GET("/assistant-report", dh.AssistantReport)
	}

	registerInfoRoutes(api.Group("/info"), ch)
	registerDeliveryRoutes(api.Group("/delivery"))
	registerDigitalRoutes(api.Group("/digital"))
	registerSystemRoutes(api.Group("/system"))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}

func registerInfoRoutes(group *gin.RouterGroup, ch *handler.CMDBHandler) {
	// CMDB - 值班排班
	duty := group.Group("/duty")
	{
		duty.GET("/schedules", ch.GetDutySchedules)
		duty.GET("/schedules/:id", ch.GetDutyScheduleByID)
		duty.POST("/schedules", ch.CreateDutySchedule)
		duty.PUT("/schedules/:id", ch.UpdateDutySchedule)
		duty.DELETE("/schedules/:id", ch.DeleteDutySchedule)
	}

	// CMDB - 大事记
	milestones := group.Group("/milestones")
	{
		milestones.GET("", ch.GetMilestoneEvents)
		milestones.GET("/:id", ch.GetMilestoneEventByID)
		milestones.POST("", ch.CreateMilestoneEvent)
		milestones.PUT("/:id", ch.UpdateMilestoneEvent)
		milestones.DELETE("/:id", ch.DeleteMilestoneEvent)
	}

	// 其他占位接口
	group.GET("/cmdb/environments", placeholder("信息管理", "CMDB环境列表"))
	group.GET("/alerts/ignores", placeholder("信息管理", "可忽略告警"))
	group.GET("/itsm/tickets", placeholder("信息管理", "安全工单"))
	group.GET("/batches/windows", placeholder("信息管理", "批量时间窗口"))
}

func registerDeliveryRoutes(group *gin.RouterGroup) {
	group.GET("/inspection/reports", placeholder("智能交付", "巡检分析"))
	group.GET("/alerts/analysis", placeholder("智能交付", "告警分析"))
	group.POST("/tickets/ai", placeholder("智能交付", "智能提单"))
	group.GET("/events", placeholder("智能交付", "事件管理"))
	group.GET("/releases", placeholder("智能交付", "上线检视"))
	group.GET("/versions", placeholder("智能交付", "版本分析"))
	group.GET("/sql/insights", placeholder("智能交付", "SQL洞察"))
	group.GET("/knowledge", placeholder("智能交付", "知识库管理"))
}

func registerDigitalRoutes(group *gin.RouterGroup) {
	group.GET("/sessions", placeholder("数字员工", "对话历史"))
	group.POST("/feedback", placeholder("数字员工", "服务回传"))
}

func registerSystemRoutes(group *gin.RouterGroup) {
	group.GET("/tasks", placeholder("系统管理", "任务调度"))
	group.GET("/pipelines", placeholder("系统管理", "流程编排"))
	group.GET("/models", placeholder("系统管理", "模型管理"))
	group.GET("/mail/templates", placeholder("系统管理", "邮件模板"))
	group.GET("/rbac/roles", placeholder("系统管理", "RBAC角色列表"))
}

func placeholder(module, feature string) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.OK(c, gin.H{
			"module":  module,
			"feature": feature,
			"status":  "pending",
		})
	}
}
