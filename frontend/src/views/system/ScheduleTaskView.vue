<template>
  <div class="schedule-task-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">调度任务管理</h1>
        <p class="view-subtitle">管理系统中的定时任务，支持创建、编辑、启用/停用和立即执行等操作</p>
      </div>
      <div class="header-actions">
        <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
          <span>➕</span> 新增任务
        </el-button>
        <el-dropdown @command="handleBatchOperation" :disabled="!selectedTasks.length">
          <el-button :disabled="!selectedTasks.length" class="batch-btn">
            批量操作
            <el-icon class="el-icon--right"><ArrowDown /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="enable">批量启用</el-dropdown-item>
              <el-dropdown-item command="disable">批量停用</el-dropdown-item>
              <el-dropdown-item command="delete" divided>批量删除</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>

    <!-- 搜索筛选 -->
    <TaskFilter @search="handleSearch" />

    <!-- 数据表格 -->
    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        v-loading="loading"
        :data="tableData"
        @selection-change="handleSelectionChange"
        @sort-change="handleSortChange"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        
        <el-table-column prop="id" label="ID" width="80" sortable="custom" />
        
        <el-table-column prop="name" label="任务名称" min-width="150" show-overflow-tooltip>
          <template #default="{ row }">
            <div class="task-name">
              <span>{{ row.name }}</span>
              <el-tag v-if="row.category" :type="getCategoryType(row.category)" size="small" class="category-tag">
                {{ getCategoryLabel(row.category) }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="status_text" label="状态" width="100">
          <template #default="{ row }">
            <el-switch
              :model-value="row.active"
              :active-value="1"
              :inactive-value="0"
              @change="(val) => handleStatusChange(row, val)"
              :loading="statusLoading[row.id]"
            />
          </template>
        </el-table-column>
        
        <el-table-column prop="cron_desc" label="执行时间" min-width="120" show-overflow-tooltip :show-overflow-tooltip="{ popperClass: 'table-tooltip' }" />
        
        <el-table-column prop="next_run_time" label="下次执行" width="160">
          <template #default="{ row }">
            <span v-if="row.next_run_time" class="next-run-time">
              {{ formatDateTime(row.next_run_time) }}
            </span>
            <el-text v-else type="info" size="small">已停用</el-text>
          </template>
        </el-table-column>
        
        <el-table-column prop="dcn" label="数据中心" width="100">
          <template #default="{ row }">
            <el-tag type="info" size="small">{{ row.dcn }}</el-tag>
          </template>
        </el-table-column>
        
        <el-table-column prop="workflow" label="工作流ID" width="120" show-overflow-tooltip />
        
        <el-table-column prop="updated_at" label="更新时间" width="160" sortable="custom">
          <template #default="{ row }">
            {{ formatDateTime(row.updated_at) }}
          </template>
        </el-table-column>
        
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              @click="handleExecute(row)"
              :disabled="row.active === 0"
              :loading="executeLoading[row.id]"
            >
              执行
            </el-button>
            <el-button link type="primary" @click="handleEdit(row)">
              编辑
            </el-button>
            <el-button
              link
              type="danger"
              @click="handleDelete(row)"
              :loading="deleteLoading[row.id]"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 任务表单对话框 -->
    <TaskForm
      v-model="formVisible"
      :task="currentTask"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, ArrowDown } from '@element-plus/icons-vue'
import TaskFilter from '@/components/TaskFilter.vue'
import TaskForm from '@/components/TaskForm.vue'
import {
  getScheduleTaskList,
  deleteScheduleTask,
  updateScheduleTaskStatus,
  executeScheduleTask,
  batchOperateScheduleTasks,
  type ScheduleTask,
  type ScheduleTaskQuery
} from '@/api/schedule-task'

// 响应式数据
const loading = ref(false)
const tableData = ref<ScheduleTask[]>([])
const selectedTasks = ref<ScheduleTask[]>([])
const formVisible = ref(false)
const currentTask = ref<ScheduleTask | null>(null)

// 加载状态
const statusLoading = ref<Record<number, boolean>>({})
const executeLoading = ref<Record<number, boolean>>({})
const deleteLoading = ref<Record<number, boolean>>({})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 查询条件
const queryParams = ref<ScheduleTaskQuery>({})

// 分类标签映射
const categoryLabels: Record<string, string> = {
  'rota': '值班提醒',
  'weekly-report': '周报生成',
  'version-align': '版本通知'
}

const getCategoryLabel = (category: string) => {
  return categoryLabels[category] || category
}

const getCategoryType = (category: string) => {
  const typeMap: Record<string, string> = {
    'rota': 'warning',
    'weekly-report': 'success',
    'version-align': 'info'
  }
  return typeMap[category] || 'info'
}

// 格式化日期时间
const formatDateTime = (dateTime: string) => {
  if (!dateTime) return ''
  return new Date(dateTime).toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

// 加载数据
const loadData = async () => {
  try {
    loading.value = true
    const params = {
      ...queryParams.value,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    
    const response = await getScheduleTaskList(params)
    console.log('API响应:', response)
    // 后端返回的数据结构是直接返回对象
    tableData.value = response.list || []
    pagination.total = response.total || 0
    console.log('表格数据:', tableData.value)
    console.log('总数:', pagination.total)
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败')
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = (filters: ScheduleTaskQuery) => {
  queryParams.value = filters
  pagination.page = 1
  loadData()
}

// 排序处理
const handleSortChange = ({ prop, order }: { prop: string; order: string | null }) => {
  if (order) {
    queryParams.value.order_by = prop
    queryParams.value.order_dir = order === 'ascending' ? 'asc' : 'desc'
  } else {
    queryParams.value.order_by = undefined
    queryParams.value.order_dir = undefined
  }
  loadData()
}

// 分页处理
const handlePageChange = (page: number) => {
  pagination.page = page
  loadData()
}

const handlePageSizeChange = (pageSize: number) => {
  pagination.pageSize = pageSize
  pagination.page = 1
  loadData()
}

// 选择处理
const handleSelectionChange = (selection: ScheduleTask[]) => {
  selectedTasks.value = selection
}

// 新增任务
const handleCreate = () => {
  currentTask.value = null
  formVisible.value = true
}

// 编辑任务
const handleEdit = (task: ScheduleTask) => {
  currentTask.value = task
  formVisible.value = true
}

// 删除任务
const handleDelete = async (task: ScheduleTask) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除任务"${task.name}"吗？删除后无法恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    deleteLoading.value[task.id] = true
    await deleteScheduleTask(task.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  } finally {
    deleteLoading.value[task.id] = false
  }
}

// 状态切换
const handleStatusChange = async (task: ScheduleTask, active: number) => {
  try {
    statusLoading.value[task.id] = true
    await updateScheduleTaskStatus(task.id, active)
    task.active = active
    task.status_text = active === 1 ? '启用' : '停用'
    ElMessage.success(`任务已${active === 1 ? '启用' : '停用'}`)
    loadData() // 重新加载以更新下次执行时间
  } catch (error) {
    console.error('状态更新失败:', error)
    ElMessage.error('状态更新失败')
  } finally {
    statusLoading.value[task.id] = false
  }
}

// 立即执行
const handleExecute = async (task: ScheduleTask) => {
  try {
    executeLoading.value[task.id] = true
    await executeScheduleTask(task.id)
    ElMessage.success('任务执行成功')
  } catch (error) {
    console.error('任务执行失败:', error)
    ElMessage.error('任务执行失败')
  } finally {
    executeLoading.value[task.id] = false
  }
}

// 批量操作
const handleBatchOperation = async (command: string) => {
  if (!selectedTasks.value.length) return
  
  const operationMap: Record<string, string> = {
    enable: '启用',
    disable: '停用',
    delete: '删除'
  }
  
  const operation = operationMap[command]
  if (!operation) return
  
  try {
    await ElMessageBox.confirm(
      `确定要${operation}选中的 ${selectedTasks.value.length} 个任务吗？`,
      `确认${operation}`,
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: command === 'delete' ? 'warning' : 'info'
      }
    )
    
    const ids = selectedTasks.value.map(task => task.id)
    await batchOperateScheduleTasks({
      ids,
      operation: command as 'enable' | 'disable' | 'delete'
    })
    
    ElMessage.success(`批量${operation}成功`)
    selectedTasks.value = []
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(`批量${operation}失败:`, error)
      ElMessage.error(`批量${operation}失败`)
    }
  }
}

// 表单成功处理
const handleFormSuccess = () => {
  loadData()
}

// 初始化
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.schedule-task-view {
  padding: var(--spacing-md);
}

.view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: var(--radius-xl);
}

.create-btn {
  padding: var(--spacing-md) var(--spacing-xl);
  border-radius: var(--radius-md);
}

.batch-btn {
  padding: var(--spacing-md) var(--spacing-lg);
  border-radius: var(--radius-md);
}

.table-card {
  border-radius: var(--radius-xl);
}

.custom-table {
  background: transparent;
}

.custom-table :deep(.el-table__row) {
  transition: all var(--transition-base);
}

.custom-table :deep(.el-table__row:hover) {
  background: var(--bg-glass-hover);
}

.task-name {
  display: flex;
  align-items: center;
  gap: 8px;
}

.category-tag {
  font-size: 11px;
}

.next-run-time {
  color: var(--accent-color);
  font-weight: 500;
}

.pagination-container {
  display: flex;
  justify-content: flex-end;
  margin-top: var(--spacing-md);
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}

.table-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}

/* 修复表格tooltip定位问题 */
.custom-table :deep(.el-tooltip__trigger) {
  display: inline-block;
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.custom-table :deep(.el-table__cell) {
  overflow: visible !important;
}
</style>