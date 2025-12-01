package model

// DashboardSummary 概览指标。
type DashboardSummary struct {
	ServiceCounts         map[string]int  `json:"serviceCounts"`
	UserCount             int             `json:"userCount"`
	ServiceUserCount      int             `json:"serviceUserCount"`
	Inspection            MetricBreakdown `json:"inspection"`
	Alerts                MetricBreakdown `json:"alerts"`
	ServiceTrend          MetricBreakdown `json:"serviceTrend"`
	WeeklyServiceCount    int             `json:"weeklyServiceCount"`
	ServiceUsers          int             `json:"serviceUsers"`
	WeeklyInspectionCount int             `json:"weeklyInspectionCount"`
}

// MetricBreakdown 细分指标。
type MetricBreakdown struct {
	Total  int            `json:"total"`
	Trend  []int          `json:"trend"`
	Labels []string       `json:"labels"`
	Detail map[string]int `json:"detail"`
}

// KnowledgeSnapshot 知识库动态。
type KnowledgeSnapshot struct {
	RecentUpdates []KnowledgeItem `json:"recentUpdates"`
	HotTopics     []KnowledgeItem `json:"hotTopics"`
	RecallRate    float64         `json:"recallRate"`
}

// KnowledgeItem 单条知识信息。
type KnowledgeItem struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	UpdatedAt string `json:"updatedAt"`
	Views     int    `json:"views"`
}

// TimelineInfo 日例视图。
type TimelineInfo struct {
	DutyOfficer string       `json:"dutyOfficer"`
	Date        string       `json:"date"`
	Milestones  []Milestone  `json:"milestones"`
	DeployPlans []DeployPlan `json:"deployPlans"`
}

// Milestone 事件节点。
type Milestone struct {
	Time  string `json:"time"`
	Title string `json:"title"`
	Level string `json:"level"`
}

// DeployPlan 投产计划。
type DeployPlan struct {
	System string `json:"system"`
	Window string `json:"window"`
	Owner  string `json:"owner"`
}

// AssistantReport 助手总结。
type AssistantReport struct {
	Range      string   `json:"range"`
	Highlights []string `json:"highlights"`
	Risks      []string `json:"risks"`
	NextSteps  []string `json:"nextSteps"`
}
