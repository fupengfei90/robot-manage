<template>
  <div class="message-records-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">历史会话记录</h1>
        <p class="view-subtitle">查看和管理数字员工与用户的对话历史</p>
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
        <el-form-item label="用户">
          <el-input v-model="filters.request_user_name" placeholder="输入用户名" clearable style="width: 150px" />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="filters.types" placeholder="全部" multiple collapse-tags style="width: 200px">
            <el-option label="问答" value="question" />
            <el-option label="命令" value="command" />
            <el-option label="聊天" value="chat" />
            <el-option label="定时任务" value="scheduled_task" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始时间">
          <el-date-picker
            v-model="filters.startTime"
            type="datetime"
            placeholder="选择开始时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="结束时间">
          <el-date-picker
            v-model="filters.endTime"
            type="datetime"
            placeholder="选择结束时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            clearable
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="关键词">
          <el-input v-model="filters.keyword" placeholder="搜索对话内容" clearable style="width: 200px" />
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
        <el-table-column prop="request_user_name" label="用户" width="120" />
        <el-table-column label="类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.type === 'question'" type="primary" size="small">问答</el-tag>
            <el-tag v-else-if="row.type === 'command'" type="warning" size="small">命令</el-tag>
            <el-tag v-else-if="row.type === 'chat'" type="success" size="small">聊天</el-tag>
            <el-tag v-else-if="row.type === 'scheduled_task'" type="info" size="small">定时任务</el-tag>
            <el-tag v-else type="info" size="small">{{ row.type || '-' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="input" label="用户输入" min-width="200" show-overflow-tooltip />
        <el-table-column prop="output" label="AI回复" min-width="200" show-overflow-tooltip />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleView(row)">查看</el-button>
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
    <el-dialog v-model="dialogVisible" title="会话详情" width="1100px" class="custom-dialog">
      <div v-if="currentRecord" class="message-detail">
        <el-descriptions :column="1" border label-width="120px">
          <el-descriptions-item label="消息ID">{{ currentRecord.msg_id }}</el-descriptions-item>
          <el-descriptions-item label="会话ID">{{ currentRecord.conversation_id }}</el-descriptions-item>
          <el-descriptions-item label="用户">{{ currentRecord.request_user_name }}</el-descriptions-item>
          <el-descriptions-item label="用户ID">{{ currentRecord.request_user_id }}</el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ formatDateTime(currentRecord.created_at) }}</el-descriptions-item>
          <el-descriptions-item label="备注">{{ currentRecord.extra_data || '-' }}</el-descriptions-item>
        </el-descriptions>
        <div class="message-content">
          <h4>用户输入：</h4>
          <div class="content-box">{{ currentRecord.input }}</div>
          <h4>AI回复：</h4>
          <div class="content-box">{{ currentRecord.output }}</div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useCommon } from '../../composables/useCommon'

const { t } = useCommon()
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getMessageRecords,
  deleteMessageRecord,
  type MessageRecord,
  type MessageRecordQuery
} from '@/api/digital-employee'

const loading = ref(false)
const tableData = ref<MessageRecord[]>([])
const dialogVisible = ref(false)
const currentRecord = ref<MessageRecord | null>(null)

const filters = reactive({
  conversation_id: '',
  request_user_name: '',
  types: ['question', 'command', 'chat'] as string[],
  is_valid: 1,
  keyword: '',
  startTime: '',
  endTime: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const loadData = async () => {
  loading.value = true
  try {
    const params: any = {
      conversation_id: filters.conversation_id,
      request_user_name: filters.request_user_name,
      type: filters.types.length > 0 ? filters.types.join(',') + ',null' : '',
      is_valid: filters.is_valid,
      keyword: filters.keyword,
      start_time: filters.startTime,
      end_time: filters.endTime,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    const res = await getMessageRecords(params)
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
  filters.conversation_id = ''
  filters.request_user_name = ''
  filters.types = ['question', 'command', 'chat']
  filters.is_valid = 1
  filters.keyword = ''
  filters.startTime = ''
  filters.endTime = ''
  handleSearch()
}

const handleView = (row: MessageRecord) => {
  currentRecord.value = row
  dialogVisible.value = true
}

const handleDelete = async (row: MessageRecord) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除这条会话记录吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await deleteMessageRecord(row.msg_id)
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
.message-records-view {
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

.message-detail {
  padding: var(--spacing-md);
}

.message-content {
  margin-top: var(--spacing-lg);
}

.message-content h4 {
  margin: var(--spacing-md) 0 var(--spacing-sm) 0;
  color: var(--text-primary);
}

.content-box {
  padding: var(--spacing-md);
  background: var(--bg-tertiary);
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
  white-space: pre-wrap;
  word-break: break-word;
  line-height: 1.6;
}

.custom-dialog :deep(.el-dialog) {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
}
</style>
