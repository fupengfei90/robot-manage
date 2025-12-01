<template>
  <div class="digital-human-view">
    <el-card shadow="never" class="feature-card">
      <div class="card-header">
        <h3 class="card-title">数字人形象动态展示 <el-tag size="small" type="warning">TODO</el-tag></h3>
      </div>
      <div class="card-content">
        <p class="placeholder-text">数字人形象动态展示功能开发中...</p>
      </div>
    </el-card>

    <el-card shadow="never" class="feature-card">
      <div class="card-header">
        <h3 class="card-title">数字人互动功能 <el-tag size="small" type="warning">TODO</el-tag></h3>
      </div>
      <div class="card-content">
        <p class="placeholder-text">数字人互动功能开发中...</p>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.digital-human-view {
  display: grid;
  gap: var(--spacing-xl);
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
}

.feature-card {
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
  transition: all var(--transition-base);
}

.feature-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-card-hover);
  border-color: var(--border-color-active);
}

.card-header {
  padding-bottom: var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
  margin-bottom: var(--spacing-lg);
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--text-primary);
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin: 0;
}

.card-content {
  min-height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder-text {
  color: var(--text-muted);
  font-size: 1rem;
}

:deep(.el-tag) {
  background-color: #fef3c7 !important;
  color: #92400e !important;
  border-color: #fbbf24 !important;
}

@media (max-width: 768px) {
  .digital-human-view {
    grid-template-columns: 1fr;
  }
}
</style>
