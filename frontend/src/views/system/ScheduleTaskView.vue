<template>
  <div class="schedule-task-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">{{ t('system.task.title') }}</h1>
        <p class="view-subtitle">{{ t('system.task.subtitle') }}</p>
      </div>
      <div class="header-actions">
        <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
          <span>➕</span> {{ t('system.task.addTask') }}
        </el-button>
        <el-dropdown @command="handleBatchOperation" :disabled="!selectedTasks.length">
          <el-button :disabled="!selectedTasks.length" class="batch-btn">
            {{ t('system.task.batchOperations') }}
            <el-icon class="el-icon--right"><ArrowDown /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="enable">{{ t('system.task.batchEnable') }}</el-dropdown-item>
              <el-dropdown-item command="disable">{{ t('system.task.batchDisable') }}</el-dropdown-item>
              <el-dropdown-item command="delete" divided>{{ t('system.task.batchDelete') }}</el-dropdown-item>
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
        
        <el-table-column prop="name" :label="t('system.task.taskName')" min-width="200" show-overflow-tooltip>
          <template #default="{ row }">
            <div class="task-name">
              <span>{{ row.name }}</span>
              <el-tag v-if="row.category" :type="getCategoryType(row.category)" size="small" class="category-tag">
                {{ getCategoryLabel(row.category) }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="status_text" :label="t('system.task.status')" width="100">
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
        
        <el-table-column prop="cron_desc" :label="t('system.task.executionTime')" min-width="140" show-overflow-tooltip :show-overflow-tooltip="{ popperClass: 'table-tooltip' }" />
        
        <el-table-column prop="next_run_time" :label="t('system.task.nextExecution')" width="140">
          <template #default="{ row }">
            <span v-if="row.next_run_time" class="next-run-time">
              {{ formatDateTime(row.next_run_time) }}
            </span>
            <el-text v-else type="info" size="small">{{ t('system.task.disabled') }}</el-text>
          </template>
        </el-table-column>
        

        
        <el-table-column prop="updated_at" :label="t('system.task.updateTime')" width="140" sortable="custom">
          <template #default="{ row }">
            {{ formatDateTime(row.updated_at) }}
          </template>
        </el-table-column>
        
        <el-table-column :label="t('system.task.actions')" width="120" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              @click="handleExecute(row)"
              :disabled="row.active === 0"
              :loading="executeLoading[row.id]"
            >
              {{ t('system.task.executeNow') }}
            </el-button>
            <el-button link type="primary" @click="handleEdit(row)">
              {{ t('system.task.edit') }}
            </el-button>
            <el-button
              link
              type="danger"
              @click="handleDelete(row)"
              :loading="deleteLoading[row.id]"
            >
              {{ t('system.task.delete') }}
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
import { useCommon } from '../../composables/useCommon'

const { t } = useCommon()
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
    ElMessage.error(t('system.task.loadDataFailed'))
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
      t('system.task.deleteConfirmMsg').replace('{name}', task.name),
      t('system.task.confirmDelete'),
      {
        confirmButtonText: t('system.task.confirm'),
        cancelButtonText: t('system.task.cancel'),
        type: 'warning'
      }
    )
    
    deleteLoading.value[task.id] = true
    await deleteScheduleTask(task.id)
    ElMessage.success(t('system.task.deleteSuccess'))
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error(t('system.task.deleteFailed'))
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
    task.status_text = active === 1 ? t('common.enabled') : t('common.disabled')
    const status = active === 1 ? t('common.enabled') : t('common.disabled')
    ElMessage.success(t('system.task.statusUpdateSuccess').replace('{status}', status))
    loadData() // 重新加载以更新下次执行时间
  } catch (error) {
    console.error('状态更新失败:', error)
    ElMessage.error(t('system.task.statusUpdateFailed'))
  } finally {
    statusLoading.value[task.id] = false
  }
}

// 立即执行
const handleExecute = async (task: ScheduleTask) => {
  try {
    executeLoading.value[task.id] = true
    await executeScheduleTask(task.id)
    ElMessage.success(t('system.task.executeSuccess'))
  } catch (error) {
    console.error('任务执行失败:', error)
    ElMessage.error(t('system.task.executeFailed'))
  } finally {
    executeLoading.value[task.id] = false
  }
}

// 批量操作
const handleBatchOperation = async (command: string) => {
  if (!selectedTasks.value.length) return
  
  const operationMap: Record<string, string> = {
    enable: t('system.task.batchEnable'),
    disable: t('system.task.batchDisable'),
    delete: t('system.task.batchDelete')
  }
  
  const operation = operationMap[command]
  if (!operation) return
  
  try {
    const message = t('system.task.batchOperationConfirm')
      .replace('{operation}', operation)
      .replace('{count}', selectedTasks.value.length.toString())
    await ElMessageBox.confirm(
      message,
      t('common.confirm'),
      {
        confirmButtonText: t('system.task.confirm'),
        cancelButtonText: t('system.task.cancel'),
        type: command === 'delete' ? 'warning' : 'info'
      }
    )
    
    const ids = selectedTasks.value.map(task => task.id)
    await batchOperateScheduleTasks({
      ids,
      operation: command as 'enable' | 'disable' | 'delete'
    })
    
    ElMessage.success(t('system.task.batchOperationSuccess').replace('{operation}', operation))
    selectedTasks.value = []
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(`批量操作失败:`, error)
      ElMessage.error(t('system.task.batchOperationFailed').replace('{operation}', operation))
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

.custom-table :deep(.el-table__header) {
  font-weight: 600;
}

.custom-table :deep(.el-table__row) {
  transition: all var(--transition-base);
}

.custom-table :deep(.el-table__row:hover) {
  transform: translateY(-1px);
}

.custom-table :deep(.el-table__cell) {
  padding: 12px 0;
}

.custom-table :deep(.el-button + .el-button) {
  margin-left: 8px;
}

.custom-table :deep(.el-switch) {
  --el-switch-on-color: var(--accent-color);
  --el-switch-off-color: rgba(255, 255, 255, 0.2);
}

.task-name {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.category-tag {
  font-size: 11px;
}

.next-run-time {
  color: var(--accent-color);
  font-weight: 500;
  font-size: 13px;
}

.custom-table :deep(.el-button--text) {
  padding: 4px 8px;
  font-size: 13px;
}

.custom-table :deep(.el-button--text.is-disabled) {
  opacity: 0.4;
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