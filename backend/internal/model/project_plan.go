package model

import "time"

type ProjectModule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Icon      string    `gorm:"size:10" json:"icon"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectItem struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ModuleID     uint      `gorm:"not null;index" json:"module_id"`
	Name         string    `gorm:"size:200;not null" json:"name"`
	Status       string    `gorm:"size:20;default:'not_started'" json:"status"`
	Priority     string    `gorm:"size:20;default:'medium'" json:"priority"`
	Owner        string    `gorm:"size:50" json:"owner"`
	PlannedHours *float64  `json:"planned_hours"`
	ActualHours  *float64  `json:"actual_hours"`
	Description  string    `gorm:"type:text" json:"description"`
	SortOrder    int       `gorm:"default:0" json:"sort_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (ProjectModule) TableName() string {
	return "project_modules"
}

func (ProjectItem) TableName() string {
	return "project_items"
}
