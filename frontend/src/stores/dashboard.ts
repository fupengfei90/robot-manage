import { defineStore } from 'pinia'
import { ElMessage } from 'element-plus'
import {
  fetchAssistantReport,
  fetchDashboardSummary,
  fetchKnowledgeSnapshot,
  fetchTimelineInfo
} from '../api/dashboard'
import type { DashboardState } from '../types/dashboard'

export const useDashboardStore = defineStore('dashboard', {
  state: (): DashboardState => ({
    summary: null,
    knowledge: null,
    timeline: null,
    assistant: null,
    loading: false
  }),
  actions: {
    async initialize(range = 'daily') {
      this.loading = true
      try {
        const [summary, knowledge, timeline, assistant] = await Promise.all([
          fetchDashboardSummary(),
          fetchKnowledgeSnapshot(),
          fetchTimelineInfo(),
          fetchAssistantReport(range)
        ])

        this.summary = summary
        this.knowledge = knowledge
        this.timeline = timeline
        this.assistant = assistant
      } catch (error) {
        ElMessage.error('Dashboard 数据加载失败')
        throw error
      } finally {
        this.loading = false
      }
    }
  }
})
