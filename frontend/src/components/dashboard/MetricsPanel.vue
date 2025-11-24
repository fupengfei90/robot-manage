<template>
  <el-card shadow="never" class="panel card-hover glass-effect">
    <div class="panel__header">
      <div class="header__content">
        <div class="header__icon">📊</div>
        <div>
          <h2 class="panel__title">{{ t('dashboard.metrics.title') }}</h2>
          <p class="panel__subtitle">{{ t('dashboard.metrics.subtitle') }}</p>
        </div>
      </div>
      <el-tag type="info" class="realtime-tag animate-pulse">
        <span class="tag__dot"></span>
        {{ t('dashboard.metrics.realtime') }}
      </el-tag>
    </div>

    <div class="panel__metrics">
      <el-tooltip
        v-for="key in serviceOrder" 
        :key="key"
        :content="getTimeRange(key)"
        placement="top"
      >
        <div class="metric card-hover">
          <div class="metric__icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M13 2L3 14H12L11 22L21 10H12L13 2Z" stroke="currentColor" stroke-width="2" fill="none"/>
            </svg>
          </div>
          <div class="metric__content">
            <p class="metric__label">{{ serviceDict[key as keyof typeof serviceDict] }}</p>
            <p class="metric__value">{{ summary.serviceCounts[key]?.toLocaleString() || 0 }}</p>
          </div>
        </div>
      </el-tooltip>
      <div class="metric metric--highlight card-hover">
        <div class="metric__icon metric__icon--highlight">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none"/>
            <path d="M12 6V12L16 14" stroke="currentColor" stroke-width="2" fill="none"/>
          </svg>
        </div>
        <div class="metric__content">
          <p class="metric__label">在册用户数</p>
          <p class="metric__value">{{ summary.userCount.toLocaleString() }}</p>
        </div>
      </div>
      <div class="metric metric--highlight card-hover">
        <div class="metric__icon metric__icon--highlight">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M17 21V19C17 17.9391 16.5786 16.9217 15.8284 16.1716C15.0783 15.4214 14.0609 15 13 15H5C3.93913 15 2.92172 15.4214 2.17157 16.1716C1.42143 16.9217 1 17.9391 1 19V21" stroke="currentColor" stroke-width="2" fill="none"/>
            <circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2" fill="none"/>
            <path d="M23 21V19C22.9993 18.1137 22.7044 17.2528 22.1614 16.5523C21.6184 15.8519 20.8581 15.3516 20 15.13" stroke="currentColor" stroke-width="2" fill="none"/>
            <path d="M16 3.13C16.8604 3.35031 17.623 3.85071 18.1676 4.55232C18.7122 5.25392 19.0078 6.11683 19.0078 7.005C19.0078 7.89318 18.7122 8.75608 18.1676 9.45769C17.623 10.1593 16.8604 10.6597 16 10.88" stroke="currentColor" stroke-width="2" fill="none"/>
          </svg>
        </div>
        <div class="metric__content">
          <p class="metric__label">服务用户数</p>
          <p class="metric__value">{{ summary.serviceUserCount.toLocaleString() }}</p>
        </div>
      </div>
    </div>

    <div class="panel__charts">
      <div class="chart-wrapper chart-wrapper--wide">
        <div class="chart-header">
          <el-radio-group v-model="trendDimension" size="small" @change="updateServiceTrendChart">
            <el-radio-button label="day">天</el-radio-button>
            <el-radio-button label="week">周</el-radio-button>
          </el-radio-group>
        </div>
        <div ref="serviceTrendRef" class="chart chart--service"></div>
        <div class="chart-summary">
          <div class="summary-title">{{ trendDimension === 'day' ? '本周汇总' : '本月汇总' }}</div>
          <div class="summary-item">
            <span class="summary-label">用户服务</span>
            <span class="summary-value">{{ currentPeriodSummary.userService.toLocaleString() }}</span>
          </div>
          <div class="summary-item">
            <span class="summary-label">定时任务</span>
            <span class="summary-value">{{ currentPeriodSummary.scheduledTask.toLocaleString() }}</span>
          </div>
        </div>
      </div>
      <div ref="inspectionRef" class="chart chart--inspection"></div>
      <div ref="alertsRef" class="chart chart--alerts"></div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref, watch, computed } from 'vue'
import * as echarts from 'echarts'
import type { DashboardSummary } from '../../types/dashboard'
import { useThemeStore } from '../../stores/theme'
import { useI18n } from '../../composables/useI18n'

const props = defineProps<{ summary: DashboardSummary }>()
const themeStore = useThemeStore()
const { t } = useI18n()

const serviceTrendRef = ref<HTMLDivElement | null>(null)
const inspectionRef = ref<HTMLDivElement | null>(null)
const alertsRef = ref<HTMLDivElement | null>(null)
const trendDimension = ref<'day' | 'week'>('day')
const currentPeriodSummary = computed(() => {
  if (!props.summary) return { userService: 0, scheduledTask: 0, total: 0 }
  
  const { serviceTrend, serviceCounts } = props.summary
  let userService = 0
  let scheduledTask = 0
  
  if (trendDimension.value === 'day') {
    // 天纬度：使用后端统计的weekly数据（最近7天）
    const total = serviceCounts.weekly || 0
    // 从最近7天的明细数据中计算分类
    const now = new Date()
    const last7Days: string[] = []
    for (let i = 6; i >= 0; i--) {
      const date = new Date(now)
      date.setDate(date.getDate() - i)
      const label = `${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
      last7Days.push(label)
    }
    last7Days.forEach(label => {
      userService += serviceTrend.detail[`用户服务_${label}`] || 0
      scheduledTask += serviceTrend.detail[`定时任务_${label}`] || 0
    })
  } else {
    // 周纬度：统计本月数据
    const currentMonth = new Date().getMonth() + 1
    serviceTrend.labels.forEach(label => {
      const [month] = label.split('-').map(Number)
      if (month === currentMonth) {
        userService += serviceTrend.detail[`用户服务_${label}`] || 0
        scheduledTask += serviceTrend.detail[`定时任务_${label}`] || 0
      }
    })
  }
  
  return { userService, scheduledTask, total: userService + scheduledTask }
})

let serviceTrendChart: echarts.ECharts | null = null
let inspectionChart: echarts.ECharts | null = null
let alertsChart: echarts.ECharts | null = null

const textColor = computed(() => {
  return themeStore.mode === 'dark' ? '#e2e8f0' : '#1e293b'
})

const axisLabelColor = computed(() => {
  return themeStore.mode === 'dark' ? '#94a3b8' : '#64748b'
})

const serviceOrder = ['daily', 'weekly', 'monthly', 'yearly']

const serviceDict = computed(() => ({
  daily: t('dashboard.metrics.daily'),
  weekly: t('dashboard.metrics.weekly'),
  monthly: t('dashboard.metrics.monthly'),
  yearly: t('dashboard.metrics.yearly')
}))

const getTimeRange = (key: string) => {
  const now = new Date()
  const formatDate = (date: Date) => {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    return `${year}-${month}-${day}`
  }
  
  const today = formatDate(now)
  
  switch(key) {
    case 'daily':
      return `统计时间: ${today}`
    case 'weekly': {
      const weekAgo = new Date(now)
      weekAgo.setDate(weekAgo.getDate() - 7)
      return `统计时间: ${formatDate(weekAgo)} ~ ${today}`
    }
    case 'monthly': {
      const monthAgo = new Date(now)
      monthAgo.setMonth(monthAgo.getMonth() - 1)
      return `统计时间: ${formatDate(monthAgo)} ~ ${today}`
    }
    case 'yearly': {
      const yearAgo = new Date(now)
      yearAgo.setFullYear(yearAgo.getFullYear() - 1)
      return `统计时间: ${formatDate(yearAgo)} ~ ${today}`
    }
    default:
      return ''
  }
}

const initCharts = () => {
  if (serviceTrendRef.value) {
    serviceTrendChart = echarts.init(serviceTrendRef.value)
  }
  if (inspectionRef.value) {
    inspectionChart = echarts.init(inspectionRef.value)
  }
  if (alertsRef.value) {
    alertsChart = echarts.init(alertsRef.value)
  }
  updateCharts()
}

const handleResize = () => {
  serviceTrendChart?.resize()
  inspectionChart?.resize()
  alertsChart?.resize()
}

const updateServiceTrendChart = () => {
  if (!props.summary) return
  const { serviceTrend } = props.summary
  
  let labels = serviceTrend.labels
  let trend = serviceTrend.trend
  let detailData: Record<string, number> = {}
  
  // 如果是周纬度，需要聚合数据
  if (trendDimension.value === 'week') {
    const weekData: Record<string, { total: number, userService: number, scheduledTask: number }> = {}
    
    serviceTrend.labels.forEach((label, index) => {
      // 解析日期 MM-DD
      const [month, day] = label.split('-').map(Number)
      const year = new Date().getFullYear()
      const date = new Date(year, month - 1, day)
      
      // 计算周数（从当月第一天开始）
      const firstDayOfMonth = new Date(year, month - 1, 1)
      const dayOfMonth = date.getDate()
      const weekNum = Math.ceil(dayOfMonth / 7)
      const weekKey = `${String(month).padStart(2, '0')}-W${weekNum}`
      
      if (!weekData[weekKey]) {
        weekData[weekKey] = { total: 0, userService: 0, scheduledTask: 0 }
      }
      
      weekData[weekKey].total += serviceTrend.trend[index]
      weekData[weekKey].userService += serviceTrend.detail[`用户服务_${label}`] || 0
      weekData[weekKey].scheduledTask += serviceTrend.detail[`定时任务_${label}`] || 0
    })
    
    labels = Object.keys(weekData).sort()
    trend = labels.map(key => weekData[key].total)
    labels.forEach(key => {
      detailData[`用户服务_${key}`] = weekData[key].userService
      detailData[`定时任务_${key}`] = weekData[key].scheduledTask
    })
  } else {
    detailData = serviceTrend.detail
  }

  serviceTrendChart?.setOption({
    title: { 
      text: '服务次数趋势', 
      textStyle: { 
        color: textColor.value,
        fontSize: 16,
        fontWeight: 600
      },
      left: 'center'
    },
    tooltip: {
      trigger: 'axis',
      formatter: (params: any) => {
        const data = params[0]
        const date = data.name
        const total = data.value
        const userService = detailData[`用户服务_${date}`] || 0
        const scheduledTask = detailData[`定时任务_${date}`] || 0
        return `${date}<br/>总服务次数: ${total}<br/>用户服务: ${userService}<br/>定时任务: ${scheduledTask}`
      }
    },
    xAxis: { 
      type: 'category', 
      data: labels, 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { show: false }
    },
    yAxis: { 
      type: 'value', 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { 
        lineStyle: { 
          color: 'rgba(148, 163, 184, 0.1)',
          type: 'dashed'
        }
      }
    },
    series: [
      {
        data: trend,
        type: 'line',
        smooth: true,
        symbol: 'circle',
        symbolSize: 8,
        lineStyle: {
          width: 3,
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: '#10b981' },
              { offset: 1, color: '#3b82f6' }
            ]
          }
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: 'rgba(16, 185, 129, 0.3)' },
              { offset: 1, color: 'rgba(59, 130, 246, 0.1)' }
            ]
          }
        },
        itemStyle: {
          color: '#10b981',
          borderColor: '#fff',
          borderWidth: 2
        },
        emphasis: {
          focus: 'series',
          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(16, 185, 129, 0.5)'
          }
        },
        animation: true,
        animationDuration: 1000,
        animationEasing: 'cubicOut'
      }
    ],
    grid: { left: 50, right: 30, top: 60, bottom: 40 },
    backgroundColor: 'transparent'
  })
}

const updateCharts = () => {
  if (!props.summary) return
  const { inspection, alerts } = props.summary

  updateServiceTrendChart()

  inspectionChart?.setOption({
    title: { 
      text: t('dashboard.metrics.inspectionTrend'), 
      textStyle: { 
        color: textColor.value,
        fontSize: 16,
        fontWeight: 600
      },
      left: 'center'
    },
    xAxis: { 
      type: 'category', 
      data: inspection.labels, 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { show: false }
    },
    yAxis: { 
      type: 'value', 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { 
        lineStyle: { 
          color: 'rgba(148, 163, 184, 0.1)',
          type: 'dashed'
        }
      }
    },
    series: [
      {
        data: inspection.trend,
        type: 'line',
        smooth: true,
        symbol: 'circle',
        symbolSize: 8,
        lineStyle: {
          width: 3,
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: '#7c3aed' },
              { offset: 1, color: '#3b82f6' }
            ]
          }
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: 'rgba(124, 58, 237, 0.3)' },
              { offset: 1, color: 'rgba(59, 130, 246, 0.1)' }
            ]
          }
        },
        itemStyle: {
          color: '#7c3aed',
          borderColor: '#fff',
          borderWidth: 2
        },
        emphasis: {
          focus: 'series',
          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(124, 58, 237, 0.5)'
          }
        },
        animation: true,
        animationDuration: 1000,
        animationEasing: 'cubicOut'
      }
    ],
    grid: { left: 50, right: 30, top: 60, bottom: 40 },
    backgroundColor: 'transparent'
  })

  alertsChart?.setOption({
    title: { 
      text: t('dashboard.metrics.alertLevel'), 
      textStyle: { 
        color: textColor.value,
        fontSize: 16,
        fontWeight: 600
      },
      left: 'center'
    },
    xAxis: { 
      type: 'category', 
      data: alerts.labels, 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { show: false }
    },
    yAxis: { 
      type: 'value', 
      axisLabel: { color: axisLabelColor.value },
      axisLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.2)' } },
      splitLine: { 
        lineStyle: { 
          color: 'rgba(148, 163, 184, 0.1)',
          type: 'dashed'
        }
      }
    },
    series: [
      {
        data: alerts.trend,
        type: 'bar',
        barWidth: '60%',
        itemStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: '#f59e0b' },
              { offset: 1, color: '#ef4444' }
            ]
          },
          borderRadius: [8, 8, 0, 0],
          shadowBlur: 10,
          shadowColor: 'rgba(245, 158, 11, 0.3)'
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 20,
            shadowColor: 'rgba(245, 158, 11, 0.5)'
          }
        },
        label: {
          show: true,
          position: 'top',
          color: textColor.value,
          fontWeight: 600
        },
        animation: true,
        animationDuration: 1000,
        animationEasing: 'bounceOut'
      }
    ],
    grid: { left: 50, right: 30, top: 60, bottom: 40 },
    backgroundColor: 'transparent'
  })
}

onMounted(() => {
  initCharts()
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  serviceTrendChart?.dispose()
  inspectionChart?.dispose()
  alertsChart?.dispose()
})

watch(
  () => props.summary,
  () => updateCharts(),
  { deep: true }
)

watch(
  () => themeStore.mode,
  () => updateCharts()
)
</script>

<style scoped>
.panel {
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-card);
  color: var(--text-primary);
  transition: all var(--transition-base);
  position: relative;
  overflow: hidden;
}

.panel::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: var(--gradient-accent);
  opacity: 0;
  transition: opacity var(--transition-base);
}

.panel:hover::before {
  opacity: 1;
}

.panel:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-card-hover);
  border-color: var(--border-color-active);
}

.panel__header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-xl);
  padding-bottom: var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
}

.header__content {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.header__icon {
  font-size: 2rem;
  filter: drop-shadow(0 0 8px rgba(124, 58, 237, 0.5));
}

.panel__title {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.panel__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  margin-top: var(--spacing-xs);
}

.realtime-tag {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-full);
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
}

.tag__dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--color-success);
  box-shadow: 0 0 8px var(--color-success);
}

.panel__metrics {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: var(--spacing-md);
  margin: var(--spacing-xl) 0;
}

.metric {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  border-radius: var(--radius-lg);
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  transition: all var(--transition-base);
  position: relative;
  overflow: hidden;
}

.metric::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 3px;
  height: 100%;
  background: var(--gradient-accent);
  transform: scaleY(0);
  transition: transform var(--transition-base);
}

.metric:hover::before {
  transform: scaleY(1);
}

.metric:hover {
  background: var(--bg-glass-hover);
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
  border-color: var(--border-color-hover);
}

.metric--highlight {
  background: var(--gradient-card);
  border: 2px solid var(--accent-color);
  box-shadow: var(--shadow-glow);
}

.metric--highlight::before {
  background: var(--gradient-accent);
  width: 4px;
}

.metric__icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-md);
  background: var(--accent-color-light);
  color: var(--accent-color);
  flex-shrink: 0;
}

.metric__icon svg {
  width: 24px;
  height: 24px;
}

.metric__icon--highlight {
  background: var(--gradient-accent);
  color: #fff;
  box-shadow: var(--shadow-glow);
}

.metric__content {
  flex: 1;
  min-width: 0;
}

.metric__label {
  font-size: 0.875rem;
  color: var(--text-muted);
  margin-bottom: var(--spacing-xs);
  font-weight: 500;
}

.metric__value {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.3;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  word-break: break-all;
}

.metric--highlight .metric__value {
  font-size: 2rem;
}

.panel__charts {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-lg);
  margin-top: var(--spacing-xl);
}

.chart-wrapper--wide {
  grid-column: span 2;
}

.chart {
  height: 280px;
  border-radius: var(--radius-lg);
  background: var(--bg-glass);
  padding: var(--spacing-md);
  transition: all var(--transition-base);
}

.chart:hover {
  background: var(--bg-glass-hover);
  box-shadow: var(--shadow-md);
}

.chart-wrapper {
  position: relative;
}

.chart-header {
  position: absolute;
  top: 15px;
  left: 20px;
  z-index: 10;
}

.chart--service {
  border-left: 3px solid var(--color-success);
}

.chart-summary {
  position: absolute;
  top: 50px;
  right: 20px;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  font-size: 0.875rem;
}

.summary-title {
  font-size: 0.75rem;
  color: var(--text-muted);
  font-weight: 600;
  text-align: center;
  padding-bottom: var(--spacing-xs);
  border-bottom: 1px solid var(--border-color);
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: var(--spacing-md);
}

.summary-label {
  color: var(--text-muted);
  font-size: 0.75rem;
}

.summary-value {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 1rem;
}

.chart--inspection {
  border-left: 3px solid var(--color-tech-purple-500);
}

.chart--alerts {
  border-left: 3px solid var(--color-warning);
}

/* 响应式 */
@media (max-width: 768px) {
  .panel__metrics {
    grid-template-columns: 1fr;
  }
  
  .panel__charts {
    grid-template-columns: 1fr;
  }
  
  .chart-wrapper--wide {
    grid-column: span 1;
  }
  
  .metric {
    flex-direction: column;
    text-align: center;
  }
}
</style>
