package repository

import (
	"encoding/json"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/logger"
	"github.com/example/robot-manage/internal/model"
)

// DashboardRepository 定义数据访问接口，可落地至GORM或其他存储。
type DashboardRepository struct {
	db *gorm.DB
}

// NewDashboardRepository 构造函数。
func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

// FetchSummary 查询概览指标，若数据库不可用则返回模拟数据。
func (r *DashboardRepository) FetchSummary() (*model.DashboardSummary, error) {
	if r.db == nil {
		return mockSummary(), nil
	}

	summary := &model.DashboardSummary{
		ServiceCounts: map[string]int{"daily": 0, "weekly": 0, "monthly": 0, "yearly": 0},
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		type serviceStat struct {
			Period string
			Count  int
		}
		var stats []serviceStat
		if err := tx.Table("ops_service_stats").Select("period,count").Where("metric = ?", "service_count").Find(&stats).Error; err != nil {
			return err
		}
		for _, stat := range stats {
			summary.ServiceCounts[strings.ToLower(stat.Period)] = stat.Count
		}

		if err := tx.Table("ops_user_profile").Select("COUNT(1)").Scan(&summary.UserCount).Error; err != nil {
			return err
		}

		summary.Inspection.Detail = map[string]int{}
		type metricRow struct {
			Label   string
			Value   int
			Segment string
		}
		var inspectionRows []metricRow
		if err := tx.Table("ops_inspection_trend").Select("label,value,segment").Order("sequence asc").Find(&inspectionRows).Error; err != nil {
			return err
		}
		for _, row := range inspectionRows {
			summary.Inspection.Trend = append(summary.Inspection.Trend, row.Value)
			summary.Inspection.Labels = append(summary.Inspection.Labels, row.Label)
			summary.Inspection.Total += row.Value
			if row.Segment != "" {
				summary.Inspection.Detail[row.Segment] += row.Value
			}
		}

		summary.Alerts.Detail = map[string]int{}
		var alertRows []metricRow
		if err := tx.Table("ops_alert_metrics").Select("label,value,segment").Order("severity asc").Find(&alertRows).Error; err != nil {
			return err
		}
		for _, row := range alertRows {
			summary.Alerts.Trend = append(summary.Alerts.Trend, row.Value)
			summary.Alerts.Labels = append(summary.Alerts.Labels, row.Label)
			summary.Alerts.Total += row.Value
			if row.Segment != "" {
				summary.Alerts.Detail[row.Segment] += row.Value
			}
		}

		return nil
	})

	if err != nil {
		logger.L().Warn("查询Dashboard概览失败，返回模拟数据", zap.Error(err))
		return mockSummary(), nil
	}

	return summary, nil
}

// FetchKnowledge 查询知识库动态。
func (r *DashboardRepository) FetchKnowledge() (*model.KnowledgeSnapshot, error) {
	if r.db == nil {
		return mockKnowledge(), nil
	}

	snapshot := &model.KnowledgeSnapshot{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		type knowledgeRow struct {
			Title     string
			Author    string
			UpdatedAt time.Time
			Views     int
		}
		var recents []knowledgeRow
		if err := tx.Table("ops_knowledge_articles").
			Select("title,author,updated_at,views").
			Order("updated_at desc").Limit(5).Find(&recents).Error; err != nil {
			return err
		}
		for _, row := range recents {
			snapshot.RecentUpdates = append(snapshot.RecentUpdates, model.KnowledgeItem{
				Title:     row.Title,
				Author:    row.Author,
				UpdatedAt: row.UpdatedAt.Format("2006-01-02"),
				Views:     row.Views,
			})
		}

		var hots []knowledgeRow
		if err := tx.Table("ops_knowledge_articles").
			Select("title,author,updated_at,views").
			Order("views desc").Limit(5).Find(&hots).Error; err != nil {
			return err
		}
		for _, row := range hots {
			snapshot.HotTopics = append(snapshot.HotTopics, model.KnowledgeItem{
				Title:     row.Title,
				Author:    row.Author,
				UpdatedAt: row.UpdatedAt.Format("2006-01-02"),
				Views:     row.Views,
			})
		}

		if err := tx.Table("ops_knowledge_metrics").Select("recall_rate").Order("snapshot_at desc").Limit(1).Scan(&snapshot.RecallRate).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.L().Warn("查询知识库动态失败，返回模拟数据", zap.Error(err))
		return mockKnowledge(), nil
	}

	return snapshot, nil
}

// FetchTimeline 查询日例数据。
func (r *DashboardRepository) FetchTimeline() (*model.TimelineInfo, error) {
	if r.db == nil {
		return mockTimeline(), nil
	}

	result := &model.TimelineInfo{
		Date:        time.Now().Format("2006-01-02"),
		Milestones:  []model.Milestone{},
		DeployPlans: []model.DeployPlan{},
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 查询今天的值班人员
		today := time.Now().Format("2006-01-02")
		type dutyRow struct {
			WBStaffName string
			FBStaffName string
		}
		var duty dutyRow
		if err := tx.Table("duty_schedule").
			Select("wb_staff_name, fb_staff_name").
			Where("duty_date = ?", today).
			First(&duty).Error; err == nil {
			// 组合WB和FB值班人员
			if duty.WBStaffName != "" && duty.FBStaffName != "" {
				result.DutyOfficer = duty.WBStaffName + " / " + duty.FBStaffName
			} else if duty.WBStaffName != "" {
				result.DutyOfficer = duty.WBStaffName
			} else if duty.FBStaffName != "" {
				result.DutyOfficer = duty.FBStaffName
			} else {
				result.DutyOfficer = "未安排"
			}
		} else {
			result.DutyOfficer = "未安排"
		}

		// 查询最近的大事记（只显示激活的，最近30天）
		thirtyDaysAgo := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
		type milestoneRow struct {
			EventDate    time.Time
			EventContent string
		}
		var milestones []milestoneRow
		if err := tx.Table("milestone_events").
			Select("event_date, event_content").
			Where("is_active = ? AND event_date >= ?", true, thirtyDaysAgo).
			Order("event_date desc, id desc").
			Limit(20).
			Find(&milestones).Error; err == nil {
			for _, m := range milestones {
				// 过滤空内容和无效内容
				content := strings.TrimSpace(m.EventContent)
				if content == "" || content == "--" || content == "-" {
					continue
				}
				// 格式化日期为 MM-DD
				dateStr := m.EventDate.Format("01-02")
				result.Milestones = append(result.Milestones, model.Milestone{
					Time:  dateStr,
					Title: content,
					Level: "info",
				})
				// 最多显示10条有效的大事记
				if len(result.Milestones) >= 10 {
					break
				}
			}
		}

		// 投产计划暂时保留模拟数据，因为数据库中没有对应表
		result.DeployPlans = []model.DeployPlan{
			{System: "会员中心", Window: "22:00-23:00", Owner: "待定"},
			{System: "风控引擎", Window: "23:30-00:30", Owner: "待定"},
		}

		return nil
	})

	if err != nil {
		logger.L().Warn("查询日例失败，返回模拟数据", zap.Error(err))
		return mockTimeline(), nil
	}

	// 如果没有查询到数据，使用默认值
	if result.DutyOfficer == "" {
		result.DutyOfficer = "未安排"
	}
	if len(result.Milestones) == 0 {
		result.Milestones = []model.Milestone{
			{Time: "今日", Title: "暂无大事记", Level: "info"},
		}
	}

	return result, nil
}

// FetchAssistantReport 生成助手总结。
func (r *DashboardRepository) FetchAssistantReport(rangeType string) (*model.AssistantReport, error) {
	if r.db == nil {
		return mockAssistant(rangeType), nil
	}

	report := &model.AssistantReport{Range: rangeType}

	type assistantRow struct {
		Highlights string
		Risks      string
		NextSteps  string
	}
	var row assistantRow
	err := r.db.Table("ops_assistant_reports").
		Select("highlights,risks,next_steps").
		Where("range = ?", rangeType).
		Order("generated_at desc").
		Limit(1).
		Scan(&row).Error
	if err != nil {
		logger.L().Warn("查询助手总结失败，返回模拟数据", zap.Error(err))
		return mockAssistant(rangeType), nil
	}

	report.Highlights = parseStringArray(row.Highlights)
	report.Risks = parseStringArray(row.Risks)
	report.NextSteps = parseStringArray(row.NextSteps)

	return report, nil
}

func mockSummary() *model.DashboardSummary {
	return &model.DashboardSummary{
		ServiceCounts: map[string]int{
			"daily":   120,
			"weekly":  560,
			"monthly": 2100,
			"yearly":  14500,
		},
		UserCount: 86,
		Inspection: model.MetricBreakdown{
			Total:  320,
			Trend:  []int{12, 18, 22, 25, 21, 30, 35},
			Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
			Detail: map[string]int{"自动": 260, "人工": 60},
		},
		Alerts: model.MetricBreakdown{
			Total:  42,
			Trend:  []int{5, 9, 8, 6, 7, 4, 3},
			Labels: []string{"P1", "P2", "P3", "P4"},
			Detail: map[string]int{"P1": 3, "P2": 11, "P3": 20, "P4": 8},
		},
	}
}

func mockKnowledge() *model.KnowledgeSnapshot {
	return &model.KnowledgeSnapshot{
		RecentUpdates: []model.KnowledgeItem{
			{Title: "K8s集群容量评估", Author: "OpsBot", UpdatedAt: "2025-11-18", Views: 97},
			{Title: "零信任接入排障指南", Author: "SecTeam", UpdatedAt: "2025-11-17", Views: 142},
		},
		HotTopics: []model.KnowledgeItem{
			{Title: "数据库巡检模版", Author: "DBA", UpdatedAt: "2025-10-30", Views: 260},
			{Title: "AI助手训练指北", Author: "AI团队", UpdatedAt: "2025-11-01", Views: 211},
		},
		RecallRate: 0.87,
	}
}

func mockTimeline() *model.TimelineInfo {
	return &model.TimelineInfo{
		DutyOfficer: "李明",
		Date:        "2025-11-20",
		Milestones: []model.Milestone{
			{Time: "09:00", Title: "智能巡检完成", Level: "info"},
			{Time: "13:30", Title: "支付系统版本回归", Level: "warning"},
			{Time: "18:00", Title: "晚间值班交接", Level: "info"},
		},
		DeployPlans: []model.DeployPlan{
			{System: "会员中心", Window: "22:00-23:00", Owner: "孙嘉"},
			{System: "风控引擎", Window: "23:30-00:30", Owner: "赵衡"},
		},
	}
}

func mockAssistant(rangeType string) *model.AssistantReport {
	return &model.AssistantReport{
		Range: rangeType,
		Highlights: []string{
			"AI巡检覆盖率达到 93% 并自动闭环 24 条告警",
			"智能提单平均耗时缩短至 18 秒",
		},
		Risks: []string{
			"东区机房温湿度波动需持续监控",
		},
		NextSteps: []string{
			"完成Hybrid云备份演练",
			"发布RPA数字员工2.0训练计划",
		},
	}
}

func parseStringArray(payload string) []string {
	if payload == "" {
		return []string{}
	}
	var data []string
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return []string{}
	}
	return data
}
