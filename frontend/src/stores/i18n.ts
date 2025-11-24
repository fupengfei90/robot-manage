import { defineStore } from 'pinia'

export type Locale = 'zh' | 'en'

interface Translations {
  [key: string]: string | Translations
}

const translations: Record<Locale, Translations> = {
  zh: {
    // Layout
    layout: {
      title: '小娇智能运维平台',
      subtitle: '一站式运维指挥中枢',
      dashboard: 'Dashboard',
      cmdb: 'CMDB',
      delivery: '智能交付',
      digital: '数字员工',
      system: '系统管理',
      toggleTheme: '切换主题',
      toggleLanguage: '切换语言',
      dutySchedule: '值班管理',
      milestone: '大事记'
    },
    // Dashboard
    dashboard: {
      metrics: {
        title: '运营指标',
        subtitle: '服务能力与巡检告警总览',
        realtime: '实时刷新',
        daily: '日服务次数',
        weekly: '周服务次数',
        monthly: '月服务次数',
        yearly: '年服务次数',
        users: '在册用户',
        inspectionTrend: '巡检趋势',
        alertLevel: '告警等级'
      },
      knowledge: {
        title: '知识库动态',
        recentUpdates: '最近更新',
        hotTopics: '热门知识',
        author: '作者',
        views: '阅读',
        recallRate: '召回'
      },
      timeline: {
        title: '值班与大事记',
        dutyOfficer: '今日值班',
        milestones: '大事记',
        deployPlans: '投产计划',
        system: '系统',
        window: '窗口',
        owner: '负责人',
        viewSchedule: '查看排班'
      },
      assistant: {
        title: '助手总结',
        range: '范围',
        highlights: '亮点',
        risks: '风险',
        nextSteps: '下一步',
        daily: '日报',
        weekly: '周报',
        monthly: '月报',
        yearly: '年报'
      },
      digitalHuman: {
        title: '数字人助手',
        speak: '语音播报'
      }
    }
  },
  en: {
    // Layout
    layout: {
      title: 'Intelligent Operations Platform',
      subtitle: 'Enterprise-Grade One-Stop Operations Command Center',
      dashboard: 'Dashboard',
      cmdb: 'CMDB',
      delivery: 'Smart Delivery',
      digital: 'Digital Employee',
      system: 'System Management',
      toggleTheme: 'Toggle Theme',
      toggleLanguage: 'Toggle Language',
      dutySchedule: 'Duty Schedule',
      milestone: 'Milestones'
    },
    // Dashboard
    dashboard: {
      metrics: {
        title: 'Operations Metrics',
        subtitle: 'Service Capability & Inspection Alert Overview',
        realtime: 'Real-time Refresh',
        daily: 'Daily Services',
        weekly: 'Weekly Services',
        monthly: 'Monthly Services',
        yearly: 'Yearly Services',
        users: 'Registered Users',
        inspectionTrend: 'Inspection Trend',
        alertLevel: 'Alert Level'
      },
      knowledge: {
        title: 'Knowledge Base Updates',
        recentUpdates: 'Recent Updates',
        hotTopics: 'Hot Topics',
        author: 'Author',
        views: 'Views',
        recallRate: 'Recall Rate'
      },
      timeline: {
        title: 'Duty & Milestones',
        dutyOfficer: 'Today\'s Duty Officer',
        milestones: 'Milestones',
        deployPlans: 'Deployment Plans',
        system: 'System',
        window: 'Window',
        owner: 'Owner',
        viewSchedule: 'View Schedule'
      },
      assistant: {
        title: 'Assistant Summary',
        range: 'Range',
        highlights: 'Highlights',
        risks: 'Risks',
        nextSteps: 'Next Steps',
        daily: 'Daily',
        weekly: 'Weekly',
        monthly: 'Monthly',
        yearly: 'Yearly'
      },
      digitalHuman: {
        title: 'Digital Assistant',
        speak: 'Voice Broadcast'
      }
    }
  }
}

export const useI18nStore = defineStore('i18n', {
  state: (): { locale: Locale } => ({
    locale: (localStorage.getItem('locale') as Locale) || 'zh'
  }),
  getters: {
    t: (state) => {
      return (key: string): string => {
        const keys = key.split('.')
        let value: any = translations[state.locale]
        for (const k of keys) {
          value = value?.[k]
        }
        return typeof value === 'string' ? value : key
      }
    }
  },
  actions: {
    setLocale(locale: Locale) {
      this.locale = locale
      localStorage.setItem('locale', locale)
    },
    toggle() {
      this.setLocale(this.locale === 'zh' ? 'en' : 'zh')
    }
  }
})

