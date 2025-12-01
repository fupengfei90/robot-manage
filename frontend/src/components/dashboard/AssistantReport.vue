<template>
  <el-card shadow="never" class="panel card-hover glass-effect">
    <div class="panel__header">
      <div class="header__content">
        <div class="header__icon">📝</div>
        <div>
          <h2 class="panel__title">{{ t('dashboard.assistant.title') }} <el-tag size="small" type="warning">TODO</el-tag></h2>
          <p class="panel__subtitle">{{ t('dashboard.assistant.range') }}：{{ report.range }}</p>
        </div>
      </div>
      <el-button-group class="range-buttons">
        <el-button 
          v-for="range in ranges" 
          :key="range" 
          size="small" 
          :type="report.range === range ? 'primary' : 'default'"
          class="range-btn hover-lift"
          @click="$emit('change-range', range)"
        >
          {{ rangeLabels[range as keyof typeof rangeLabels] }}
        </el-button>
      </el-button-group>
    </div>

    <div class="report">
      <section class="report-section report-section--highlights">
        <div class="section-header">
          <h3 class="section-title">{{ t('dashboard.assistant.highlights') }}</h3>
          <span class="section-icon">✨</span>
        </div>
        <ul class="report-list">
          <li v-for="(item, index) in report.highlights" :key="item" class="list-item" :style="{ animationDelay: `${index * 0.1}s` }">
            <span class="list-bullet"></span>
            {{ item }}
          </li>
        </ul>
      </section>
      <section class="report-section report-section--risks">
        <div class="section-header">
          <h3 class="section-title">{{ t('dashboard.assistant.risks') }}</h3>
          <span class="section-icon">⚠️</span>
        </div>
        <ul class="report-list">
          <li v-for="(item, index) in report.risks" :key="item" class="list-item" :style="{ animationDelay: `${index * 0.1}s` }">
            <span class="list-bullet list-bullet--warning"></span>
            {{ item }}
          </li>
        </ul>
      </section>
      <section class="report-section report-section--next">
        <div class="section-header">
          <h3 class="section-title">{{ t('dashboard.assistant.nextSteps') }}</h3>
          <span class="section-icon">🎯</span>
        </div>
        <ul class="report-list">
          <li v-for="(item, index) in report.nextSteps" :key="item" class="list-item" :style="{ animationDelay: `${index * 0.1}s` }">
            <span class="list-bullet list-bullet--success"></span>
            {{ item }}
          </li>
        </ul>
      </section>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { AssistantReport } from '../../types/dashboard'
import { useI18n } from '../../composables/useI18n'

const { t } = useI18n()
defineProps<{ report: AssistantReport }>()

defineEmits<{ (e: 'change-range', range: string): void }>()

const ranges = ['daily', 'weekly', 'monthly', 'yearly']
const rangeLabels = computed(() => ({
  daily: t('dashboard.assistant.daily'),
  weekly: t('dashboard.assistant.weekly'),
  monthly: t('dashboard.assistant.monthly'),
  yearly: t('dashboard.assistant.yearly')
}))
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

.panel__title :deep(.el-tag) {
  -webkit-text-fill-color: #92400e !important;
}

.panel__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  margin-top: var(--spacing-xs);
}

.range-buttons {
  display: flex;
  gap: var(--spacing-xs);
  background: var(--bg-glass);
  padding: var(--spacing-xs);
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
}

.range-btn {
  border-radius: var(--radius-sm);
  transition: all var(--transition-base);
}

.report {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: var(--spacing-lg);
  margin-top: var(--spacing-lg);
}

.report-section {
  background: var(--bg-glass);
  padding: var(--spacing-lg);
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color);
  transition: all var(--transition-base);
}

.report-section:hover {
  background: var(--bg-glass-hover);
  border-color: var(--border-color-hover);
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}

.report-section--highlights {
  border-left: 4px solid var(--color-success);
}

.report-section--risks {
  border-left: 4px solid var(--color-warning);
}

.report-section--next {
  border-left: 4px solid var(--accent-color);
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: var(--spacing-md);
}

.section-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--text-secondary);
  margin: 0;
}

.section-icon {
  font-size: 1.5rem;
}

.report-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.list-item {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  margin-bottom: var(--spacing-sm);
  background: var(--bg-glass);
  border-radius: var(--radius-md);
  color: var(--text-secondary);
  transition: all var(--transition-base);
  opacity: 0;
  animation: fadeInUp 0.5s ease-out forwards;
}

.list-item:hover {
  background: var(--bg-glass-hover);
  color: var(--text-primary);
  transform: translateX(4px);
}

.list-bullet {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--color-success);
  margin-top: 6px;
  flex-shrink: 0;
  box-shadow: 0 0 8px var(--color-success);
}

.list-bullet--warning {
  background: var(--color-warning);
  box-shadow: 0 0 8px var(--color-warning);
}

.list-bullet--success {
  background: var(--accent-color);
  box-shadow: 0 0 8px var(--accent-color);
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
  .report {
    grid-template-columns: 1fr;
  }
}

.panel__title :deep(.el-tag) {
  background-color: #fef3c7 !important;
  color: #92400e !important;
  border-color: #fbbf24 !important;
}
</style>
