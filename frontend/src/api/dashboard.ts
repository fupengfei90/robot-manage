import http from './http'
import type {
  DashboardSummary,
  KnowledgeSnapshot,
  TimelineInfo,
  AssistantReport
} from '../types/dashboard'
import type { ApiResponse } from '../types/api'

export const fetchDashboardSummary = async (): Promise<DashboardSummary> => {
  const res = await http.get<ApiResponse<DashboardSummary>, ApiResponse<DashboardSummary>>(
    '/dashboard/summary'
  )
  return res.data
}

export const fetchKnowledgeSnapshot = async (): Promise<KnowledgeSnapshot> => {
  const res = await http.get<ApiResponse<KnowledgeSnapshot>, ApiResponse<KnowledgeSnapshot>>(
    '/dashboard/knowledge'
  )
  return res.data
}

export const fetchTimelineInfo = async (): Promise<TimelineInfo> => {
  const res = await http.get<ApiResponse<TimelineInfo>, ApiResponse<TimelineInfo>>(
    '/dashboard/timeline'
  )
  return res.data
}

export const fetchAssistantReport = async (range = 'daily'): Promise<AssistantReport> => {
  const res = await http.get<ApiResponse<AssistantReport>, ApiResponse<AssistantReport>>(
    '/dashboard/assistant-report',
    {
      params: { range }
    }
  )
  return res.data
}
