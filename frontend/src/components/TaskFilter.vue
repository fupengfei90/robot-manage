<template>
  <el-card shadow="never" class="task-filter glass-effect animate-fade-in">
    <el-form :model="filters" inline size="default" class="filter-form">
      <el-form-item label="任务名称">
        <el-input
          v-model="filters.name"
          placeholder="请输入任务名称"
          clearable
          style="width: 200px"
          @keyup.enter="handleSearch"
        />
      </el-form-item>
      
      <el-form-item label="任务状态">
        <el-select
          v-model="filters.active"
          placeholder="请选择状态"
          clearable
          style="width: 120px"
        >
          <el-option label="启用" :value="1" />
          <el-option label="停用" :value="0" />
        </el-select>
      </el-form-item>
      
      <el-form-item label="任务分类">
        <el-select
          v-model="filters.category"
          placeholder="请选择分类"
          clearable
          style="width: 150px"
        >
          <el-option
            v-for="category in categories"
            :key="category"
            :label="getCategoryLabel(category)"
            :value="category"
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="数据中心">
        <el-select
          v-model="filters.dcn"
          placeholder="请选择数据中心"
          clearable
          style="width: 120px"
        >
          <el-option
            v-for="dcn in dcns"
            :key="dcn"
            :label="dcn"
            :value="dcn"
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="创建时间">
        <el-date-picker
          v-model="dateRange"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          format="YYYY-MM-DD HH:mm:ss"
          value-format="YYYY-MM-DD HH:mm:ss"
          style="width: 350px"
          @change="handleDateChange"
        />
      </el-form-item>
      
      <el-form-item>
        <el-space>
          <el-button type="primary" @click="handleSearch" :icon="Search">
            搜索
          </el-button>
          <el-button @click="handleReset" :icon="Refresh">
            重置
          </el-button>
        </el-space>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Search, Refresh } from '@element-plus/icons-vue'
import { getScheduleTaskCategories, getScheduleTaskDCNs, type ScheduleTaskQuery } from '@/api/schedule-task'

interface Emits {
  (e: 'search', filters: ScheduleTaskQuery): void
}

const emit = defineEmits<Emits>()

const categories = ref<string[]>([])
const dcns = ref<string[]>([])
const dateRange = ref<[string, string] | null>(null)

const filters = reactive<ScheduleTaskQuery>({
  name: '',
  active: undefined,
  category: '',
  dcn: '',
  start_time: '',
  end_time: ''
})

// 分类标签映射
const categoryLabels: Record<string, string> = {
  'rota': '值班提醒',
  'weekly-report': '周报生成',
  'version-align': '版本通知'
}

const getCategoryLabel = (category: string) => {
  return categoryLabels[category] || category
}

// 处理日期范围变化
const handleDateChange = (dates: [string, string] | null) => {
  if (dates) {
    filters.start_time = dates[0]
    filters.end_time = dates[1]
  } else {
    filters.start_time = ''
    filters.end_time = ''
  }
}

// 搜索
const handleSearch = () => {
  const searchFilters: ScheduleTaskQuery = {
    name: filters.name || undefined,
    active: filters.active,
    category: filters.category || undefined,
    dcn: filters.dcn || undefined,
    start_time: filters.start_time || undefined,
    end_time: filters.end_time || undefined
  }
  emit('search', searchFilters)
}

// 重置
const handleReset = () => {
  Object.assign(filters, {
    name: '',
    active: undefined,
    category: '',
    dcn: '',
    start_time: '',
    end_time: ''
  })
  dateRange.value = null
  handleSearch()
}

// 加载选项
const loadOptions = async () => {
  try {
    const [categoriesRes, dcnsRes] = await Promise.all([
      getScheduleTaskCategories(),
      getScheduleTaskDCNs()
    ])
    categories.value = categoriesRes.data
    dcns.value = dcnsRes.data
  } catch (error) {
    console.error('加载选项失败:', error)
  }
}

onMounted(() => {
  loadOptions()
})
</script>

<style scoped>
.task-filter {
  margin-bottom: var(--spacing-sm);
  border-radius: var(--radius-xl);
}

.task-filter :deep(.el-card__body) {
  padding: var(--spacing-md);
}

.filter-form {
  margin: 0;
}
</style>