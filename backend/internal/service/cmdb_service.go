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
func (s *CMDBService) GetDutySchedules(startDate, endDate, staffName string, page, pageSize int) ([]model.DutySchedule, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetDutySchedules(startDate, endDate, staffName, pageSize, offset)
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
func (s *CMDBService) GetMilestoneEvents(startDate, endDate, eventContent string, isActive *bool, page, pageSize int) ([]model.MilestoneEvent, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetMilestoneEvents(startDate, endDate, eventContent, isActive, pageSize, offset)
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

// ========== 批量时间相关 ==========

// GetBatchTimes 获取批量时间列表
func (s *CMDBService) GetBatchTimes(systemName, subsysName, batchTime string, page, pageSize int) ([]model.BatchTime, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetBatchTimes(systemName, subsysName, batchTime, pageSize, offset)
}

// GetBatchTimeByID 根据ID获取批量时间
func (s *CMDBService) GetBatchTimeByID(id uint) (*model.BatchTime, error) {
	return s.repo.GetBatchTimeByID(id)
}

// CreateBatchTime 创建批量时间
func (s *CMDBService) CreateBatchTime(req *model.BatchTimeCreateRequest) (*model.BatchTime, error) {
	batchTime := &model.BatchTime{
		SystemName: req.SystemName,
		SubsysName: req.SubsysName,
		BatchTime:  req.BatchTime,
	}

	if err := s.repo.CreateBatchTime(batchTime); err != nil {
		return nil, err
	}

	return batchTime, nil
}

// UpdateBatchTime 更新批量时间
func (s *CMDBService) UpdateBatchTime(id uint, req *model.BatchTimeUpdateRequest) (*model.BatchTime, error) {
	batchTime, err := s.repo.GetBatchTimeByID(id)
	if err != nil {
		return nil, fmt.Errorf("批量时间不存在: %w", err)
	}

	if req.SystemName != "" {
		batchTime.SystemName = req.SystemName
	}
	if req.SubsysName != "" {
		batchTime.SubsysName = req.SubsysName
	}
	if req.BatchTime != "" {
		batchTime.BatchTime = req.BatchTime
	}

	if err := s.repo.UpdateBatchTime(id, batchTime); err != nil {
		return nil, err
	}

	return batchTime, nil
}

// DeleteBatchTime 删除批量时间
func (s *CMDBService) DeleteBatchTime(id uint) error {
	_, err := s.repo.GetBatchTimeByID(id)
	if err != nil {
		return fmt.Errorf("批量时间不存在: %w", err)
	}
	return s.repo.DeleteBatchTime(id)
}

// GetSystems 获取系统列表
func (s *CMDBService) GetSystems() ([]model.System, error) {
	return s.repo.GetSystems()
}

// GetSubsystemsBySystemID 根据系统ID获取子系统列表
func (s *CMDBService) GetSubsystemsBySystemID(systemID uint) ([]model.Subsystem, error) {
	return s.repo.GetSubsystemsBySystemID(systemID)
}

// ========== WB CMDB相关 ==========

func (s *CMDBService) GetWBCMDBList(systemName, environment string, page, pageSize int) ([]model.WBCMDBInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetWBCMDBList(systemName, environment, pageSize, offset)
}

func (s *CMDBService) GetWBCMDBByID(id uint) (*model.WBCMDBInfo, error) {
	return s.repo.GetWBCMDBByID(id)
}

func (s *CMDBService) CreateWBCMDB(req *model.WBCMDBInfoRequest) (*model.WBCMDBInfo, error) {
	info := &model.WBCMDBInfo{
		SystemName:  req.SystemName,
		Environment: req.Environment,
		IPAddress:   req.IPAddress,
		Port:        req.Port,
		Status:      req.Status,
		Owner:       req.Owner,
		Remark:      req.Remark,
	}
	if err := s.repo.CreateWBCMDB(info); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *CMDBService) UpdateWBCMDB(id uint, req *model.WBCMDBInfoRequest) (*model.WBCMDBInfo, error) {
	info, err := s.repo.GetWBCMDBByID(id)
	if err != nil {
		return nil, fmt.Errorf("记录不存在: %w", err)
	}

	info.SystemName = req.SystemName
	info.Environment = req.Environment
	info.IPAddress = req.IPAddress
	info.Port = req.Port
	info.Status = req.Status
	info.Owner = req.Owner
	info.Remark = req.Remark

	if err := s.repo.UpdateWBCMDB(id, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *CMDBService) DeleteWBCMDB(id uint) error {
	_, err := s.repo.GetWBCMDBByID(id)
	if err != nil {
		return fmt.Errorf("记录不存在: %w", err)
	}
	return s.repo.DeleteWBCMDB(id)
}

// ========== VB CMDB相关 ==========

func (s *CMDBService) GetVBCMDBList(systemName, environment string, page, pageSize int) ([]model.VBCMDBInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetVBCMDBList(systemName, environment, pageSize, offset)
}

func (s *CMDBService) GetVBCMDBByID(id uint) (*model.VBCMDBInfo, error) {
	return s.repo.GetVBCMDBByID(id)
}

func (s *CMDBService) CreateVBCMDB(req *model.VBCMDBInfoRequest) (*model.VBCMDBInfo, error) {
	info := &model.VBCMDBInfo{
		SystemName:  req.SystemName,
		Environment: req.Environment,
		IPAddress:   req.IPAddress,
		Port:        req.Port,
		Status:      req.Status,
		Owner:       req.Owner,
		Remark:      req.Remark,
	}
	if err := s.repo.CreateVBCMDB(info); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *CMDBService) UpdateVBCMDB(id uint, req *model.VBCMDBInfoRequest) (*model.VBCMDBInfo, error) {
	info, err := s.repo.GetVBCMDBByID(id)
	if err != nil {
		return nil, fmt.Errorf("记录不存在: %w", err)
	}

	info.SystemName = req.SystemName
	info.Environment = req.Environment
	info.IPAddress = req.IPAddress
	info.Port = req.Port
	info.Status = req.Status
	info.Owner = req.Owner
	info.Remark = req.Remark

	if err := s.repo.UpdateVBCMDB(id, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *CMDBService) DeleteVBCMDB(id uint) error {
	_, err := s.repo.GetVBCMDBByID(id)
	if err != nil {
		return fmt.Errorf("记录不存在: %w", err)
	}
	return s.repo.DeleteVBCMDB(id)
}

// ========== ITSM出包记录相关 ==========

func (s *CMDBService) GetITSMPackageRecords(subsystem, status, environment string, page, pageSize int) ([]model.ITSMPackageRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.GetITSMPackageRecords(subsystem, status, environment, pageSize, offset)
}

func (s *CMDBService) GetITSMPackageRecordByID(id uint) (*model.ITSMPackageRecord, error) {
	return s.repo.GetITSMPackageRecordByID(id)
}

func (s *CMDBService) CreateITSMPackageRecord(req *model.ITSMPackageRecordRequest) (*model.ITSMPackageRecord, error) {
	record := &model.ITSMPackageRecord{
		Subsystem:     req.Subsystem,
		PackageName:   req.PackageName,
		RequirementID: req.RequirementID,
		ITSMTicket:    req.ITSMTicket,
		Status:        req.Status,
		Owner:         req.Owner,
		Environment:   req.Environment,
	}
	if err := s.repo.CreateITSMPackageRecord(record); err != nil {
		return nil, err
	}
	return record, nil
}

func (s *CMDBService) UpdateITSMPackageRecord(id uint, req *model.ITSMPackageRecordRequest) (*model.ITSMPackageRecord, error) {
	record, err := s.repo.GetITSMPackageRecordByID(id)
	if err != nil {
		return nil, fmt.Errorf("记录不存在: %w", err)
	}

	record.Subsystem = req.Subsystem
	record.PackageName = req.PackageName
	record.RequirementID = req.RequirementID
	record.ITSMTicket = req.ITSMTicket
	record.Status = req.Status
	record.Owner = req.Owner
	record.Environment = req.Environment

	if err := s.repo.UpdateITSMPackageRecord(id, record); err != nil {
		return nil, err
	}
	return record, nil
}

func (s *CMDBService) DeleteITSMPackageRecord(id uint) error {
	_, err := s.repo.GetITSMPackageRecordByID(id)
	if err != nil {
		return fmt.Errorf("记录不存在: %w", err)
	}
	return s.repo.DeleteITSMPackageRecord(id)
}
