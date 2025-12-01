<template>
  <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
    <el-form :inline="true" :model="filters" class="filter-form">
      <el-form-item :label="t('system.task.taskName')">
        <el-input
          v-model="filters.name"
          :placeholder="t('system.task.searchPlaceholder')"
          clearable
          style="width: 200px"
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </el-form-item>
      
      <el-form-item :label="t('system.task.status')">
        <el-select
          v-model="filters.active"
          :placeholder="t('system.task.selectStatus')"
          clearable
          style="width: 120px"
        >
          <el-option :label="t('common.enabled')" :value="1" />
          <el-option :label="t('common.disabled')" :value="0" />
        </el-select>
      </el-form-item>
      
      <el-form-item :label="t('system.task.category')">
        <el-input
          v-model="filters.category"
          :placeholder="t('system.task.selectCategory')"
          clearable
          style="width: 150px"
          @keyup.enter="handleSearch"
        />
      </el-form-item>
      
      <el-form-item>
        <el-space>
          <el-button type="primary" @click="handleSearch" :icon="Search">
            {{ t('common.search') }}
          </el-button>
          <el-button @click="handleReset" :icon="Refresh">
            {{ t('common.reset') }}
          </el-button>
        </el-space>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { Search, Refresh } from '@element-plus/icons-vue'
import { type ScheduleTaskQuery } from '@/api/schedule-task'
import { useCommon } from '@/composables/useCommon'

const { t } = useCommon()

interface Emits {
  (e: 'search', filters: ScheduleTaskQuery): void
}

const emit = defineEmits<Emits>()

const filters = reactive<ScheduleTaskQuery>({
  name: '',
  active: undefined,
  category: ''
})

const handleSearch = () => {
  const searchFilters: ScheduleTaskQuery = {
    name: filters.name || undefined,
    active: filters.active,
    category: filters.category || undefined
  }
  emit('search', searchFilters)
}

const handleReset = () => {
  Object.assign(filters, {
    name: '',
    active: undefined,
    category: ''
  })
  handleSearch()
}
</script>

<style scoped>
.filter-card {
  margin-bottom: var(--spacing-sm);
  border-radius: var(--radius-xl);
}

.filter-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}

.filter-form {
  margin: 0;
}
</style>