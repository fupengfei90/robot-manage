import http from './http'
import type {
  DashboardSummary,
  KnowledgeSnapshot,
  TimelineInfo,
  AssistantReport
} from '../types/dashboard'
import type { ApiResponse } from '../types/api'

export const fetchDashboardSummary = async (): Promise<DashboardSummary> => {
  const res = await http.get<DashboardSummary>('/dashboard/summary')
  return res
}

export const fetchKnowledgeSnapshot = async (): Promise<KnowledgeSnapshot> => {
  const res = await http.get<KnowledgeSnapshot>('/dashboard/knowledge')
  return res
}

export const fetchTimelineInfo = async (): Promise<TimelineInfo> => {
  const res = await http.get<TimelineInfo>('/dashboard/timeline')
  return res
}

export const fetchAssistantReport = async (range = 'daily'): Promise<AssistantReport> => {
  const res = await http.get<AssistantReport>('/dashboard/assistant-report', {
    params: { range }
  })
  return res
}
