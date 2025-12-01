<template>
  <el-card shadow="never" class="panel card-hover glass-effect">
    <div class="panel__header">
      <div class="header__content">
        <div class="header__icon">📅</div>
        <div>
          <h2 class="panel__title">{{ t('dashboard.timeline.title') }}</h2>
          <p class="panel__subtitle">
            {{ t('dashboard.timeline.dutyOfficer') }}：<span class="highlight">{{ timeline.dutyOfficer }}</span> · {{ timeline.date }}
          </p>
        </div>
      </div>
      <el-button text type="primary" class="view-btn hover-lift" @click="goToDutySchedule">
        {{ t('dashboard.timeline.viewSchedule') }}
      </el-button>
    </div>

    <div class="panel__content">
      <section class="section-card">
        <div class="section-header">
          <h3 class="section-title">{{ t('dashboard.timeline.milestones') }}</h3>
          <el-button text type="primary" size="small" @click="goToMilestone">查看更多</el-button>
        </div>
        <el-steps direction="vertical" class="custom-steps">
          <el-step
            v-for="(item, index) in timeline.milestones"
            :key="item.time"
            :title="item.title"
            :description="item.time"
            class="step-item"
            :style="{ animationDelay: `${index * 0.1}s` }"
          />
        </el-steps>
      </section>
      <section class="section-card">
        <h3 class="section-title">{{ t('dashboard.timeline.deployPlans') }} <el-tag size="small" type="info">TODO</el-tag></h3>
        <el-table 
          :data="timeline.deployPlans" 
          size="small"
          class="custom-table"
          :row-class-name="tableRowClassName"
        >
          <el-table-column prop="system" :label="t('dashboard.timeline.system')" />
          <el-table-column prop="window" :label="t('dashboard.timeline.window')" />
          <el-table-column prop="owner" :label="t('dashboard.timeline.owner')" />
        </el-table>
      </section>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import type { TimelineInfo } from '../../types/dashboard'
import { useI18n } from '../../composables/useI18n'

const { t } = useI18n()
const router = useRouter()
defineProps<{ timeline: TimelineInfo }>()

const tableRowClassName = ({ rowIndex }: { rowIndex: number }) => {
  return rowIndex % 2 === 1 ? 'table-row-even' : ''
}

const goToDutySchedule = () => {
  router.push({ name: 'dutySchedule' })
}

const goToMilestone = () => {
  router.push({ name: 'milestone' })
}
</script>

<style scoped>
.panel {
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-card);
  transition: all var(--transition-base);
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
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.panel__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  margin-top: var(--spacing-xs);
}

.highlight {
  color: var(--accent-color);
  font-weight: 600;
}

.view-btn {
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  transition: all var(--transition-base);
}

.view-btn:hover {
  background: var(--bg-glass-hover);
  transform: translateY(-2px);
}

.panel__content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-xl);
}

.section-card {
  background: var(--bg-glass);
  padding: var(--spacing-lg);
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color);
  transition: all var(--transition-base);
}

.section-card:hover {
  background: var(--bg-glass-hover);
  border-color: var(--border-color-hover);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.section-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin: 0;
}

.section-title::before {
  content: '';
  width: 4px;
  height: 20px;
  background: var(--gradient-accent);
  border-radius: var(--radius-full);
}

.custom-steps {
  padding-left: var(--spacing-md);
  max-height: 500px;
  overflow-y: auto;
  overflow-x: hidden;
}

.custom-steps::-webkit-scrollbar {
  width: 6px;
}

.custom-steps::-webkit-scrollbar-track {
  background: var(--bg-glass);
  border-radius: var(--radius-full);
}

.custom-steps::-webkit-scrollbar-thumb {
  background: var(--border-color);
  border-radius: var(--radius-full);
}

.custom-steps::-webkit-scrollbar-thumb:hover {
  background: var(--accent-color);
}

.step-item {
  opacity: 0;
  animation: fadeInUp 0.5s ease-out forwards;
}

.custom-table {
  background: transparent;
}

.custom-table :deep(.el-table__row) {
  background: transparent;
  transition: all var(--transition-base);
}

.custom-table :deep(.el-table__row:hover) {
  background: var(--bg-glass-hover);
  transform: scale(1.01);
}

.custom-table :deep(.table-row-even) {
  background: var(--bg-glass);
}

.custom-table :deep(.el-table__header) {
  background: var(--bg-glass);
}

.custom-table :deep(.el-table th) {
  background: transparent;
  color: var(--text-secondary);
  font-weight: 600;
  border-bottom: 2px solid var(--border-color);
}

.custom-table :deep(.el-table td) {
  border-bottom: 1px solid var(--border-color);
  color: var(--text-primary);
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 768px) {
  .panel__content {
    grid-template-columns: 1fr;
  }
}

:deep(.el-tag) {
  background-color: #fef3c7 !important;
  color: #92400e !important;
  border-color: #fbbf24 !important;
}
</style>
