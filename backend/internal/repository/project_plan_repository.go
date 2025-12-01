package repository

import (
	"github.com/example/robot-manage/internal/model"
	"gorm.io/gorm"
)

type ProjectPlanRepository struct {
	db *gorm.DB
}

func NewProjectPlanRepository(db *gorm.DB) *ProjectPlanRepository {
	return &ProjectPlanRepository{db: db}
}

func (r *ProjectPlanRepository) GetAllModules() ([]model.ProjectModule, error) {
	var modules []model.ProjectModule
	err := r.db.Order("sort_order, id").Find(&modules).Error
	return modules, err
}

func (r *ProjectPlanRepository) CreateModule(module *model.ProjectModule) error {
	return r.db.Create(module).Error
}

func (r *ProjectPlanRepository) UpdateModule(module *model.ProjectModule) error {
	return r.db.Save(module).Error
}

func (r *ProjectPlanRepository) DeleteModule(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("module_id = ?", id).Delete(&model.ProjectItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.ProjectModule{}, id).Error
	})
}

func (r *ProjectPlanRepository) GetItemsByModuleID(moduleID uint) ([]model.ProjectItem, error) {
	var items []model.ProjectItem
	err := r.db.Where("module_id = ?", moduleID).Order("sort_order, id").Find(&items).Error
	return items, err
}

func (r *ProjectPlanRepository) CreateItem(item *model.ProjectItem) error {
	return r.db.Create(item).Error
}

func (r *ProjectPlanRepository) UpdateItem(item *model.ProjectItem) error {
	return r.db.Save(item).Error
}

func (r *ProjectPlanRepository) DeleteItem(id uint) error {
	return r.db.Delete(&model.ProjectItem{}, id).Error
}

func (r *ProjectPlanRepository) BatchUpdateItems(items []model.ProjectItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Save(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
