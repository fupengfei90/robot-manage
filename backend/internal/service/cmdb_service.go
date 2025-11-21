package service

import (
	"fmt"
	"time"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

// CMDBService CMDB业务逻辑层
type CMDBService struct {
	repo *repository.CMDBRepository
}

// NewCMDBService 创建CMDB Service
func NewCMDBService(repo *repository.CMDBRepository) *CMDBService {
	return &CMDBService{repo: repo}
}

// ========== 值班排班相关 ==========

// GetDutySchedules 获取值班排班列表
func (s *CMDBService) GetDutySchedules(startDate, endDate string, page, pageSize int) ([]model.DutySchedule, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetDutySchedules(startDate, endDate, pageSize, offset)
}

// GetDutyScheduleByID 根据ID获取值班排班
func (s *CMDBService) GetDutyScheduleByID(id int) (*model.DutySchedule, error) {
	return s.repo.GetDutyScheduleByID(id)
}

// CreateDutySchedule 创建值班排班
func (s *CMDBService) CreateDutySchedule(req *model.DutyScheduleCreateRequest) (*model.DutySchedule, error) {
	dutyDate, err := time.Parse("2006-01-02", req.DutyDate)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %w", err)
	}

	// 检查是否已存在
	existing, _ := s.repo.GetDutyScheduleByDate(dutyDate)
	if existing != nil {
		return nil, fmt.Errorf("该日期已存在值班排班")
	}

	schedule := &model.DutySchedule{
		DutyDate:    dutyDate,
		WBStaffName: req.WBStaffName,
		FBStaffName: req.FBStaffName,
		Notified:    false,
	}

	if err := s.repo.CreateDutySchedule(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// UpdateDutySchedule 更新值班排班
func (s *CMDBService) UpdateDutySchedule(id int, req *model.DutyScheduleUpdateRequest) (*model.DutySchedule, error) {
	schedule, err := s.repo.GetDutyScheduleByID(id)
	if err != nil {
		return nil, fmt.Errorf("值班排班不存在: %w", err)
	}

	if req.WBStaffName != "" {
		schedule.WBStaffName = req.WBStaffName
	}
	if req.FBStaffName != "" {
		schedule.FBStaffName = req.FBStaffName
	}
	if req.Notified != nil {
		schedule.Notified = *req.Notified
	}

	if err := s.repo.UpdateDutySchedule(id, schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// DeleteDutySchedule 删除值班排班
func (s *CMDBService) DeleteDutySchedule(id int) error {
	_, err := s.repo.GetDutyScheduleByID(id)
	if err != nil {
		return fmt.Errorf("值班排班不存在: %w", err)
	}
	return s.repo.DeleteDutySchedule(id)
}

// ========== 大事记相关 ==========

// GetMilestoneEvents 获取大事记列表
func (s *CMDBService) GetMilestoneEvents(startDate, endDate string, isActive *bool, page, pageSize int) ([]model.MilestoneEvent, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetMilestoneEvents(startDate, endDate, isActive, pageSize, offset)
}

// GetMilestoneEventByID 根据ID获取大事记
func (s *CMDBService) GetMilestoneEventByID(id int) (*model.MilestoneEvent, error) {
	return s.repo.GetMilestoneEventByID(id)
}

// CreateMilestoneEvent 创建大事记
func (s *CMDBService) CreateMilestoneEvent(req *model.MilestoneEventCreateRequest) (*model.MilestoneEvent, error) {
	eventDate, err := time.Parse("2006-01-02", req.EventDate)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %w", err)
	}

	dayOfWeek := eventDate.Weekday().String()
	dayNames := map[time.Weekday]string{
		time.Sunday:    "星期日",
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
	}
	if cnName, ok := dayNames[eventDate.Weekday()]; ok {
		dayOfWeek = cnName
	}

	event := &model.MilestoneEvent{
		EventDate:    eventDate,
		DayOfWeek:    dayOfWeek,
		EventContent: req.EventContent,
		IsActive:     true,
	}

	if err := s.repo.CreateMilestoneEvent(event); err != nil {
		return nil, err
	}

	return event, nil
}

// UpdateMilestoneEvent 更新大事记
func (s *CMDBService) UpdateMilestoneEvent(id int, req *model.MilestoneEventUpdateRequest) (*model.MilestoneEvent, error) {
	event, err := s.repo.GetMilestoneEventByID(id)
	if err != nil {
		return nil, fmt.Errorf("大事记不存在: %w", err)
	}

	if req.EventDate != "" {
		eventDate, err := time.Parse("2006-01-02", req.EventDate)
		if err != nil {
			return nil, fmt.Errorf("日期格式错误: %w", err)
		}
		event.EventDate = eventDate

		dayOfWeek := eventDate.Weekday().String()
		dayNames := map[time.Weekday]string{
			time.Sunday:    "星期日",
			time.Monday:    "星期一",
			time.Tuesday:   "星期二",
			time.Wednesday: "星期三",
			time.Thursday:  "星期四",
			time.Friday:    "星期五",
			time.Saturday:  "星期六",
		}
		if cnName, ok := dayNames[eventDate.Weekday()]; ok {
			dayOfWeek = cnName
		}
		event.DayOfWeek = dayOfWeek
	}

	if req.EventContent != "" {
		event.EventContent = req.EventContent
	}

	if req.IsActive != nil {
		event.IsActive = *req.IsActive
	}

	if err := s.repo.UpdateMilestoneEvent(id, event); err != nil {
		return nil, err
	}

	return event, nil
}

// DeleteMilestoneEvent 删除大事记
func (s *CMDBService) DeleteMilestoneEvent(id int) error {
	_, err := s.repo.GetMilestoneEventByID(id)
	if err != nil {
		return fmt.Errorf("大事记不存在: %w", err)
	}
	return s.repo.DeleteMilestoneEvent(id)
}
