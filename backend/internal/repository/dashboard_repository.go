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
		// 使用真实数据统计服务次数（包括定时任务，排除无效数据）
		now := time.Now()
		today := now.Format("2006-01-02")
		weekAgo := now.AddDate(0, 0, -6).Format("2006-01-02") // 修正：最近7天（包括今天）
		monthAgo := now.AddDate(0, -1, 0).Format("2006-01-02")
		yearAgo := now.AddDate(-1, 0, 0).Format("2006-01-02")

		var dailyCount, weeklyCount, monthlyCount, yearlyCount int64
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) = ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", today).Count(&dailyCount)
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", weekAgo).Count(&weeklyCount)
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", monthAgo).Count(&monthlyCount)
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", yearAgo).Count(&yearlyCount)

		summary.ServiceCounts["daily"] = int(dailyCount)
		summary.ServiceCounts["weekly"] = int(weeklyCount)
		summary.ServiceCounts["monthly"] = int(monthlyCount)
		summary.ServiceCounts["yearly"] = int(yearlyCount)

		// 统计最近7天的用户服务和定时任务分类
		var weeklyUserService, weeklyScheduledTask int64
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND (type IS NULL OR type != 'scheduled_task') AND (is_valid = 1 OR is_valid IS NULL)", "test-%", weekAgo).Count(&weeklyUserService)
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND type = 'scheduled_task' AND (is_valid = 1 OR is_valid IS NULL)", "test-%", weekAgo).Count(&weeklyScheduledTask)
		summary.ServiceCounts["weekly_user_service"] = int(weeklyUserService)
		summary.ServiceCounts["weekly_scheduled_task"] = int(weeklyScheduledTask)

		// 统计最近7天的提单操作次数
		var weeklyTicketCount int64
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND output LIKE ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", weekAgo, "%【待确认】请回复【确认】后开始自动生成提单信息%").Count(&weeklyTicketCount)
		summary.ServiceCounts["weekly_ticket_count"] = int(weeklyTicketCount)

		// 在册用户数统计，从wecom_users表获取
		var userCount int64
		if err := tx.Table("wecom_users").Where("deleted_at IS NULL").Count(&userCount).Error; err == nil {
			summary.UserCount = int(userCount)
		}

		// 服务用户数统计，从会话记录表获取去重用户数
		var serviceUserCount int64
		if err := tx.Table("wecom_message_records").Where("msg_id NOT LIKE ?", "test-%").Distinct("request_user_id").Count(&serviceUserCount).Error; err == nil {
			summary.ServiceUserCount = int(serviceUserCount)
		}

		// 服务次数趋势统计（最近30天，只统计有效数据）
		summary.ServiceTrend.Detail = map[string]int{}
		type serviceTrendRow struct {
			Date      string
			NonTask   int
			Scheduled int
		}
		var trendRows []serviceTrendRow
		if err := tx.Raw(`
			SELECT 
				DATE(created_at) as date,
				SUM(CASE WHEN (type IS NULL OR type != 'scheduled_task') AND (is_valid = 1 OR is_valid IS NULL) THEN 1 ELSE 0 END) as non_task,
				SUM(CASE WHEN type = 'scheduled_task' AND (is_valid = 1 OR is_valid IS NULL) THEN 1 ELSE 0 END) as scheduled
			FROM wecom_message_records
			WHERE msg_id NOT LIKE 'test-%' AND created_at >= DATE_SUB(NOW(), INTERVAL 30 DAY)
			GROUP BY DATE(created_at)
			ORDER BY date ASC
		`).Scan(&trendRows).Error; err != nil {
			logger.L().Warn("查询服务次数趋势失败", zap.Error(err))
		}
		logger.L().Info("服务次数趋势查询结果", zap.Int("rows", len(trendRows)))
		for _, row := range trendRows {
			// 格式化日期为 MM-DD
			if len(row.Date) >= 10 {
				label := row.Date[5:10]
				total := row.NonTask + row.Scheduled
				summary.ServiceTrend.Labels = append(summary.ServiceTrend.Labels, label)
				summary.ServiceTrend.Trend = append(summary.ServiceTrend.Trend, total)
				summary.ServiceTrend.Total += total
				summary.ServiceTrend.Detail["用户服务_"+label] = row.NonTask
				summary.ServiceTrend.Detail["定时任务_"+label] = row.Scheduled
			}
		}

		summary.Inspection.Detail = map[string]int{}
		type metricRow struct {
			Label   string
			Value   int
			Segment string
		}
		// 巡检数据，失败不影响服务次数
		var inspectionRows []metricRow
		tx.Table("ops_inspection_trend").Select("label,value,segment").Order("sequence asc").Find(&inspectionRows)
		for _, row := range inspectionRows {
			summary.Inspection.Trend = append(summary.Inspection.Trend, row.Value)
			summary.Inspection.Labels = append(summary.Inspection.Labels, row.Label)
			summary.Inspection.Total += row.Value
			if row.Segment != "" {
				summary.Inspection.Detail[row.Segment] += row.Value
			}
		}

		// 告警数据，失败不影响服务次数
		summary.Alerts.Detail = map[string]int{}
		var alertRows []metricRow
		tx.Table("ops_alert_metrics").Select("label,value,segment").Order("severity asc").Find(&alertRows)
		for _, row := range alertRows {
			summary.Alerts.Trend = append(summary.Alerts.Trend, row.Value)
			summary.Alerts.Labels = append(summary.Alerts.Labels, row.Label)
			summary.Alerts.Total += row.Value
			if row.Segment != "" {
				summary.Alerts.Detail[row.Segment] += row.Value
			}
		}

		// 填充周报所需字段
		summary.WeeklyServiceCount = summary.ServiceCounts["weekly"]
		// 统计本周的覆盖用户数
		var weeklyServiceUsers int64
		tx.Table("wecom_message_records").Where("msg_id NOT LIKE ? AND DATE(created_at) >= ? AND (is_valid = 1 OR is_valid IS NULL)", "test-%", weekAgo).Distinct("request_user_id").Count(&weeklyServiceUsers)
		summary.ServiceUsers = int(weeklyServiceUsers)
		// 统计最近7天的巡检次数
		var weeklyInspectionCount int64
		tx.Table("ops_inspection_trend").Where("DATE(created_at) >= ?", weekAgo).Select("COALESCE(SUM(value), 0)").Scan(&weeklyInspectionCount)
		summary.WeeklyInspectionCount = int(weeklyInspectionCount)

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
			{System: "Product-B-20251120", Window: "用户风险等级问题修复", Owner: "ruoxinhuang, harrischen"},
			{System: "Product-B-20251127", Window: "对公企业网银限额清零问题修复", Owner: "yuanlinbao, fulali"},
			{System: "Product-O-20251125", Window: "修复REDIS安全漏洞组件", Owner: "yolandawang"},
			{System: "Product-O-20251127", Window: "资产信息管理需求", Owner: "v_alipeng, borfyqiu"},
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
		UserCount:             86,
		ServiceUserCount:      52,
		WeeklyServiceCount:    560,
		ServiceUsers:          52,
		WeeklyInspectionCount: 163,
		ServiceTrend: model.MetricBreakdown{
			Total:  520,
			Trend:  []int{65, 72, 88, 91, 78, 83, 83},
			Labels: []string{"11-18", "11-19", "11-20", "11-21", "11-22", "11-23", "11-24"},
			Detail: map[string]int{"用户服务": 420, "定时任务": 100},
		},
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
			{Title: "IDP批量异常SOP", Author: "pengfeifu", UpdatedAt: "2025-11-28", Views: 97},
			{Title: "CMDB知识库更新", Author: "交付团队", UpdatedAt: "2025-11-27", Views: 142},
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
			{System: "Product-B-20251120", Window: "用户风险等级问题修复", Owner: "ruoxinhuang, harrischen"},
			{System: "Product-B-20251127", Window: "对公企业网银限额清零问题修复", Owner: "yuanlinbao, fulali"},
			{System: "Product-O-20251125", Window: "修复REDIS安全漏洞组件", Owner: "yolandawang"},
			{System: "Product-O-20251127", Window: "资产信息管理需求", Owner: "v_alipeng, borfyqiu"},
		},
	}
}

func mockAssistant(rangeType string) *model.AssistantReport {
	return &model.AssistantReport{
		Range: rangeType,
		Highlights: []string{
			"日常巡检和服务数据上报稳定运行，11月5日成功上报约300条数据",
			"完成富小娇能力汇报文档修改和视频反馈材料准备",
		},
		Risks: []string{
			"服务次数数据异常问题需持续关注和优化",
		},
		NextSteps: []string{
			"推进慢SQL分析功能开发",
			"扩展富小娇能力应用到更多团队和场景",
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
