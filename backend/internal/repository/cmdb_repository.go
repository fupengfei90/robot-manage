package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/model"
)

// CMDBRepository CMDB数据访问层
type CMDBRepository struct {
	db *gorm.DB
}

// NewCMDBRepository 创建CMDB Repository
func NewCMDBRepository(db *gorm.DB) *CMDBRepository {
	return &CMDBRepository{db: db}
}

// ========== 值班排班相关 ==========

// GetDutySchedules 获取值班排班列表
func (r *CMDBRepository) GetDutySchedules(startDate, endDate, staffName string, limit, offset int) ([]model.DutySchedule, int64, error) {
	var schedules []model.DutySchedule
	var total int64

	query := r.db.Model(&model.DutySchedule{})

	if startDate != "" {
		query = query.Where("duty_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("duty_date <= ?", endDate)
	}
	if staffName != "" {
		query = query.Where("wb_staff_name LIKE ? OR fb_staff_name LIKE ?", "%"+staffName+"%", "%"+staffName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("duty_date DESC").Limit(limit).Offset(offset).Find(&schedules).Error; err != nil {
		return nil, 0, err
	}

	return schedules, total, nil
}

// GetDutyScheduleByID 根据ID获取值班排班
func (r *CMDBRepository) GetDutyScheduleByID(id int) (*model.DutySchedule, error) {
	var schedule model.DutySchedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

// GetDutyScheduleByDate 根据日期获取值班排班
func (r *CMDBRepository) GetDutyScheduleByDate(date time.Time) (*model.DutySchedule, error) {
	var schedule model.DutySchedule
	if err := r.db.Where("duty_date = ?", date.Format("2006-01-02")).First(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

// CreateDutySchedule 创建值班排班
func (r *CMDBRepository) CreateDutySchedule(schedule *model.DutySchedule) error {
	return r.db.Create(schedule).Error
}

// UpdateDutySchedule 更新值班排班
func (r *CMDBRepository) UpdateDutySchedule(id int, schedule *model.DutySchedule) error {
	return r.db.Model(&model.DutySchedule{}).Where("id = ?", id).Updates(schedule).Error
}

// DeleteDutySchedule 删除值班排班
func (r *CMDBRepository) DeleteDutySchedule(id int) error {
	return r.db.Delete(&model.DutySchedule{}, id).Error
}

// ========== 大事记相关 ==========

// GetMilestoneEvents 获取大事记列表
func (r *CMDBRepository) GetMilestoneEvents(startDate, endDate, eventContent string, isActive *bool, limit, offset int) ([]model.MilestoneEvent, int64, error) {
	var events []model.MilestoneEvent
	var total int64

	query := r.db.Model(&model.MilestoneEvent{})

	if startDate != "" {
		query = query.Where("event_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("event_date <= ?", endDate)
	}
	if eventContent != "" {
		query = query.Where("event_content LIKE ?", "%"+eventContent+"%")
	}
	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("event_date DESC, id DESC").Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// GetMilestoneEventByID 根据ID获取大事记
func (r *CMDBRepository) GetMilestoneEventByID(id int) (*model.MilestoneEvent, error) {
	var event model.MilestoneEvent
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// CreateMilestoneEvent 创建大事记
func (r *CMDBRepository) CreateMilestoneEvent(event *model.MilestoneEvent) error {
	return r.db.Create(event).Error
}

// UpdateMilestoneEvent 更新大事记
func (r *CMDBRepository) UpdateMilestoneEvent(id int, event *model.MilestoneEvent) error {
	return r.db.Model(&model.MilestoneEvent{}).Where("id = ?", id).Updates(event).Error
}

// DeleteMilestoneEvent 删除大事记
func (r *CMDBRepository) DeleteMilestoneEvent(id int) error {
	return r.db.Delete(&model.MilestoneEvent{}, id).Error
}
