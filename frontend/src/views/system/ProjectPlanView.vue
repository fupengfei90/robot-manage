<template>
  <div class="project-plan-view">
    <div class="header-summary glass-effect">
      <div class="header-left">
        <h2>项目规划</h2>
        <p>支持模块管理、任务跟踪、进度可视化</p>
      </div>
      <div class="header-stats">
        <div class="summary-item">
          <div class="summary-value">{{ completedCount }}/{{ totalCount }}</div>
          <div class="summary-label">已完成任务</div>
        </div>
        <div class="summary-item">
          <div class="summary-value">{{ progressPercentage }}%</div>
          <div class="summary-label">总体进度</div>
        </div>
        <div class="summary-item">
          <div class="summary-value">{{ totalPlannedHours }}h</div>
          <div class="summary-label">计划工时</div>
        </div>
        <div class="summary-item">
          <div class="summary-value">{{ totalActualHours }}h</div>
          <div class="summary-label">实际工时</div>
        </div>
      </div>
    </div>

    <div class="toolbar glass-effect">
      <el-button type="primary" @click="showAddModuleDialog">
        <el-icon><Plus /></el-icon>
        添加模块
      </el-button>
      <el-button @click="expandAll">展开全部</el-button>
      <el-button @click="collapseAll">收起全部</el-button>
    </div>

    <div class="modules-container">
      <div v-for="module in modules" :key="module.id" class="module-card glass-effect">
        <div class="module-header">
          <div class="module-left">
            <el-icon
              class="expand-icon"
              :class="{ 'is-expanded': expandedKeys.has(module.id) }"
              @click="toggleExpand(module.id)"
            >
              <ArrowRight v-if="module.items.length > 0" />
              <span v-else style="width: 16px; display: inline-block;"></span>
            </el-icon>
            <span class="module-icon">{{ module.icon }}</span>
            <span class="module-name">{{ module.name }}</span>
            <el-tag :type="getModuleStatusType(module)" size="small">
              {{ getModuleStatusText(module) }}
            </el-tag>
            <span class="module-stats">{{ getCompletedItems(module) }}/{{ module.items.length }}</span>
          </div>
          <div class="module-actions">
            <el-button size="small" text @click="showAddItemDialog(module)">
              <el-icon><Plus /></el-icon>
              任务
            </el-button>
            <el-button size="small" text @click="showEditModuleDialog(module)">
              <el-icon><Edit /></el-icon>
            </el-button>
            <el-button size="small" text type="danger" @click="handleDeleteModule(module)">
              <el-icon><Delete /></el-icon>
            </el-button>
          </div>
        </div>

        <el-collapse-transition>
          <div v-show="expandedKeys.has(module.id)" class="module-content">
            <draggable
              v-if="module.items.length > 0"
              v-model="module.items"
              item-key="id"
              class="items-list"
              handle=".drag-handle"
              animation="200"
              @end="handleDragEnd(module)"
            >
              <template #item="{ element: item }">
              <div class="item-row">
                <div class="item-info">
                  <span class="drag-handle">☰</span>
                  <span class="item-name">{{ item.name }}</span>
                  <el-tag :type="getStatusType(item.status)" size="small">
                    {{ getStatusText(item.status) }}
                  </el-tag>
                  <el-tag :type="getPriorityType(item.priority)" size="small">
                    {{ getPriorityText(item.priority) }}
                  </el-tag>
                  <span v-if="item.owner" class="item-owner">👤 {{ item.owner }}</span>
                  <span v-if="item.planned_hours" class="item-hours">
                    📅 {{ item.planned_hours }}h
                  </span>
                  <span v-if="item.actual_hours" class="item-hours">
                    ⏱️ {{ item.actual_hours }}h
                  </span>
                </div>
                <div class="item-actions">
                  <el-select
                    v-model="item.status"
                    size="small"
                    class="quick-select status-select"
                    @change="handleUpdateItem(item)"
                  >
                    <el-option label="未开始" value="not_started">
                      <span class="option-dot dot-gray"></span>未开始
                    </el-option>
                    <el-option label="进行中" value="in_progress">
                      <span class="option-dot dot-blue"></span>进行中
                    </el-option>
                    <el-option label="已完成" value="completed">
                      <span class="option-dot dot-green"></span>已完成
                    </el-option>
                    <el-option label="阻塞" value="blocked">
                      <span class="option-dot dot-red"></span>阻塞
                    </el-option>
                    <el-option label="已取消" value="cancelled">
                      <span class="option-dot dot-gray"></span>已取消
                    </el-option>
                  </el-select>
                  <el-select
                    v-model="item.priority"
                    size="small"
                    class="quick-select priority-select"
                    @change="handleUpdateItem(item)"
                  >
                    <el-option label="低" value="low">
                      <span class="priority-icon">🟢</span>低
                    </el-option>
                    <el-option label="中" value="medium">
                      <span class="priority-icon">🟡</span>中
                    </el-option>
                    <el-option label="高" value="high">
                      <span class="priority-icon">🟠</span>高
                    </el-option>
                    <el-option label="紧急" value="urgent">
                      <span class="priority-icon">🔴</span>紧急
                    </el-option>
                  </el-select>
                  <el-button size="small" text @click="showEditItemDialog(item)">
                    <el-icon><Edit /></el-icon>
                  </el-button>
                  <el-button size="small" text type="danger" @click="handleDeleteItem(item)">
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </div>
              </div>
              </template>
            </draggable>
          </div>
        </el-collapse-transition>
      </div>
    </div>

    <el-dialog v-model="moduleDialogVisible" :title="moduleDialogTitle" width="500px">
      <el-form :model="moduleForm" label-width="100px">
        <el-form-item label="模块名称" required>
          <el-input v-model="moduleForm.name" placeholder="请输入模块名称" />
        </el-form-item>
        <el-form-item label="图标">
          <el-input v-model="moduleForm.icon" placeholder="请输入图标 emoji" maxlength="10" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="moduleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSaveModule">保存</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="itemDialogVisible" :title="itemDialogTitle" width="600px">
      <el-form :model="itemForm" label-width="100px">
        <el-form-item label="任务名称" required>
          <el-input v-model="itemForm.name" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="itemForm.status" style="width: 100%">
            <el-option label="未开始" value="not_started" />
            <el-option label="进行中" value="in_progress" />
            <el-option label="已完成" value="completed" />
            <el-option label="阻塞" value="blocked" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
        <el-form-item label="优先级">
          <el-select v-model="itemForm.priority" style="width: 100%">
            <el-option label="低" value="low" />
            <el-option label="中" value="medium" />
            <el-option label="高" value="high" />
            <el-option label="紧急" value="urgent" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="itemForm.owner" placeholder="请输入负责人" />
        </el-form-item>
        <el-form-item label="计划工时">
          <el-input-number v-model="itemForm.planned_hours" :min="0" :step="0.5" placeholder="小时" style="width: 100%" />
        </el-form-item>
        <el-form-item label="实际工时">
          <el-input-number v-model="itemForm.actual_hours" :min="0" :step="0.5" placeholder="小时" style="width: 100%" />
        </el-form-item>
        <el-form-item label="详细描述">
          <el-input v-model="itemForm.description" type="textarea" :rows="4" placeholder="请输入详细描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="itemDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSaveItem">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { Plus, Edit, Delete, ArrowRight } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import draggable from 'vuedraggable'
import {
  getModulesWithItems,
  createModule,
  updateModule,
  deleteModule,
  createItem,
  updateItem,
  deleteItem,
  batchUpdateItems,
  type ModuleWithItems,
  type ProjectModule,
  type ProjectItem
} from '@/api/project-plan'

const modules = ref<ModuleWithItems[]>([])
const expandedKeys = ref<Set<number>>(new Set())

const moduleDialogVisible = ref(false)
const moduleDialogTitle = ref('添加模块')
const moduleForm = ref<Partial<ProjectModule>>({
  name: '',
  icon: '📦'
})

const itemDialogVisible = ref(false)
const itemDialogTitle = ref('添加任务')
const itemForm = ref<Partial<ProjectItem>>({
  name: '',
  status: 'not_started',
  priority: 'medium',
  owner: '',
  planned_hours: null,
  actual_hours: null,
  description: '',
  module_id: 0
})

const totalCount = computed(() => {
  return modules.value.reduce((sum, module) => sum + module.items.length, 0)
})

const completedCount = computed(() => {
  return modules.value.reduce((sum, module) => 
    sum + module.items.filter(item => item.status === 'completed').length, 0
  )
})

const progressPercentage = computed(() => {
  return totalCount.value > 0 ? Math.round((completedCount.value / totalCount.value) * 100) : 0
})

const totalPlannedHours = computed(() => {
  const total = modules.value.reduce((sum, module) => 
    sum + module.items.reduce((s, item) => s + (item.planned_hours || 0), 0), 0
  )
  return Math.round(total * 10) / 10
})

const totalActualHours = computed(() => {
  const total = modules.value.reduce((sum, module) => 
    sum + module.items.reduce((s, item) => s + (item.actual_hours || 0), 0), 0
  )
  return Math.round(total * 10) / 10
})

const loadData = async () => {
  try {
    const data = await getModulesWithItems()
    modules.value = data
  } catch (error) {
    console.error('加载失败:', error)
    ElMessage.error('加载数据失败')
  }
}

const toggleExpand = (id: number) => {
  if (expandedKeys.value.has(id)) {
    expandedKeys.value.delete(id)
  } else {
    expandedKeys.value.add(id)
  }
}

const expandAll = () => {
  expandedKeys.value = new Set(modules.value.map(m => m.id))
}

const collapseAll = () => {
  expandedKeys.value.clear()
}

const getCompletedItems = (module: ModuleWithItems) => {
  return module.items.filter(item => item.status === 'completed').length
}

const getModuleStatusType = (module: ModuleWithItems) => {
  const completed = getCompletedItems(module)
  const total = module.items.length
  if (total === 0) return ''
  const percentage = (completed / total) * 100
  if (percentage === 100) return 'success'
  if (percentage >= 50) return 'warning'
  if (percentage > 0) return 'info'
  return ''
}

const getModuleStatusText = (module: ModuleWithItems) => {
  const completed = getCompletedItems(module)
  const total = module.items.length
  if (total === 0) return '无任务'
  const percentage = (completed / total) * 100
  if (percentage === 100) return '已完成'
  if (percentage >= 50) return '进行中'
  if (percentage > 0) return '开发中'
  return '未开始'
}

const getStatusType = (status: string) => {
  const map: Record<string, any> = {
    not_started: '',
    in_progress: 'warning',
    completed: 'success',
    blocked: 'danger',
    cancelled: 'info'
  }
  return map[status] || ''
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    not_started: '未开始',
    in_progress: '进行中',
    completed: '已完成',
    blocked: '阻塞',
    cancelled: '已取消'
  }
  return map[status] || status
}

const getPriorityType = (priority: string) => {
  const map: Record<string, any> = {
    low: 'info',
    medium: '',
    high: 'warning',
    urgent: 'danger'
  }
  return map[priority] || ''
}

const getPriorityText = (priority: string) => {
  const map: Record<string, string> = {
    low: '低',
    medium: '中',
    high: '高',
    urgent: '紧急'
  }
  return map[priority] || priority
}

const showAddModuleDialog = () => {
  moduleDialogTitle.value = '添加模块'
  moduleForm.value = {
    name: '',
    icon: '📦'
  }
  moduleDialogVisible.value = true
}

const showEditModuleDialog = (module: ModuleWithItems) => {
  moduleDialogTitle.value = '编辑模块'
  moduleForm.value = { ...module }
  moduleDialogVisible.value = true
}

const handleSaveModule = async () => {
  if (!moduleForm.value.name) {
    ElMessage.warning('请输入模块名称')
    return
  }

  try {
    if (moduleForm.value.id) {
      await updateModule(moduleForm.value.id, moduleForm.value)
      ElMessage.success('更新成功')
    } else {
      await createModule(moduleForm.value)
      ElMessage.success('创建成功')
    }
    moduleDialogVisible.value = false
    await loadData()
  } catch (error) {
    console.error('保存失败:', error)
    ElMessage.error('保存失败')
  }
}

const handleDeleteModule = async (module: ModuleWithItems) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除模块"${module.name}"吗？这将同时删除所有任务。`,
      '确认删除',
      { type: 'warning' }
    )
    await deleteModule(module.id)
    ElMessage.success('删除成功')
    await loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

const showAddItemDialog = (module: ModuleWithItems) => {
  itemDialogTitle.value = '添加任务'
  itemForm.value = {
    name: '',
    status: 'not_started',
    priority: 'medium',
    owner: '',
    planned_hours: null,
    actual_hours: null,
    description: '',
    module_id: module.id
  }
  itemDialogVisible.value = true
}

const showEditItemDialog = (item: ProjectItem) => {
  itemDialogTitle.value = '编辑任务'
  itemForm.value = { ...item }
  itemDialogVisible.value = true
}

const handleSaveItem = async () => {
  if (!itemForm.value.name) {
    ElMessage.warning('请输入任务名称')
    return
  }

  try {
    if (itemForm.value.id) {
      await updateItem(itemForm.value.id, itemForm.value)
      ElMessage.success('更新成功')
    } else {
      await createItem(itemForm.value)
      ElMessage.success('创建成功')
    }
    itemDialogVisible.value = false
    await loadData()
  } catch (error) {
    console.error('保存失败:', error)
    ElMessage.error('保存失败')
  }
}

const handleDeleteItem = async (item: ProjectItem) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除任务"${item.name}"吗？`,
      '确认删除',
      { type: 'warning' }
    )
    await deleteItem(item.id)
    ElMessage.success('删除成功')
    await loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

const handleUpdateItem = async (item: ProjectItem) => {
  try {
    await updateItem(item.id, item)
  } catch (error) {
    console.error('更新失败:', error)
    ElMessage.error('更新失败')
  }
}

const handleDragEnd = async (module: ModuleWithItems) => {
  try {
    module.items.forEach((item, index) => {
      item.sort_order = index
    })
    await batchUpdateItems(module.items)
  } catch (error) {
    console.error('排序失败:', error)
    ElMessage.error('保存排序失败')
    await loadData()
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.project-plan-view {
  padding: var(--spacing-md);
}

.header-summary {
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-md);
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: var(--spacing-lg);
}

.header-left h2 {
  font-size: 1.5rem;
  font-weight: 700;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: var(--spacing-xs);
}

.header-left p {
  color: var(--text-muted);
  font-size: 0.875rem;
}

.header-stats {
  display: flex;
  gap: var(--spacing-lg);
  align-items: center;
}

.summary-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.summary-value {
  font-size: 1.5rem;
  font-weight: 700;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.summary-label {
  color: var(--text-muted);
  font-size: 0.75rem;
}

.toolbar {
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-md);
  display: flex;
  gap: var(--spacing-sm);
}

.modules-container {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.module-card {
  border-radius: var(--radius-md);
  overflow: hidden;
  margin-bottom: var(--spacing-xs);
}

.module-header {
  padding: var(--spacing-sm) var(--spacing-md);
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: background var(--transition-base);
}

.module-header:hover {
  background: var(--bg-glass-hover);
}

.module-left {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  flex: 1;
}

.expand-icon {
  transition: transform var(--transition-base);
  cursor: pointer;
}

.expand-icon.is-expanded {
  transform: rotate(90deg);
}

.module-icon {
  font-size: 1.25rem;
}

.module-name {
  font-weight: 600;
  color: var(--text-primary);
}

.module-stats {
  font-weight: 600;
  color: var(--accent-color);
  font-size: 0.875rem;
  margin-left: auto;
  margin-right: var(--spacing-md);
}

.module-actions {
  display: flex;
  gap: var(--spacing-xs);
}

.module-content {
  padding: 0 var(--spacing-md) var(--spacing-sm) calc(var(--spacing-md) + 32px);
}

.items-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
  margin-bottom: var(--spacing-sm);
}

.item-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-xs) var(--spacing-sm);
  background: var(--bg-glass);
  border-radius: var(--radius-sm);
  transition: all var(--transition-base);
}

.item-row:hover {
  background: var(--bg-glass-hover);
  transform: translateX(2px);
}

.item-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  flex: 1;
}

.item-name {
  color: var(--text-primary);
  font-size: 0.875rem;
}

.item-owner,
.item-hours {
  color: var(--text-muted);
  font-size: 0.75rem;
}

.item-actions {
  display: flex;
  gap: var(--spacing-xs);
  align-items: center;
}

.quick-select {
  transition: all var(--transition-base);
}

.quick-select :deep(.el-input__wrapper) {
  background: var(--bg-glass);
  border: 1px solid transparent;
  box-shadow: none;
  transition: all var(--transition-base);
}

.quick-select:hover :deep(.el-input__wrapper) {
  background: var(--bg-glass-hover);
  border-color: var(--accent-color);
}

.quick-select :deep(.el-input__wrapper.is-focus) {
  background: var(--bg-glass-hover);
  border-color: var(--accent-color);
  box-shadow: 0 0 0 2px rgba(var(--accent-rgb), 0.1);
}

.status-select {
  width: 110px;
}

.priority-select {
  width: 95px;
}

.option-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 8px;
}

.dot-gray { background: #909399; }
.dot-blue { background: #409eff; }
.dot-green { background: #67c23a; }
.dot-red { background: #f56c6c; }

.priority-icon {
  margin-right: 6px;
  font-size: 12px;
}

.drag-handle {
  cursor: move;
  color: var(--text-muted);
  margin-right: var(--spacing-xs);
  opacity: 0.4;
  transition: opacity var(--transition-base);
}

.item-row:hover .drag-handle {
  opacity: 1;
}
</style>
