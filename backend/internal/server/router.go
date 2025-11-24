package server

import (
	"github.com/gin-gonic/gin"

	"github.com/example/robot-manage/internal/config"
	"github.com/example/robot-manage/internal/handler"
	"github.com/example/robot-manage/internal/pkg/response"
	"github.com/example/robot-manage/internal/server/middleware"
)

// NewRouter 注册所有HTTP路由。
func NewRouter(cfg *config.AppConfig, dh *handler.DashboardHandler, ch *handler.CMDBHandler, ah *handler.AuthHandler, sth *handler.ScheduleTaskHandler, rbh *handler.RBACHandler, deh *handler.DigitalEmployeeHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID(), middleware.AccessLog())

	// 认证相关路由（不需要JWT认证）
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", ah.Login)
		auth.POST("/logout", ah.Logout)
	}

	// 需要JWT认证的API路由
	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth(cfg.Security.JWTSecret, cfg.App.Env == "local"))
	{
		// 获取当前用户信息
		api.GET("/auth/me", ah.GetCurrentUser)

		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/summary", dh.Summary)
			dashboard.GET("/knowledge", dh.Knowledge)
			dashboard.GET("/timeline", dh.Timeline)
			dashboard.GET("/assistant-report", dh.AssistantReport)
		}

		registerInfoRoutes(api.Group("/info"), ch)
		registerDeliveryRoutes(api.Group("/delivery"))
		registerDigitalRoutes(api.Group("/digital"), deh)
		registerSystemRoutes(api.Group("/system"), sth, rbh)
	}

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

func registerDigitalRoutes(group *gin.RouterGroup, deh *handler.DigitalEmployeeHandler) {
	// 历史会话记录
	messages := group.Group("/message-records")
	{
		messages.GET("", deh.GetMessageRecords)
		messages.GET("/:id", deh.GetMessageRecordByID)
		messages.POST("", deh.CreateMessageRecord)
		messages.PUT("/:id", deh.UpdateMessageRecord)
		messages.DELETE("/:id", deh.DeleteMessageRecord)
		messages.POST("/batch", deh.BatchOperateMessageRecords)
	}

	group.GET("/conversations/:conversation_id/messages", deh.GetMessagesByConversationID)
	group.GET("/conversation-groups", deh.GetConversationGroups)
	group.GET("/message-statistics", deh.GetMessageStatistics)

	// 服务回传记录
	exports := group.Group("/export-records")
	{
		exports.GET("", deh.GetExportRecords)
		exports.GET("/:id", deh.GetExportRecordByID)
		exports.POST("", deh.CreateExportRecord)
		exports.PUT("/:id", deh.UpdateExportRecord)
		exports.DELETE("/:id", deh.DeleteExportRecord)
		exports.POST("/batch", deh.BatchOperateExportRecords)
		exports.POST("/:id/retry", deh.RetryExportEmail)
	}

	group.GET("/export-statistics", deh.GetExportStatistics)
}

func registerSystemRoutes(group *gin.RouterGroup, sth *handler.ScheduleTaskHandler, rbh *handler.RBACHandler) {
	// 调度任务管理
	tasks := group.Group("/schedule-tasks")
	{
		tasks.GET("", sth.GetList)
		tasks.GET("/:id", sth.GetByID)
		tasks.POST("", sth.Create)
		tasks.PUT("/:id", sth.Update)
		tasks.DELETE("/:id", sth.Delete)
		tasks.PUT("/:id/status", sth.UpdateStatus)
		tasks.POST("/:id/execute", sth.ExecuteTask)
		tasks.POST("/batch", sth.BatchOperation)
		tasks.GET("/categories", sth.GetCategories)
		tasks.GET("/dcns", sth.GetDCNs)
	}

	// RBAC用户角色权限管理
	// 用户管理
	users := group.Group("/users")
	{
		users.GET("", rbh.GetUsers)
		users.GET("/:id", rbh.GetUserByID)
		users.POST("", rbh.CreateUser)
		users.PUT("/:id", rbh.UpdateUser)
		users.DELETE("/:id", rbh.DeleteUser)
		users.PUT("/:id/status", rbh.UpdateUserStatus)
		users.POST("/batch", rbh.BatchOperateUsers)
		users.POST("/:id/roles", rbh.AssignUserRoles)
	}

	// 角色管理
	roles := group.Group("/roles")
	{
		roles.GET("", rbh.GetRoles)
		roles.GET("/:id", rbh.GetRoleByID)
		roles.POST("", rbh.CreateRole)
		roles.PUT("/:id", rbh.UpdateRole)
		roles.DELETE("/:id", rbh.DeleteRole)
		roles.POST("/:id/permissions", rbh.AssignRolePermissions)
	}

	// 权限管理
	permissions := group.Group("/permissions")
	{
		permissions.GET("", rbh.GetPermissions)
		permissions.GET("/tree", rbh.GetPermissionTree)
		permissions.GET("/:id", rbh.GetPermissionByID)
		permissions.POST("", rbh.CreatePermission)
		permissions.PUT("/:id", rbh.UpdatePermission)
		permissions.DELETE("/:id", rbh.DeletePermission)
	}

	group.GET("/pipelines", placeholder("系统管理", "流程编排"))
	group.GET("/models", placeholder("系统管理", "模型管理"))
	group.GET("/mail/templates", placeholder("系统管理", "邮件模板"))
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
