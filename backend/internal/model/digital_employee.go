package model

import (
	"time"
)

// MessageRecord 历史会话记录模型
type MessageRecord struct {
	MsgID           string     `gorm:"primaryKey;column:msg_id;size:128" json:"msg_id"`
	Input           string     `gorm:"column:input;type:text" json:"input"`
	Output          string     `gorm:"column:output;type:text" json:"output"`
	ConversationID  string     `gorm:"column:conversation_id;size:128;index" json:"conversation_id"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	RequestUserID   string     `gorm:"column:request_user_id;size:128;index" json:"request_user_id"`
	RequestUserName string     `gorm:"column:request_user_name;size:128;index" json:"request_user_name"`
	ExtraData       string     `gorm:"column:extra_data;type:text" json:"extra_data"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Debug           string     `gorm:"column:debug;size:10" json:"debug"`
	Type            string     `gorm:"column:type;size:50;index" json:"type"` // 会话类型：question, command, chat, scheduled_task
	IsValid         *int       `gorm:"column:is_valid;default:1;index" json:"is_valid"` // 是否有效：1有效，0无效
}

func (MessageRecord) TableName() string {
	return "wecom_message_records"
}

// ExportRecord 服务回传记录模型
type ExportRecord struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ExportTime      time.Time  `gorm:"column:export_time;not null" json:"export_time"`
	StartDate       string     `gorm:"column:start_date;size:10;not null" json:"start_date"`
	EndDate         string     `gorm:"column:end_date;size:10;not null" json:"end_date"`
	TotalRecords    int64      `gorm:"column:total_records;not null" json:"total_records"`
	FileName        string     `gorm:"column:file_name;size:255;not null" json:"file_name"`
	FilePath        string     `gorm:"column:file_path;size:512;not null" json:"file_path"`
	EmailSendTime   *time.Time `gorm:"column:email_send_time" json:"email_send_time"`
	EmailSubject    string     `gorm:"column:email_subject;size:255" json:"email_subject"`
	EmailStatus     string     `gorm:"column:email_status;size:20" json:"email_status"`
	EmailRecipients string     `gorm:"column:email_recipients;type:text" json:"email_recipients"`
	EmailCC         string     `gorm:"column:email_cc;type:text" json:"email_cc"`
	EmailBCC        string     `gorm:"column:email_bcc;type:text" json:"email_bcc"`
	ErrorReason     string     `gorm:"column:error_reason;type:text" json:"error_reason"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ExportRecord) TableName() string {
	return "wecom_export_records"
}

// MessageRecordQuery 会话记录查询参数
type MessageRecordQuery struct {
	MsgID           string `form:"msg_id"`
	ConversationID  string `form:"conversation_id"`
	RequestUserID   string `form:"request_user_id"`
	RequestUserName string `form:"request_user_name"`
	Debug           string `form:"debug"`
	Type            string `form:"type"`
	IsValid         *int   `form:"is_valid"`
	StartTime       string `form:"start_time"`
	EndTime         string `form:"end_time"`
	Keyword         string `form:"keyword"`
	Page            int    `form:"page,default=1"`
	PageSize        int    `form:"page_size,default=10"`
	OrderBy         string `form:"order_by,default=created_at"`
	OrderDir        string `form:"order_dir,default=desc"`
}

// ExportRecordQuery 导出记录查询参数
type ExportRecordQuery struct {
	FileName     string `form:"file_name"`
	EmailSubject string `form:"email_subject"`
	EmailStatus  string `form:"email_status"`
	StartDate    string `form:"start_date"`
	EndDate      string `form:"end_date"`
	StartTime    string `form:"start_time"`
	EndTime      string `form:"end_time"`
	Page         int    `form:"page,default=1"`
	PageSize     int    `form:"page_size,default=10"`
	OrderBy      string `form:"order_by,default=export_time"`
	OrderDir     string `form:"order_dir,default=desc"`
}

// MessageRecordRequest 会话记录请求
type MessageRecordRequest struct {
	MsgID           string `json:"msg_id" binding:"required"`
	Input           string `json:"input"`
	Output          string `json:"output"`
	ConversationID  string `json:"conversation_id"`
	RequestUserID   string `json:"request_user_id"`
	RequestUserName string `json:"request_user_name"`
	ExtraData       string `json:"extra_data"`
	Debug           string `json:"debug"`
	Type            string `json:"type"`
	IsValid         *int   `json:"is_valid"`
}

// ExportRecordRequest 导出记录请求
type ExportRecordRequest struct {
	StartDate       string `json:"start_date" binding:"required"`
	EndDate         string `json:"end_date" binding:"required"`
	TotalRecords    int64  `json:"total_records" binding:"required"`
	FileName        string `json:"file_name" binding:"required"`
	FilePath        string `json:"file_path" binding:"required"`
	EmailSubject    string `json:"email_subject"`
	EmailRecipients string `json:"email_recipients"`
	EmailCC         string `json:"email_cc"`
	EmailBCC        string `json:"email_bcc"`
}

// BatchOperationRequest 批量操作请求
type BatchMessageOperationRequest struct {
	MsgIDs    []string `json:"msg_ids" binding:"required"`
	Operation string   `json:"operation" binding:"required,oneof=delete mark_debug unmark_debug"`
}

// BatchExportOperationRequest 批量导出操作请求
type BatchExportOperationRequest struct {
	IDs       []uint `json:"ids" binding:"required"`
	Operation string `json:"operation" binding:"required,oneof=delete retry"`
}

// ConversationGroup 会话分组
type ConversationGroup struct {
	ConversationID  string           `json:"conversation_id"`
	RequestUserID   string           `json:"request_user_id"`
	RequestUserName string           `json:"request_user_name"`
	MessageCount    int              `json:"message_count"`
	FirstMessageAt  *time.Time       `json:"first_message_at"`
	LastMessageAt   *time.Time       `json:"last_message_at"`
	Messages        []MessageRecord  `json:"messages"`
}

// MessageStatistics 会话统计
type MessageStatistics struct {
	TotalMessages      int64              `json:"total_messages"`
	TotalConversations int64              `json:"total_conversations"`
	TotalUsers         int64              `json:"total_users"`
	DebugMessages      int64              `json:"debug_messages"`
	TodayMessages      int64              `json:"today_messages"`
	TypeDistribution   []TypeDistribution `json:"type_distribution"`
	TrendData          []TrendDataPoint   `json:"trend_data"`
	TopUsers           []UserActivityData `json:"top_users"`
}

// ExportStatistics 导出统计
type ExportStatistics struct {
	TotalExports   int64            `json:"total_exports"`
	SuccessCount   int64            `json:"success_count"`
	FailedCount    int64            `json:"failed_count"`
	PendingCount   int64            `json:"pending_count"`
	SuccessRate    float64          `json:"success_rate"`
	TotalRecords   int64            `json:"total_records"`
	TrendData      []TrendDataPoint `json:"trend_data"`
}

// TrendDataPoint 趋势数据点
type TrendDataPoint struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// UserActivityData 用户活跃度数据
type UserActivityData struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Count    int64  `json:"count"`
}

// TypeDistribution 类型分布数据
type TypeDistribution struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}
