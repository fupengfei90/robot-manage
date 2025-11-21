<template>
  <div class="dashboard">
    <el-skeleton v-if="store.loading" :rows="6" animated class="skeleton-loader" />
    <template v-else>
      <div class="dashboard__grid">
        <div class="dashboard__item animate-fade-in-up" style="animation-delay: 0.1s">
          <MetricsPanel v-if="store.summary" :summary="store.summary" />
        </div>
        <div class="dashboard__item animate-fade-in-up" style="animation-delay: 0.2s">
          <KnowledgePanel v-if="store.knowledge" :knowledge="store.knowledge" />
        </div>
        <div class="dashboard__item animate-fade-in-up" style="animation-delay: 0.3s">
          <TimelinePanel v-if="store.timeline" :timeline="store.timeline" />
        </div>
        <div class="dashboard__item animate-fade-in-up" style="animation-delay: 0.4s">
          <AssistantReport
            v-if="store.assistant"
            :report="store.assistant"
            @change-range="changeRange"
          />
        </div>
        <div class="dashboard__item animate-fade-in-up" style="animation-delay: 0.5s">
          <DigitalHuman />
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import MetricsPanel from '../../components/dashboard/MetricsPanel.vue'
import KnowledgePanel from '../../components/dashboard/KnowledgePanel.vue'
import TimelinePanel from '../../components/dashboard/TimelinePanel.vue'
import AssistantReport from '../../components/dashboard/AssistantReport.vue'
import DigitalHuman from '../../components/dashboard/DigitalHuman.vue'
import { useDashboardStore } from '../../stores/dashboard'

const store = useDashboardStore()

const changeRange = (range: string) => {
  store.initialize(range)
}

onMounted(() => {
  store.initialize('daily')
})
</script>

<style scoped>
.dashboard {
  position: relative;
  width: 100%;
}

.dashboard__grid {
  display: grid;
  gap: var(--spacing-xl);
  grid-template-columns: repeat(auto-fit, minmax(380px, 1fr));
  grid-auto-rows: minmax(300px, auto);
}

.dashboard__item {
  position: relative;
  height: 100%;
}

/* 不对称布局 - 突出重要区域 */
.dashboard__item:first-child {
  grid-column: span 2;
}

.dashboard__item:nth-child(2) {
  grid-row: span 2;
}

@media (max-width: 1200px) {
  .dashboard__grid {
    grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  }
  
  .dashboard__item:first-child {
    grid-column: span 1;
  }
}

@media (max-width: 768px) {
  .dashboard__grid {
    grid-template-columns: 1fr;
    gap: var(--spacing-lg);
  }
}

.skeleton-loader {
  background: var(--bg-card);
  border-radius: var(--radius-xl);
  padding: var(--spacing-xl);
  backdrop-filter: blur(20px);
}
</style>
