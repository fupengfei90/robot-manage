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

// ========== 批量时间相关 ==========

// GetBatchTimes 获取批量时间列表
func (r *CMDBRepository) GetBatchTimes(systemName, subsysName, batchTime string, limit, offset int) ([]model.BatchTime, int64, error) {
	var batchTimes []model.BatchTime
	var total int64

	query := r.db.Model(&model.BatchTime{}).Where("deleted_at IS NULL")

	if systemName != "" {
		query = query.Where("system_name LIKE ?", "%"+systemName+"%")
	}
	if subsysName != "" {
		query = query.Where("subsys_name LIKE ?", "%"+subsysName+"%")
	}
	if batchTime != "" {
		query = query.Where("batch_time LIKE ?", "%"+batchTime+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("updated_at DESC").Limit(limit).Offset(offset).Find(&batchTimes).Error; err != nil {
		return nil, 0, err
	}

	return batchTimes, total, nil
}

// GetBatchTimeByID 根据ID获取批量时间
func (r *CMDBRepository) GetBatchTimeByID(id uint) (*model.BatchTime, error) {
	var batchTime model.BatchTime
	if err := r.db.Where("deleted_at IS NULL").First(&batchTime, id).Error; err != nil {
		return nil, err
	}
	return &batchTime, nil
}

// CreateBatchTime 创建批量时间
func (r *CMDBRepository) CreateBatchTime(batchTime *model.BatchTime) error {
	return r.db.Create(batchTime).Error
}

// UpdateBatchTime 更新批量时间
func (r *CMDBRepository) UpdateBatchTime(id uint, batchTime *model.BatchTime) error {
	return r.db.Model(&model.BatchTime{}).Where("id = ? AND deleted_at IS NULL", id).Updates(batchTime).Error
}

// DeleteBatchTime 删除批量时间（软删除）
func (r *CMDBRepository) DeleteBatchTime(id uint) error {
	now := time.Now()
	return r.db.Model(&model.BatchTime{}).Where("id = ?", id).Update("deleted_at", now).Error
}

// GetSystems 获取系统列表
func (r *CMDBRepository) GetSystems() ([]model.System, error) {
	var systems []model.System
	if err := r.db.Order("name").Find(&systems).Error; err != nil {
		return nil, err
	}
	return systems, nil
}

// GetSubsystemsBySystemID 根据系统ID获取子系统列表
func (r *CMDBRepository) GetSubsystemsBySystemID(systemID uint) ([]model.Subsystem, error) {
	var subsystems []model.Subsystem
	if err := r.db.Where("system_id = ?", systemID).Order("name").Find(&subsystems).Error; err != nil {
		return nil, err
	}
	return subsystems, nil
}

// ========== WB CMDB相关 ==========

func (r *CMDBRepository) GetWBCMDBList(systemName, environment string, limit, offset int) ([]model.WBCMDBInfo, int64, error) {
	var list []model.WBCMDBInfo
	var total int64

	query := r.db.Model(&model.WBCMDBInfo{})

	if systemName != "" {
		query = query.Where("system_name LIKE ?", "%"+systemName+"%")
	}
	if environment != "" {
		query = query.Where("environment LIKE ?", "%"+environment+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id DESC").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *CMDBRepository) GetWBCMDBByID(id uint) (*model.WBCMDBInfo, error) {
	var info model.WBCMDBInfo
	if err := r.db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *CMDBRepository) CreateWBCMDB(info *model.WBCMDBInfo) error {
	return r.db.Create(info).Error
}

func (r *CMDBRepository) UpdateWBCMDB(id uint, info *model.WBCMDBInfo) error {
	return r.db.Model(&model.WBCMDBInfo{}).Where("id = ?", id).Updates(info).Error
}

func (r *CMDBRepository) DeleteWBCMDB(id uint) error {
	return r.db.Delete(&model.WBCMDBInfo{}, id).Error
}

// ========== VB CMDB相关 ==========

func (r *CMDBRepository) GetVBCMDBList(systemName, environment string, limit, offset int) ([]model.VBCMDBInfo, int64, error) {
	var list []model.VBCMDBInfo
	var total int64

	query := r.db.Model(&model.VBCMDBInfo{})

	if systemName != "" {
		query = query.Where("system_name LIKE ?", "%"+systemName+"%")
	}
	if environment != "" {
		query = query.Where("environment LIKE ?", "%"+environment+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id DESC").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *CMDBRepository) GetVBCMDBByID(id uint) (*model.VBCMDBInfo, error) {
	var info model.VBCMDBInfo
	if err := r.db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *CMDBRepository) CreateVBCMDB(info *model.VBCMDBInfo) error {
	return r.db.Create(info).Error
}

func (r *CMDBRepository) UpdateVBCMDB(id uint, info *model.VBCMDBInfo) error {
	return r.db.Model(&model.VBCMDBInfo{}).Where("id = ?", id).Updates(info).Error
}

func (r *CMDBRepository) DeleteVBCMDB(id uint) error {
	return r.db.Delete(&model.VBCMDBInfo{}, id).Error
}

// ========== ITSM出包记录相关 ==========

func (r *CMDBRepository) GetITSMPackageRecords(subsystem, status, environment string, limit, offset int) ([]model.ITSMPackageRecord, int64, error) {
	var records []model.ITSMPackageRecord
	var total int64

	query := r.db.Model(&model.ITSMPackageRecord{})

	if subsystem != "" {
		query = query.Where("subsystem LIKE ?", "%"+subsystem+"%")
	}
	if status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}
	if environment != "" {
		query = query.Where("environment LIKE ?", "%"+environment+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id DESC").Limit(limit).Offset(offset).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (r *CMDBRepository) GetITSMPackageRecordByID(id uint) (*model.ITSMPackageRecord, error) {
	var record model.ITSMPackageRecord
	if err := r.db.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *CMDBRepository) CreateITSMPackageRecord(record *model.ITSMPackageRecord) error {
	return r.db.Create(record).Error
}

func (r *CMDBRepository) UpdateITSMPackageRecord(id uint, record *model.ITSMPackageRecord) error {
	return r.db.Model(&model.ITSMPackageRecord{}).Where("id = ?", id).Updates(record).Error
}

func (r *CMDBRepository) DeleteITSMPackageRecord(id uint) error {
	return r.db.Delete(&model.ITSMPackageRecord{}, id).Error
}
