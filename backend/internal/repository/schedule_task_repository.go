package repository

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/model"
)

type ScheduleTaskRepository struct {
	db *gorm.DB
}

func NewScheduleTaskRepository(db *gorm.DB) *ScheduleTaskRepository {
	return &ScheduleTaskRepository{db: db}
}

// GetList 获取任务列表
func (r *ScheduleTaskRepository) GetList(query *model.ScheduleTaskQuery) ([]*model.ScheduleTask, int64, error) {
	var tasks []*model.ScheduleTask
	var total int64

	db := r.db.Model(&model.ScheduleTask{}).Where("deleted_at IS NULL")

	// 构建查询条件
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Active != nil {
		db = db.Where("active = ?", *query.Active)
	}
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}
	if query.DCN != "" {
		db = db.Where("dcn = ?", query.DCN)
	}
	if query.StartTime != "" {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("created_at <= ?", query.EndTime)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 排序和分页
	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}
	orderDir := strings.ToUpper(query.OrderDir)
	if orderDir != "ASC" && orderDir != "DESC" {
		orderDir = "DESC"
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", orderBy, orderDir)).
		Offset(offset).Limit(query.PageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

// GetByID 根据ID获取任务
func (r *ScheduleTaskRepository) GetByID(id uint) (*model.ScheduleTask, error) {
	var task model.ScheduleTask
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// Create 创建任务
func (r *ScheduleTaskRepository) Create(task *model.ScheduleTask) error {
	now := time.Now()
	task.CreatedAt = &now
	task.UpdatedAt = &now
	return r.db.Create(task).Error
}

// Update 更新任务
func (r *ScheduleTaskRepository) Update(id uint, task *model.ScheduleTask) error {
	now := time.Now()
	task.UpdatedAt = &now
	return r.db.Where("id = ? AND deleted_at IS NULL", id).Updates(task).Error
}

// Delete 软删除任务
func (r *ScheduleTaskRepository) Delete(id uint) error {
	now := time.Now()
	return r.db.Model(&model.ScheduleTask{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", now).Error
}

// UpdateStatus 更新任务状态
func (r *ScheduleTaskRepository) UpdateStatus(id uint, active int) error {
	now := time.Now()
	return r.db.Model(&model.ScheduleTask{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"active":     active,
			"updated_at": now,
		}).Error
}

// BatchDelete 批量删除
func (r *ScheduleTaskRepository) BatchDelete(ids []uint) error {
	now := time.Now()
	return r.db.Model(&model.ScheduleTask{}).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Update("deleted_at", now).Error
}

// BatchUpdateStatus 批量更新状态
func (r *ScheduleTaskRepository) BatchUpdateStatus(ids []uint, active int) error {
	now := time.Now()
	return r.db.Model(&model.ScheduleTask{}).
		Where("id IN ? AND deleted_at IS NULL", ids).
		Updates(map[string]interface{}{
			"active":     active,
			"updated_at": now,
		}).Error
}

// GetCategories 获取所有分类
func (r *ScheduleTaskRepository) GetCategories() ([]string, error) {
	var categories []string
	if err := r.db.Model(&model.ScheduleTask{}).
		Where("deleted_at IS NULL").
		Distinct("category").
		Pluck("category", &categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetDCNs 获取所有数据中心
func (r *ScheduleTaskRepository) GetDCNs() ([]string, error) {
	var dcns []string
	if err := r.db.Model(&model.ScheduleTask{}).
		Where("deleted_at IS NULL").
		Distinct("dcn").
		Pluck("dcn", &dcns).Error; err != nil {
		return nil, err
	}
	return dcns, nil
}

// CheckNameExists 检查任务名称是否存在
func (r *ScheduleTaskRepository) CheckNameExists(name, dcn string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&model.ScheduleTask{}).
		Where("name = ? AND dcn = ? AND deleted_at IS NULL", name, dcn)
	
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}