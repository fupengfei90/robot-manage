<template>
  <div class="milestone-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">大事记管理</h1>
        <p class="view-subtitle">记录和管理重要事件信息</p>
      </div>
      <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
        <span>➕</span> 新增大事记
      </el-button>
    </div>

    <!-- 搜索筛选 -->
    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filters" class="filter-form">
        <el-form-item label="开始日期">
          <el-date-picker
            v-model="filters.startDate"
            type="date"
            placeholder="选择开始日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            clearable
          />
        </el-form-item>
        <el-form-item label="结束日期">
          <el-date-picker
            v-model="filters.endDate"
            type="date"
            placeholder="选择结束日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            clearable
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="filters.isActive" placeholder="全部" clearable style="width: 120px">
            <el-option label="激活" :value="true" />
            <el-option label="禁用" :value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 数据表格 -->
    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        :data="tableData"
        v-loading="loading"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="eventDate" label="事件日期" width="160">
          <template #default="{ row }">
            <div>
              <span class="date-cell">{{ formatDate(row.eventDate) }}</span>
              <el-tag size="small" type="info" style="margin-left: 8px">
                {{ row.dayOfWeek }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="eventContent" label="事件内容" min-width="400">
          <template #default="{ row }">
            <div class="content-cell">{{ row.eventContent }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="isActive" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.isActive ? 'success' : 'info'">
              {{ row.isActive ? '激活' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="170">
          <template #default="{ row }">
            {{ formatDateTime(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column prop="updatedAt" label="更新时间" width="170">
          <template #default="{ row }">
            {{ formatDateTime(row.updatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
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
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="700px"
      class="custom-dialog"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <el-form-item label="事件日期" prop="eventDate">
          <el-date-picker
            v-model="formData.eventDate"
            type="date"
            placeholder="选择事件日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="事件内容" prop="eventContent">
          <el-input
            v-model="formData.eventContent"
            type="textarea"
            :rows="5"
            placeholder="请输入事件内容"
            maxlength="1000"
            show-word-limit
          />
        </el-form-item>
        <el-form-item v-if="isEdit" label="状态">
          <el-switch v-model="formData.isActive" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitting">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import {
  getMilestoneEvents,
  createMilestoneEvent,
  updateMilestoneEvent,
  deleteMilestoneEvent
} from '../../api/cmdb'
import type {
  MilestoneEvent,
  MilestoneEventCreateRequest,
  MilestoneEventUpdateRequest
} from '../../types/cmdb'

const loading = ref(false)
const submitting = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const currentId = ref<number | null>(null)
const formRef = ref<FormInstance>()

const filters = reactive({
  startDate: '',
  endDate: '',
  isActive: undefined as boolean | undefined
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const tableData = ref<MilestoneEvent[]>([])

const formData = reactive<MilestoneEventCreateRequest & { isActive?: boolean }>({
  eventDate: '',
  eventContent: '',
  isActive: true
})

const formRules: FormRules = {
  eventDate: [{ required: true, message: '请选择事件日期', trigger: 'change' }],
  eventContent: [
    { required: true, message: '请输入事件内容', trigger: 'blur' },
    { min: 1, max: 1000, message: '事件内容长度在1到1000个字符', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => (isEdit.value ? '编辑大事记' : '新增大事记'))

// 加载数据
const loadData = async () => {
  loading.value = true
  try {
    const res = await getMilestoneEvents({
      startDate: filters.startDate || undefined,
      endDate: filters.endDate || undefined,
      isActive: filters.isActive,
      page: pagination.page,
      pageSize: pagination.pageSize
    })
    tableData.value = res.list
    pagination.total = res.total
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  loadData()
}

// 重置
const handleReset = () => {
  filters.startDate = ''
  filters.endDate = ''
  filters.isActive = undefined
  handleSearch()
}

// 新增
const handleCreate = () => {
  isEdit.value = false
  currentId.value = null
  formData.eventDate = ''
  formData.eventContent = ''
  formData.isActive = true
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: MilestoneEvent) => {
  isEdit.value = true
  currentId.value = row.id
  formData.eventDate = row.eventDate
  formData.eventContent = row.eventContent
  formData.isActive = row.isActive
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: MilestoneEvent) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除这条大事记吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await deleteMilestoneEvent(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (!valid) return

    submitting.value = true
    try {
      if (isEdit.value && currentId.value) {
        const updateData: MilestoneEventUpdateRequest = {
          eventDate: formData.eventDate,
          eventContent: formData.eventContent,
          isActive: formData.isActive
        }
        await updateMilestoneEvent(currentId.value, updateData)
        ElMessage.success('更新成功')
      } else {
        const createData: MilestoneEventCreateRequest = {
          eventDate: formData.eventDate,
          eventContent: formData.eventContent
        }
        await createMilestoneEvent(createData)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadData()
    } catch (error: any) {
      ElMessage.error(error.message || '操作失败')
    } finally {
      submitting.value = false
    }
  })
}

// 分页
const handlePageChange = (page: number) => {
  pagination.page = page
  loadData()
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  loadData()
}

// 格式化日期
const formatDate = (date: string) => {
  return date ? date.split('T')[0] : ''
}

const formatDateTime = (date: string) => {
  if (!date) return ''
  const d = new Date(date)
  return d.toLocaleString('zh-CN')
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.milestone-view {
  padding: var(--spacing-md);
}

.view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-lg);
  margin-bottom: var(--spacing-md);
  border-radius: var(--radius-xl);
}

.view-title {
  font-size: 1.75rem;
  font-weight: 700;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.view-subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  margin-top: var(--spacing-xs);
}

.create-btn {
  padding: var(--spacing-md) var(--spacing-xl);
  border-radius: var(--radius-md);
}

.filter-card {
  margin-bottom: var(--spacing-md);
  border-radius: var(--radius-xl);
}

.filter-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}

.filter-form {
  margin: 0;
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

.date-cell {
  font-weight: 600;
  color: var(--accent-color);
}

.content-cell {
  line-height: 1.6;
  color: var(--text-secondary);
  word-break: break-word;
  white-space: normal;
  max-width: 100%;
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

.custom-dialog :deep(.el-dialog) {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
}

.custom-dialog :deep(.el-dialog__header) {
  border-bottom: 1px solid var(--border-color);
}

.custom-dialog :deep(.el-dialog__title) {
  color: var(--text-primary);
  font-weight: 600;
}
</style>

