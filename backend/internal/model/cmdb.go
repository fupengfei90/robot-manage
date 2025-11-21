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
