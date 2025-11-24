export interface MetricBreakdown {
  total: number
  trend: number[]
  labels: string[]
  detail: Record<string, number>
}

export interface DashboardSummary {
  serviceCounts: Record<string, number>
  userCount: number
  serviceUserCount: number
  inspection: MetricBreakdown
  alerts: MetricBreakdown
  serviceTrend: MetricBreakdown
}

export interface KnowledgeItem {
  title: string
  author: string
  updatedAt: string
  views: number
}

export interface KnowledgeSnapshot {
  recentUpdates: KnowledgeItem[]
  hotTopics: KnowledgeItem[]
  recallRate: number
}

export interface Milestone {
  time: string
  title: string
  level: string
}

export interface DeployPlan {
  system: string
  window: string
  owner: string
}

export interface TimelineInfo {
  dutyOfficer: string
  date: string
  milestones: Milestone[]
  deployPlans: DeployPlan[]
}

export interface AssistantReport {
  range: string
  highlights: string[]
  risks: string[]
  nextSteps: string[]
}

export interface DashboardState {
  summary: DashboardSummary | null
  knowledge: KnowledgeSnapshot | null
  timeline: TimelineInfo | null
  assistant: AssistantReport | null
  loading: boolean
}
