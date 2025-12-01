package service

import (
	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

type ProjectPlanService struct {
	repo *repository.ProjectPlanRepository
}

func NewProjectPlanService(repo *repository.ProjectPlanRepository) *ProjectPlanService {
	return &ProjectPlanService{repo: repo}
}

type ModuleWithItems struct {
	model.ProjectModule
	Items []model.ProjectItem `json:"items"`
}

func (s *ProjectPlanService) GetAllModulesWithItems() ([]ModuleWithItems, error) {
	modules, err := s.repo.GetAllModules()
	if err != nil {
		return nil, err
	}

	result := make([]ModuleWithItems, len(modules))
	for i, module := range modules {
		items, err := s.repo.GetItemsByModuleID(module.ID)
		if err != nil {
			return nil, err
		}
		result[i] = ModuleWithItems{
			ProjectModule: module,
			Items:         items,
		}
	}
	return result, nil
}

func (s *ProjectPlanService) CreateModule(module *model.ProjectModule) error {
	return s.repo.CreateModule(module)
}

func (s *ProjectPlanService) UpdateModule(module *model.ProjectModule) error {
	return s.repo.UpdateModule(module)
}

func (s *ProjectPlanService) DeleteModule(id uint) error {
	return s.repo.DeleteModule(id)
}

func (s *ProjectPlanService) CreateItem(item *model.ProjectItem) error {
	return s.repo.CreateItem(item)
}

func (s *ProjectPlanService) UpdateItem(item *model.ProjectItem) error {
	return s.repo.UpdateItem(item)
}

func (s *ProjectPlanService) DeleteItem(id uint) error {
	return s.repo.DeleteItem(id)
}

func (s *ProjectPlanService) BatchUpdateItems(items []model.ProjectItem) error {
	for _, item := range items {
		if err := s.repo.UpdateItem(&item); err != nil {
			return err
		}
	}
	return nil
}
