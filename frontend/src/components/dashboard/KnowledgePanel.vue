<template>
  <el-card shadow="never" class="panel card-hover glass-effect">
    <div class="panel__header">
      <div class="header__content">
        <div class="header__icon">📚</div>
        <div>
          <h2 class="panel__title">{{ t('dashboard.knowledge.title') }}</h2>
        </div>
      </div>
      <el-tag type="warning" class="recall-tag">
        <span class="tag__icon">⭐</span>
        {{ t('dashboard.knowledge.recallRate') }} {{ (knowledge.recallRate * 100).toFixed(0) }}%
      </el-tag>
    </div>
    <div class="panel__columns">
      <section class="section-card">
        <h3 class="section-title">{{ t('dashboard.knowledge.recentUpdates') }}</h3>
        <el-timeline class="custom-timeline">
          <el-timeline-item
            v-for="(item, index) in knowledge.recentUpdates"
            :key="item.title"
            :timestamp="item.updatedAt"
            class="timeline-item"
            :style="{ animationDelay: `${index * 0.1}s` }"
          >
            <div class="timeline-content">
              <p class="timeline-title">{{ item.title }}</p>
              <small class="timeline-meta">
                {{ t('dashboard.knowledge.author') }} {{ item.author }} · {{ item.views }} {{ t('dashboard.knowledge.views') }}
              </small>
            </div>
          </el-timeline-item>
        </el-timeline>
      </section>
      <section class="section-card">
        <h3 class="section-title">{{ t('dashboard.knowledge.hotTopics') }}</h3>
        <div 
          class="hot-item card-hover" 
          v-for="(item, index) in knowledge.hotTopics" 
          :key="item.title"
          :style="{ animationDelay: `${index * 0.1}s` }"
        >
          <div class="hot-item__content">
            <p class="hot-item__title">{{ item.title }}</p>
            <small class="hot-item__meta">{{ t('dashboard.knowledge.author') }} {{ item.author }}</small>
          </div>
          <div class="hot-item__views">
            <span class="views-icon">🔥</span>
            <strong>{{ item.views }}</strong>
          </div>
        </div>
      </section>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import type { KnowledgeSnapshot } from '../../types/dashboard'
import { useI18n } from '../../composables/useI18n'

const { t } = useI18n()
defineProps<{ knowledge: KnowledgeSnapshot }>()
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
}

.recall-tag {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-full);
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
}

.tag__icon {
  font-size: 1rem;
}

.panel__columns {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-xl);
  margin-top: var(--spacing-lg);
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

.section-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.section-title::before {
  content: '';
  width: 4px;
  height: 20px;
  background: var(--gradient-accent);
  border-radius: var(--radius-full);
}

.custom-timeline {
  padding-left: var(--spacing-md);
}

.timeline-item {
  opacity: 0;
  animation: fadeInUp 0.5s ease-out forwards;
}

.timeline-content {
  padding-left: var(--spacing-md);
}

.timeline-title {
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
  transition: color var(--transition-base);
}

.timeline-item:hover .timeline-title {
  color: var(--accent-color);
}

.timeline-meta {
  color: var(--text-muted);
  font-size: 0.875rem;
}

.hot-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--bg-glass);
  padding: var(--spacing-md) var(--spacing-lg);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-md);
  border: 1px solid var(--border-color);
  transition: all var(--transition-base);
  opacity: 0;
  animation: fadeInUp 0.5s ease-out forwards;
}

.hot-item:hover {
  background: var(--bg-glass-hover);
  transform: translateX(4px);
  border-color: var(--border-color-hover);
  box-shadow: var(--shadow-sm);
}

.hot-item__content {
  flex: 1;
  min-width: 0;
}

.hot-item__title {
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.hot-item__meta {
  color: var(--text-muted);
  font-size: 0.875rem;
}

.hot-item__views {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  padding: var(--spacing-sm) var(--spacing-md);
  background: var(--accent-color-light);
  border-radius: var(--radius-full);
  color: var(--accent-color);
  font-weight: 700;
}

.views-icon {
  font-size: 1.125rem;
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
  .panel__columns {
    grid-template-columns: 1fr;
  }
}
</style>
