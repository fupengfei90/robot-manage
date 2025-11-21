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
      <div 
        v-for="(value, key) in summary.serviceCounts" 
        :key="key" 
        class="metric card-hover"
        :style="{ animationDelay: `${parseInt(key) * 0.1}s` }"
      >
        <div class="metric__icon">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13 2L3 14H12L11 22L21 10H12L13 2Z" stroke="currentColor" stroke-width="2" fill="none"/>
          </svg>
        </div>
        <div class="metric__content">
          <p class="metric__label">{{ serviceDict[key as keyof typeof serviceDict] }}</p>
          <p class="metric__value">{{ value.toLocaleString() }}</p>
        </div>
      </div>
      <div class="metric metric--highlight card-hover">
        <div class="metric__icon metric__icon--highlight">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none"/>
            <path d="M12 6V12L16 14" stroke="currentColor" stroke-width="2" fill="none"/>
          </svg>
        </div>
        <div class="metric__content">
          <p class="metric__label">{{ t('dashboard.metrics.users') }}</p>
          <p class="metric__value">{{ summary.userCount.toLocaleString() }}</p>
        </div>
      </div>
    </div>

    <div class="panel__charts">
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

const inspectionRef = ref<HTMLDivElement | null>(null)
const alertsRef = ref<HTMLDivElement | null>(null)
let inspectionChart: echarts.ECharts | null = null
let alertsChart: echarts.ECharts | null = null

const textColor = computed(() => {
  return themeStore.mode === 'dark' ? '#e2e8f0' : '#1e293b'
})

const axisLabelColor = computed(() => {
  return themeStore.mode === 'dark' ? '#94a3b8' : '#64748b'
})

const serviceDict = computed(() => ({
  daily: t('dashboard.metrics.daily'),
  weekly: t('dashboard.metrics.weekly'),
  monthly: t('dashboard.metrics.monthly'),
  yearly: t('dashboard.metrics.yearly')
}))

const initCharts = () => {
  if (inspectionRef.value) {
    inspectionChart = echarts.init(inspectionRef.value)
  }
  if (alertsRef.value) {
    alertsChart = echarts.init(alertsRef.value)
  }
  updateCharts()
}

const handleResize = () => {
  inspectionChart?.resize()
  alertsChart?.resize()
}

const updateCharts = () => {
  if (!props.summary) return
  const { inspection, alerts } = props.summary

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
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
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
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.2;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.metric--highlight .metric__value {
  font-size: 2rem;
}

.panel__charts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-lg);
  margin-top: var(--spacing-xl);
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
  
  .metric {
    flex-direction: column;
    text-align: center;
  }
}
</style>
