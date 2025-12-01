package model

import "time"

// DutySchedule 值班人员表
type DutySchedule struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	DutyDate    time.Time `json:"dutyDate" gorm:"type:date;uniqueIndex;column:duty_date;comment:值班日期"`
	WBStaffName string    `json:"wbStaffName" gorm:"type:varchar(50);column:wb_staff_name;comment:WB值班人员姓名"`
	FBStaffName string    `json:"fbStaffName" gorm:"type:varchar(50);column:fb_staff_name;comment:FB值班人员姓名"`
	Notified    bool      `json:"notified" gorm:"type:tinyint(1);default:0;comment:是否已通知"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (DutySchedule) TableName() string {
	return "duty_schedule"
}

// MilestoneEvent 大事记表
type MilestoneEvent struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	EventDate    time.Time `json:"eventDate" gorm:"type:date;index;column:event_date;comment:事件日期"`
	DayOfWeek    string    `json:"dayOfWeek" gorm:"type:varchar(10);column:day_of_week;comment:周几"`
	EventContent string    `json:"eventContent" gorm:"type:text;column:event_content;comment:大事记内容"`
	IsActive     bool      `json:"isActive" gorm:"type:tinyint(1);default:1;column:is_active;comment:是否激活"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (MilestoneEvent) TableName() string {
	return "milestone_events"
}

// DutyScheduleCreateRequest 创建值班排班请求
type DutyScheduleCreateRequest struct {
	DutyDate    string `json:"dutyDate" binding:"required"`
	WBStaffName string `json:"wbStaffName" binding:"required"`
	FBStaffName string `json:"fbStaffName" binding:"required"`
}

// DutyScheduleUpdateRequest 更新值班排班请求
type DutyScheduleUpdateRequest struct {
	WBStaffName string `json:"wbStaffName"`
	FBStaffName string `json:"fbStaffName"`
	Notified    *bool  `json:"notified"`
}

// MilestoneEventCreateRequest 创建大事记请求
type MilestoneEventCreateRequest struct {
	EventDate    string `json:"eventDate" binding:"required"`
	EventContent string `json:"eventContent" binding:"required"`
}

// MilestoneEventUpdateRequest 更新大事记请求
type MilestoneEventUpdateRequest struct {
	EventDate    string `json:"eventDate"`
	EventContent string `json:"eventContent"`
	IsActive     *bool  `json:"isActive"`
}

// BatchTime 批量时间表
type BatchTime struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	SystemName string     `json:"systemName" gorm:"type:varchar(100);column:system_name;comment:系统名称"`
	SubsysName string     `json:"subsysName" gorm:"type:varchar(100);column:subsys_name;comment:子系统名称"`
	BatchTime  string     `json:"batchTime" gorm:"type:varchar(100);column:batch_time;comment:批量时间"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt  time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at;index:idx_wecom_batch_time_deleted_at"`
}

func (BatchTime) TableName() string {
	return "wecom_batch_time"
}

// System 系统表
type System struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null;uniqueIndex:uk_systems_name;comment:系统名称"`
	Code        string    `json:"code" gorm:"type:varchar(50);not null;uniqueIndex:uk_systems_code;comment:系统代码"`
	Description string    `json:"description" gorm:"type:text;comment:系统描述"`
	Owner       string    `json:"owner" gorm:"type:varchar(100);comment:负责人"`
	Status      string    `json:"status" gorm:"type:varchar(20);default:'active';comment:状态"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (System) TableName() string {
	return "systems"
}

// Subsystem 子系统表
type Subsystem struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SystemID    uint      `json:"systemId" gorm:"not null;index:fk_subsystems_system;comment:所属系统ID"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null;comment:子系统名称"`
	Code        string    `json:"code" gorm:"type:varchar(50);not null;comment:子系统代码"`
	Description string    `json:"description" gorm:"type:text;comment:子系统描述"`
	BatchWindow string    `json:"batchWindow" gorm:"type:varchar(100);column:batch_window;comment:批量时间窗口"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Subsystem) TableName() string {
	return "subsystems"
}

// BatchTimeCreateRequest 创建批量时间请求
type BatchTimeCreateRequest struct {
	SystemName string `json:"systemName" binding:"required"`
	SubsysName string `json:"subsysName" binding:"required"`
	BatchTime  string `json:"batchTime" binding:"required"`
}

// BatchTimeUpdateRequest 更新批量时间请求
type BatchTimeUpdateRequest struct {
	SystemName string `json:"systemName"`
	SubsysName string `json:"subsysName"`
	BatchTime  string `json:"batchTime"`
}

// BatchTimeWithDetails 批量时间详情（包含关联信息）
type BatchTimeWithDetails struct {
	BatchTime
	SystemCode   string `json:"systemCode,omitempty"`
	SystemStatus string `json:"systemStatus,omitempty"`
	SystemOwner  string `json:"systemOwner,omitempty"`
	SubsysCode   string `json:"subsysCode,omitempty"`
}

// WBCMDBInfo WB CMDB信息表
type WBCMDBInfo struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SystemName  string    `json:"systemName" gorm:"type:varchar(100);column:system_name;comment:系统名称"`
	Environment string    `json:"environment" gorm:"type:varchar(50);column:environment;comment:环境"`
	IPAddress   string    `json:"ipAddress" gorm:"type:varchar(50);column:ip_address;comment:IP地址"`
	Port        string    `json:"port" gorm:"type:varchar(20);column:port;comment:端口"`
	Status      string    `json:"status" gorm:"type:varchar(20);column:status;comment:状态"`
	Owner       string    `json:"owner" gorm:"type:varchar(100);column:owner;comment:负责人"`
	Remark      string    `json:"remark" gorm:"type:text;column:remark;comment:备注"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (WBCMDBInfo) TableName() string {
	return "wb_cmdb_info"
}

// VBCMDBInfo VB CMDB信息表
type VBCMDBInfo struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SystemName  string    `json:"systemName" gorm:"type:varchar(100);column:system_name;comment:系统名称"`
	Environment string    `json:"environment" gorm:"type:varchar(50);column:environment;comment:环境"`
	IPAddress   string    `json:"ipAddress" gorm:"type:varchar(50);column:ip_address;comment:IP地址"`
	Port        string    `json:"port" gorm:"type:varchar(20);column:port;comment:端口"`
	Status      string    `json:"status" gorm:"type:varchar(20);column:status;comment:状态"`
	Owner       string    `json:"owner" gorm:"type:varchar(100);column:owner;comment:负责人"`
	Remark      string    `json:"remark" gorm:"type:text;column:remark;comment:备注"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (VBCMDBInfo) TableName() string {
	return "vb_cmdb_info"
}

// WBCMDBInfoRequest WB CMDB信息请求
type WBCMDBInfoRequest struct {
	SystemName  string `json:"systemName" binding:"required"`
	Environment string `json:"environment"`
	IPAddress   string `json:"ipAddress"`
	Port        string `json:"port"`
	Status      string `json:"status"`
	Owner       string `json:"owner"`
	Remark      string `json:"remark"`
}

// VBCMDBInfoRequest VB CMDB信息请求
type VBCMDBInfoRequest struct {
	SystemName  string `json:"systemName" binding:"required"`
	Environment string `json:"environment"`
	IPAddress   string `json:"ipAddress"`
	Port        string `json:"port"`
	Status      string `json:"status"`
	Owner       string `json:"owner"`
	Remark      string `json:"remark"`
}

// ITSMPackageRecord ITSM出包记录表
type ITSMPackageRecord struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Subsystem     string    `json:"subsystem" gorm:"type:varchar(100);column:subsystem;comment:子系统"`
	PackageName   string    `json:"packageName" gorm:"type:varchar(255);column:package_name;comment:物料包"`
	RequirementID string    `json:"requirementId" gorm:"type:varchar(50);column:requirement_id;comment:需求ID"`
	ITSMTicket    string    `json:"itsmTicket" gorm:"type:varchar(50);column:itsm_ticket;comment:ITSM单"`
	Status        string    `json:"status" gorm:"type:varchar(50);column:status;comment:状态"`
	Owner         string    `json:"owner" gorm:"type:varchar(100);column:owner;comment:负责人"`
	Environment   string    `json:"environment" gorm:"type:varchar(50);column:environment;comment:发布环境"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (ITSMPackageRecord) TableName() string {
	return "itsm_package_records"
}

// ITSMPackageRecordRequest ITSM出包记录请求
type ITSMPackageRecordRequest struct {
	Subsystem     string `json:"subsystem" binding:"required"`
	PackageName   string `json:"packageName" binding:"required"`
	RequirementID string `json:"requirementId" binding:"required"`
	ITSMTicket    string `json:"itsmTicket" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Owner         string `json:"owner" binding:"required"`
	Environment   string `json:"environment" binding:"required"`
}
