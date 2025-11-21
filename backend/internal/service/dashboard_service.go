package service

import (
	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

// DashboardService 封装业务逻辑。
type DashboardService struct {
	repo *repository.DashboardRepository
}

// NewDashboardService 构造函数。
func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) Summary() (*model.DashboardSummary, error) {
	return s.repo.FetchSummary()
}

func (s *DashboardService) Knowledge() (*model.KnowledgeSnapshot, error) {
	return s.repo.FetchKnowledge()
}

func (s *DashboardService) Timeline() (*model.TimelineInfo, error) {
	return s.repo.FetchTimeline()
}

func (s *DashboardService) AssistantReport(rangeType string) (*model.AssistantReport, error) {
	return s.repo.FetchAssistantReport(rangeType)
}
