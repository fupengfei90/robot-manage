package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

type ScheduleTaskService struct {
	repo *repository.ScheduleTaskRepository
}

func NewScheduleTaskService(repo *repository.ScheduleTaskRepository) *ScheduleTaskService {
	return &ScheduleTaskService{repo: repo}
}

// GetList 获取任务列表
func (s *ScheduleTaskService) GetList(query *model.ScheduleTaskQuery) ([]*model.ScheduleTaskResponse, int64, error) {
	tasks, total, err := s.repo.GetList(query)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.ScheduleTaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = s.buildTaskResponse(task)
	}

	return responses, total, nil
}

// GetByID 根据ID获取任务
func (s *ScheduleTaskService) GetByID(id uint) (*model.ScheduleTaskResponse, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.buildTaskResponse(task), nil
}

// Create 创建任务
func (s *ScheduleTaskService) Create(req *model.ScheduleTaskRequest) (*model.ScheduleTaskResponse, error) {
	// 验证cron表达式
	if err := s.validateCron(req.Cron); err != nil {
		return nil, fmt.Errorf("cron表达式无效: %v", err)
	}

	// 检查任务名称是否重复
	exists, err := s.repo.CheckNameExists(req.Name, req.DCN, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("任务名称在数据中心 %s 中已存在", req.DCN)
	}

	task := &model.ScheduleTask{
		Name:      req.Name,
		Active:    req.Active,
		Cron:      req.Cron,
		Workflow:  req.Workflow,
		ExecToken: req.ExecToken,
		Category:  req.Category,
		DCN:       req.DCN,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return s.GetByID(task.ID)
}

// Update 更新任务
func (s *ScheduleTaskService) Update(id uint, req *model.ScheduleTaskRequest) (*model.ScheduleTaskResponse, error) {
	// 验证cron表达式
	if err := s.validateCron(req.Cron); err != nil {
		return nil, fmt.Errorf("cron表达式无效: %v", err)
	}

	// 检查任务名称是否重复
	exists, err := s.repo.CheckNameExists(req.Name, req.DCN, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("任务名称在数据中心 %s 中已存在", req.DCN)
	}

	task := &model.ScheduleTask{
		Name:      req.Name,
		Active:    req.Active,
		Cron:      req.Cron,
		Workflow:  req.Workflow,
		ExecToken: req.ExecToken,
		Category:  req.Category,
		DCN:       req.DCN,
	}

	if err := s.repo.Update(id, task); err != nil {
		return nil, err
	}

	return s.GetByID(id)
}

// Delete 删除任务
func (s *ScheduleTaskService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// UpdateStatus 更新任务状态
func (s *ScheduleTaskService) UpdateStatus(id uint, active int) error {
	return s.repo.UpdateStatus(id, active)
}

// BatchOperation 批量操作
func (s *ScheduleTaskService) BatchOperation(req *model.BatchOperationRequest) error {
	switch req.Operation {
	case "enable":
		return s.repo.BatchUpdateStatus(req.IDs, 1)
	case "disable":
		return s.repo.BatchUpdateStatus(req.IDs, 0)
	case "delete":
		return s.repo.BatchDelete(req.IDs)
	default:
		return fmt.Errorf("不支持的操作: %s", req.Operation)
	}
}

// GetCategories 获取所有分类
func (s *ScheduleTaskService) GetCategories() ([]string, error) {
	return s.repo.GetCategories()
}

// GetDCNs 获取所有数据中心
func (s *ScheduleTaskService) GetDCNs() ([]string, error) {
	return s.repo.GetDCNs()
}

// ExecuteTask 立即执行任务
func (s *ScheduleTaskService) ExecuteTask(id uint) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if task.Active == 0 {
		return fmt.Errorf("任务已停用，无法执行")
	}

	// TODO: 实现实际的任务执行逻辑
	// 这里可以调用工作流执行接口
	fmt.Printf("执行任务: %s (工作流: %s)\n", task.Name, task.Workflow)
	
	return nil
}

// buildTaskResponse 构建任务响应
func (s *ScheduleTaskService) buildTaskResponse(task *model.ScheduleTask) *model.ScheduleTaskResponse {
	response := &model.ScheduleTaskResponse{
		ScheduleTask: task,
		CronDesc:     s.parseCronDescription(task.Cron),
		StatusText:   s.getStatusText(task.Active),
	}

	// 计算下次执行时间
	if task.Active == 1 {
		if nextTime := s.getNextRunTime(task.Cron); nextTime != nil {
			timeStr := nextTime.Format("2006-01-02 15:04:05")
			response.NextRunTime = &timeStr
		}
	}

	return response
}

// validateCron 验证cron表达式
func (s *ScheduleTaskService) validateCron(cronExpr string) error {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	_, err := parser.Parse(cronExpr)
	return err
}

// getNextRunTime 获取下次执行时间
func (s *ScheduleTaskService) getNextRunTime(cronExpr string) *time.Time {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := parser.Parse(cronExpr)
	if err != nil {
		return nil
	}
	
	nextTime := schedule.Next(time.Now())
	return &nextTime
}

// getStatusText 获取状态文本
func (s *ScheduleTaskService) getStatusText(active int) string {
	if active == 1 {
		return "启用"
	}
	return "停用"
}

// parseCronDescription 解析cron表达式描述
func (s *ScheduleTaskService) parseCronDescription(cronExpr string) string {
	parts := strings.Fields(cronExpr)
	if len(parts) != 6 {
		return "无效的cron表达式"
	}

	second, minute, hour, day, _, weekday := parts[0], parts[1], parts[2], parts[3], parts[4], parts[5]

	// 简单的cron描述解析
	desc := ""

	// 处理秒
	if second == "0" {
		// 整分钟执行
	} else if second == "*" {
		desc += "每秒"
	} else {
		desc += fmt.Sprintf("第%s秒", second)
	}

	// 处理分钟
	if minute == "*" {
		if desc != "" {
			desc += "每分钟"
		}
	} else if minute == "0" {
		// 整点执行
	} else {
		if desc != "" {
			desc += " "
		}
		desc += fmt.Sprintf("%s分", minute)
	}

	// 处理小时
	if hour == "*" {
		if desc != "" {
			desc += "每小时"
		}
	} else {
		if desc != "" {
			desc += " "
		}
		desc += fmt.Sprintf("%s点", hour)
	}

	// 处理星期
	if weekday != "*" {
		weekdayMap := map[string]string{
			"0": "周日", "1": "周一", "2": "周二", "3": "周三",
			"4": "周四", "5": "周五", "6": "周六",
		}
		
		if strings.Contains(weekday, ",") {
			days := strings.Split(weekday, ",")
			dayNames := make([]string, len(days))
			for i, d := range days {
				if name, ok := weekdayMap[d]; ok {
					dayNames[i] = name
				} else {
					dayNames[i] = d
				}
			}
			desc += " " + strings.Join(dayNames, ",")
		} else if name, ok := weekdayMap[weekday]; ok {
			desc += " " + name
		}
	}

	// 处理日期
	if day != "*" && weekday == "*" {
		desc += fmt.Sprintf(" %s日", day)
	}

	if desc == "" {
		desc = "自定义时间"
	}

	return strings.TrimSpace(desc)
}