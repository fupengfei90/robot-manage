<template>
  <div class="batch-time-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">批量时间管理</h1>
        <p class="view-subtitle">管理系统批处理时间窗口配置</p>
      </div>
      <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
        <span>➕</span> 新增配置
      </el-button>
    </div>

    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filterForm" class="filter-form">
        <el-form-item label="系统名称">
          <el-input
            v-model="filterForm.systemName"
            placeholder="搜索系统名称"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="子系统名称">
          <el-input
            v-model="filterForm.subsysName"
            placeholder="搜索子系统名称"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="批量时间">
          <el-input
            v-model="filterForm.batchTime"
            placeholder="搜索批量时间"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        :data="tableData"
        v-loading="loading"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="systemName" label="系统名称" min-width="150" />
        <el-table-column prop="subsysName" label="子系统名称" min-width="150" />
        <el-table-column prop="batchTime" label="批量时间窗口" min-width="180">
          <template #default="{ row }">
            <span class="time-cell">{{ row.batchTime }}</span>
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
        <el-form-item label="系统名称" prop="systemName">
          <el-input
            v-model="formData.systemName"
            placeholder="请输入系统名称"
          />
        </el-form-item>
        <el-form-item label="子系统名称" prop="subsysName">
          <el-input
            v-model="formData.subsysName"
            placeholder="请输入子系统名称"
          />
        </el-form-item>
        <el-form-item label="批量时间" prop="batchTime">
          <el-input
            v-model="formData.batchTime"
            placeholder="例如: 22:00-23:00"
          />
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
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import {
  getBatchTimes,
  createBatchTime,
  updateBatchTime,
  deleteBatchTime
} from '../../api/cmdb'
import type { BatchTime, BatchTimeCreateRequest } from '../../types/cmdb'

const loading = ref(false)
const submitting = ref(false)
const dialogVisible = ref(false)
const formRef = ref<FormInstance>()
const currentEditId = ref<number | null>(null)

const filterForm = reactive({
  systemName: '',
  subsysName: '',
  batchTime: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const tableData = ref<BatchTime[]>([])

const formData = reactive<BatchTimeCreateRequest>({
  systemName: '',
  subsysName: '',
  batchTime: ''
})

const formRules: FormRules = {
  systemName: [
    { required: true, message: '请输入系统名称', trigger: 'blur' }
  ],
  subsysName: [
    { required: true, message: '请输入子系统名称', trigger: 'blur' }
  ],
  batchTime: [
    { required: true, message: '请输入批量时间', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => currentEditId.value ? '编辑批量时间' : '新增批量时间')

const formatDateTime = (dateStr: string) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  return d.toLocaleString('zh-CN')
}

const fetchData = async () => {
  loading.value = true
  try {
    const res = await getBatchTimes({
      systemName: filterForm.systemName || undefined,
      subsysName: filterForm.subsysName || undefined,
      batchTime: filterForm.batchTime || undefined,
      page: pagination.page,
      pageSize: pagination.pageSize
    })
    tableData.value = res.list || []
    pagination.total = res.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '获取数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  fetchData()
}

const handleReset = () => {
  filterForm.systemName = ''
  filterForm.subsysName = ''
  filterForm.batchTime = ''
  handleSearch()
}

const handlePageChange = (page: number) => {
  pagination.page = page
  fetchData()
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  fetchData()
}

const handleCreate = () => {
  currentEditId.value = null
  formData.systemName = ''
  formData.subsysName = ''
  formData.batchTime = ''
  formRef.value?.clearValidate()
  dialogVisible.value = true
}

const handleEdit = (row: BatchTime) => {
  currentEditId.value = row.id
  formData.systemName = row.systemName
  formData.subsysName = row.subsysName
  formData.batchTime = row.batchTime
  dialogVisible.value = true
}

const handleDelete = async (row: BatchTime) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除系统"${row.systemName}"的批量时间配置吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await deleteBatchTime(row.id)
    ElMessage.success('删除成功')
    fetchData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (!valid) return

    submitting.value = true
    try {
      if (currentEditId.value) {
        await updateBatchTime(currentEditId.value, formData)
        ElMessage.success('更新成功')
      } else {
        await createBatchTime(formData)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      fetchData()
    } catch (error: any) {
      ElMessage.error(error.message || '操作失败')
    } finally {
      submitting.value = false
    }
  })
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.batch-time-view {
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

.time-cell {
  font-weight: 600;
  color: var(--accent-color);
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
