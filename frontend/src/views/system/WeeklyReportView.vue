<template>
  <div class="weekly-report-view">
    <div class="report-header">
      <div class="report-title">
        <span class="title-icon">📊</span>
        <h1>【富小娇】周报</h1>
      </div>
      <el-button type="primary" @click="exportReport">导出报告</el-button>
    </div>

    <div class="report-content glass-effect">
      <h2 class="report-subtitle">富小娇系统周度工作报告【截至{{ currentDate }}】</h2>

      <div class="report-section">
        <h3>• 服务时效</h3>
        <ul>
          <li>7x24 小时在线</li>
        </ul>
      </div>

      <div class="report-section">
        <h3>• 服务规模</h3>
        <ul>
          <li>
            本周累计完成服务<span class="highlight">{{ weeklyStats.serviceCount }}</span>次，
            覆盖用户<span class="highlight">{{ weeklyStats.userCount }}</span>人，
            系统活跃度保持稳定。
          </li>
        </ul>
      </div>

      <div class="report-section">
        <h3>• 交付前支持</h3>
        <ul>
          <li>
            提供CMDB及知识库查询服务<span class="highlight">{{ weeklyStats.cmdbQueryCount }}</span>次，
            为项目交付提供准确的数据与知识基础。
          </li>
        </ul>
      </div>

      <div class="report-section">
        <h3>• 交付过程管控</h3>
        <ul>
          <li>
            完成<span class="highlight">{{ weeklyStats.ticketCount }}</span>次提单操作，
            提升交付准确率和质量，减少流程对接的人力成本。
          </li>
        </ul>
      </div>

      <div class="report-section">
        <h3>• 线上运维保障：</h3>
        <ul>
          <li>
            通过主动定时巡检与ECC事件单处理<span class="highlight">{{ weeklyStats.inspectionCount }}</span>，
            实现对线上环境的持续监控与主动运维，保障系统稳定性。
          </li>
        </ul>
      </div>

      <div class="report-section">
        <h3>• 核心价值总结：</h3>
        <ul>
          <li>
            系统本周在支撑交付效率【快速查询】、保障交付质量【异常发现】与维护系统稳定【主动运维】三大维度均提供了关键价值。
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { fetchDashboardSummary } from '../../api/dashboard'

interface WeeklyStats {
  serviceCount: number
  userCount: number
  cmdbQueryCount: number
  ticketCount: number
  inspectionCount: number
}

const weeklyStats = ref<WeeklyStats>({
  serviceCount: 0,
  userCount: 0,
  cmdbQueryCount: 0,
  ticketCount: 0,
  inspectionCount: 0
})

const currentDate = computed(() => {
  const now = new Date()
  const month = now.getMonth() + 1
  const day = now.getDate()
  return `${month}月${day}日`
})

const fetchWeeklyData = async () => {
  try {
    const data = await fetchDashboardSummary()
    const weeklyTotal = data.serviceCounts?.weekly || 0
    const weeklyUserService = data.serviceCounts?.weekly_user_service || 0
    const weeklyScheduledTask = data.serviceCounts?.weekly_scheduled_task || 0
    const weeklyTicketCount = data.serviceCounts?.weekly_ticket_count || 0
    weeklyStats.value = {
      serviceCount: weeklyTotal,
      userCount: data.serviceUsers || 0,
      cmdbQueryCount: weeklyUserService,
      ticketCount: weeklyTicketCount,
      inspectionCount: weeklyScheduledTask
    }
  } catch (error) {
    console.error('Failed to fetch weekly data:', error)
  }
}

const exportReport = () => {
  const content = document.querySelector('.report-content')
  if (content) {
    const text = content.textContent || ''
    const blob = new Blob([text], { type: 'text/plain;charset=utf-8' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `富小娇周报_${currentDate.value}.txt`
    a.click()
    URL.revokeObjectURL(url)
    ElMessage.success('报告已导出')
  }
}

onMounted(() => {
  fetchWeeklyData()
})
</script>

<style scoped>
.weekly-report-view {
  padding: var(--spacing-md);
  max-width: 1200px;
  margin: 0 auto;
}

.report-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.report-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.title-icon {
  font-size: 2rem;
}

.report-title h1 {
  font-size: 2rem;
  font-weight: 700;
  margin: 0;
  color: var(--text-primary);
}

.report-content {
  padding: var(--spacing-xl);
  border-radius: var(--radius-lg);
  line-height: 1.8;
}

.report-subtitle {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: var(--spacing-xl);
  color: var(--text-primary);
}

.report-section {
  margin-bottom: var(--spacing-lg);
}

.report-section h3 {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: var(--spacing-sm);
  color: var(--text-primary);
}

.report-section ul {
  list-style: none;
  padding-left: var(--spacing-xl);
  margin: 0;
}

.report-section li {
  font-size: 1rem;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xs);
  position: relative;
}

.report-section li::before {
  content: '•';
  position: absolute;
  left: -20px;
  color: var(--accent-color);
}

.highlight {
  color: #ef4444;
  font-weight: 700;
  font-size: 1.125rem;
  margin: 0 4px;
}
</style>
