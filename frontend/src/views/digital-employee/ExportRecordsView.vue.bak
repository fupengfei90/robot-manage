<template>
  <div class="export-records-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">服务回传记录</h1>
        <p class="view-subtitle">查看和管理数据导出和邮件发送记录</p>
      </div>
      <div class="header-actions">
        <el-button type="primary" @click="loadData" class="refresh-btn">
          <span>🔄</span> 刷新
        </el-button>
      </div>
    </div>

    <!-- 搜索筛选 -->
    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filters" class="filter-form">
        <el-form-item label="文件名">
          <el-input v-model="filters.file_name" placeholder="输入文件名" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item label="邮件状态">
          <el-select v-model="filters.email_status" placeholder="全部" clearable style="width: 120px">
            <el-option label="成功" value="success" />
            <el-option label="失败" value="failed" />
            <el-option label="待发送" value="pending" />
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
        v-loading="loading"
        :data="tableData"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="file_name" label="文件名" min-width="200" show-overflow-tooltip />
        <el-table-column label="日期范围" width="200">
          <template #default="{ row }">
            {{ row.start_date }} ~ {{ row.end_date }}
          </template>
        </el-table-column>
        <el-table-column prop="total_records" label="记录数" width="100" />
        <el-table-column label="邮件状态" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.email_status === 'success'" type="success">成功</el-tag>
            <el-tag v-else-if="row.email_status === 'failed'" type="danger">失败</el-tag>
            <el-tag v-else type="info">待发送</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="email_subject" label="邮件主题" min-width="150" show-overflow-tooltip />
        <el-table-column prop="export_time" label="导出时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.export_time) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleView(row)">查看</el-button>
            <el-button 
              v-if="row.email_status === 'failed'" 
              link 
              type="warning" 
              @click="handleRetry(row)"
            >
              重试
            </el-button>
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

    <!-- 详情对话框 -->
    <el-dialog v-model="dialogVisible" title="导出记录详情" width="800px" class="custom-dialog">
      <div v-if="currentRecord" class="export-detail">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="ID">{{ currentRecord.id }}</el-descriptions-item>
          <el-descriptions-item label="导出时间">{{ formatDateTime(currentRecord.export_time) }}</el-descriptions-item>
          <el-descriptions-item label="开始日期">{{ currentRecord.start_date }}</el-descriptions-item>
          <el-descriptions-item label="结束日期">{{ currentRecord.end_date }}</el-descriptions-item>
          <el-descriptions-item label="总记录数">{{ currentRecord.total_records }}</el-descriptions-item>
          <el-descriptions-item label="邮件状态">
            <el-tag v-if="currentRecord.email_status === 'success'" type="success">成功</el-tag>
            <el-tag v-else-if="currentRecord.email_status === 'failed'" type="danger">失败</el-tag>
            <el-tag v-else type="info">待发送</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="文件名" :span="2">{{ currentRecord.file_name }}</el-descriptions-item>
          <el-descriptions-item label="文件路径" :span="2">{{ currentRecord.file_path }}</el-descriptions-item>
          <el-descriptions-item label="邮件主题" :span="2">{{ currentRecord.email_subject }}</el-descriptions-item>
          <el-descriptions-item label="收件人" :span="2">{{ currentRecord.email_recipients }}</el-descriptions-item>
          <el-descriptions-item v-if="currentRecord.error_reason" label="错误原因" :span="2">
            <span style="color: var(--color-danger)">{{ currentRecord.error_reason }}</span>
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getExportRecords,
  deleteExportRecord,
  retryExportEmail,
  type ExportRecord,
  type ExportRecordQuery
} from '@/api/digital-employee'

const loading = ref(false)
const tableData = ref<ExportRecord[]>([])
const dialogVisible = ref(false)
const currentRecord = ref<ExportRecord | null>(null)

const filters = reactive<ExportRecordQuery>({
  file_name: '',
  email_status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...filters,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getExportRecords(params)
    tableData.value = res.list || []
    pagination.total = res.total || 0
  } catch (error: any) {
    ElMessage.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  filters.file_name = ''
  filters.email_status = ''
  handleSearch()
}

const handleView = (row: ExportRecord) => {
  currentRecord.value = row
  dialogVisible.value = true
}

const handleRetry = async (row: ExportRecord) => {
  try {
    await ElMessageBox.confirm(
      `确定要重试发送邮件吗？`,
      '确认重试',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await retryExportEmail(row.id)
    ElMessage.success('已加入重试队列')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '重试失败')
    }
  }
}

const handleDelete = async (row: ExportRecord) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除这条导出记录吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await deleteExportRecord(row.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

const handlePageChange = (page: number) => {
  pagination.page = page
  loadData()
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  loadData()
}

const formatDateTime = (date: string) => {
  if (!date) return ''
  return new Date(date).toLocaleString('zh-CN')
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.export-records-view {
  padding: var(--spacing-md);
}

.view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: var(--radius-xl);
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

.custom-table :deep(.el-table__row:hover) {
  background: var(--bg-glass-hover);
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

.export-detail {
  padding: var(--spacing-md);
}

.custom-dialog :deep(.el-dialog) {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
}
</style>
