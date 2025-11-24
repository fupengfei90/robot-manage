package model

import (
	"time"
)

// ScheduleTask 调度任务模型 (基于wecom_tasks表)
type ScheduleTask struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"not null;size:100;comment:任务名称" json:"name"`
	Active      int        `gorm:"default:0;comment:是否可用" json:"active"`
	Cron        string     `gorm:"not null;size:100;comment:定时任务表达式" json:"cron"`
	Workflow    string     `gorm:"not null;size:100;comment:工作流ID" json:"workflow"`
	ExecToken   string     `gorm:"not null;size:100;comment:执行token" json:"exec_token"`
	CreatedAt   *time.Time `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"comment:更新时间" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"comment:删除时间" json:"deleted_at"`
	CronEntryID int64      `gorm:"not null;default:0" json:"cron_entry_id"`
	Category    string     `gorm:"not null;size:32;comment:任务分类" json:"category"`
	DCN         string     `gorm:"not null;size:100;comment:数据中心" json:"dcn"`
}

// TableName 指定表名
func (ScheduleTask) TableName() string {
	return "wecom_tasks"
}

// ScheduleTaskRequest 创建/更新任务请求
type ScheduleTaskRequest struct {
	Name      string `json:"name" binding:"required,max=100"`
	Active    int    `json:"active"`
	Cron      string `json:"cron" binding:"required,max=100"`
	Workflow  string `json:"workflow" binding:"required,max=100"`
	ExecToken string `json:"exec_token" binding:"required,max=100"`
	Category  string `json:"category" binding:"required,max=32"`
	DCN       string `json:"dcn" binding:"required,max=100"`
}

// ScheduleTaskQuery 查询参数
type ScheduleTaskQuery struct {
	Name       string `form:"name"`
	Active     *int   `form:"active"`
	Category   string `form:"category"`
	DCN        string `form:"dcn"`
	StartTime  string `form:"start_time"`
	EndTime    string `form:"end_time"`
	Page       int    `form:"page,default=1"`
	PageSize   int    `form:"page_size,default=10"`
	OrderBy    string `form:"order_by,default=id"`
	OrderDir   string `form:"order_dir,default=desc"`
}

// ScheduleTaskResponse 任务响应
type ScheduleTaskResponse struct {
	*ScheduleTask
	NextRunTime   *string `json:"next_run_time"`
	CronDesc      string  `json:"cron_desc"`
	StatusText    string  `json:"status_text"`
}

// BatchOperationRequest 批量操作请求
type BatchOperationRequest struct {
	IDs       []uint `json:"ids" binding:"required"`
	Operation string `json:"operation" binding:"required,oneof=enable disable delete"`
}