<template>
  <div class="duty-schedule-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">{{ t('cmdb.duty.title') }}</h1>
        <p class="view-subtitle">{{ t('cmdb.duty.subtitle') }}</p>
      </div>
      <div class="header-actions">
        <el-radio-group v-model="viewMode" size="default">
          <el-radio-button value="list">列表视图</el-radio-button>
          <el-radio-button value="calendar">月历视图</el-radio-button>
        </el-radio-group>
        <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
          <span>➕</span> {{ t('cmdb.duty.addSchedule') }}
        </el-button>
      </div>
    </div>

    <!-- 搜索筛选 -->
    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filters" class="filter-form">
        <el-form-item :label="t('cmdb.duty.startDate')">
          <el-date-picker
            v-model="filters.startDate"
            type="date"
            placeholder="选择开始日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            clearable
          />
        </el-form-item>
        <el-form-item :label="t('cmdb.duty.endDate')">
          <el-date-picker
            v-model="filters.endDate"
            type="date"
            placeholder="选择结束日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            clearable
          />
        </el-form-item>
        <el-form-item :label="t('cmdb.duty.staffName')">
          <el-input v-model="filters.staffName" placeholder="输入值班人姓名" clearable style="width: 150px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">{{ t('common.search') }}</el-button>
          <el-button @click="handleReset">{{ t('common.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 月历视图 -->
    <el-card v-if="viewMode === 'calendar'" shadow="never" class="calendar-card glass-effect animate-fade-in-up">
      <MonthCalendar :items="calendarItems">
        <template #default="{ items }">
          <div v-for="item in items" :key="item.id" class="calendar-item">
            <div class="duty-row">
              <span class="duty-label">WB:</span>
              <span class="duty-name">{{ item.wbStaffName }}</span>
            </div>
            <div class="duty-row">
              <span class="duty-label">FB:</span>
              <span class="duty-name">{{ item.fbStaffName }}</span>
            </div>
          </div>
        </template>
      </MonthCalendar>
    </el-card>

    <!-- 数据表格 -->
    <el-card v-else shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        :data="tableData"
        v-loading="loading"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="dutyDate" :label="t('cmdb.duty.dutyDate')" width="140">
          <template #default="{ row }">
            <span class="date-cell">{{ formatDate(row.dutyDate) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="wbStaffName" :label="t('cmdb.duty.wbStaff')" min-width="200">
          <template #default="{ row }">
            <el-tag type="primary" class="staff-tag">{{ row.wbStaffName }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="fbStaffName" :label="t('cmdb.duty.fbStaff')" min-width="200">
          <template #default="{ row }">
            <el-tag type="warning" class="staff-tag">{{ row.fbStaffName }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" :label="t('cmdb.duty.createdAt')" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column :label="t('common.actions')" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">{{ t('common.edit') }}</el-button>
            <el-button link type="danger" @click="handleDelete(row)">{{ t('common.delete') }}</el-button>
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
      width="600px"
      class="custom-dialog"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
      >
        <el-form-item label="值班日期" prop="dutyDate">
          <el-date-picker
            v-model="formData.dutyDate"
            type="date"
            placeholder="选择值班日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
            style="width: 100%"
            :disabled="isEdit"
          />
        </el-form-item>
        <el-form-item label="WB值班人员" prop="wbStaffName">
          <el-input v-model="formData.wbStaffName" placeholder="请输入WB值班人员姓名" />
        </el-form-item>
        <el-form-item label="FB值班人员" prop="fbStaffName">
          <el-input v-model="formData.fbStaffName" placeholder="请输入FB值班人员姓名" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">{{ t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitting">
          {{ t('common.confirm') }}
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import {
  getDutySchedules,
  createDutySchedule,
  updateDutySchedule,
  deleteDutySchedule
} from '../../api/cmdb'
import type {
  DutySchedule,
  DutyScheduleCreateRequest,
  DutyScheduleUpdateRequest
} from '../../types/cmdb'
import { useCommon } from '../../composables/useCommon'
import MonthCalendar from '../../components/MonthCalendar.vue'

const { t } = useCommon()

const viewMode = ref<'list' | 'calendar'>('calendar')
const loading = ref(false)
const submitting = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const currentId = ref<number | null>(null)
const formRef = ref<FormInstance>()

const filters = reactive({
  startDate: '',
  endDate: '',
  staffName: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 1000,
  total: 0
})

const tableData = ref<DutySchedule[]>([])

const formData = reactive<DutyScheduleCreateRequest>({
  dutyDate: '',
  wbStaffName: '',
  fbStaffName: ''
})

const formRules: FormRules = {
  dutyDate: [{ required: true, message: '请选择值班日期', trigger: 'change' }],
  wbStaffName: [{ required: true, message: '请输入WB值班人员姓名', trigger: 'blur' }],
  fbStaffName: [{ required: true, message: '请输入FB值班人员姓名', trigger: 'blur' }]
}

const dialogTitle = computed(() => (isEdit.value ? t('cmdb.duty.editSchedule') : t('cmdb.duty.addSchedule')))

const calendarItems = computed(() => {
  return tableData.value.map(item => ({
    ...item,
    date: formatDate(item.dutyDate)
  }))
})

watch(viewMode, (newMode) => {
  if (newMode === 'calendar') {
    pagination.pageSize = 1000
    loadData()
  } else {
    pagination.pageSize = 20
    loadData()
  }
})

// 加载数据
const loadData = async () => {
  loading.value = true
  try {
    const res = await getDutySchedules({
      startDate: filters.startDate || undefined,
      endDate: filters.endDate || undefined,
      staffName: filters.staffName || undefined,
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
  filters.staffName = ''
  handleSearch()
}

// 新增
const handleCreate = () => {
  isEdit.value = false
  currentId.value = null
  formData.dutyDate = ''
  formData.wbStaffName = ''
  formData.fbStaffName = ''
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: DutySchedule) => {
  isEdit.value = true
  currentId.value = row.id
  formData.dutyDate = row.dutyDate
  formData.wbStaffName = row.wbStaffName
  formData.fbStaffName = row.fbStaffName
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: DutySchedule) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除 ${row.dutyDate} 的值班记录吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await deleteDutySchedule(row.id)
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
        const updateData: DutyScheduleUpdateRequest = {
          wbStaffName: formData.wbStaffName,
          fbStaffName: formData.fbStaffName
        }
        await updateDutySchedule(currentId.value, updateData)
        ElMessage.success('更新成功')
      } else {
        const createData: DutyScheduleCreateRequest = {
          dutyDate: formData.dutyDate,
          wbStaffName: formData.wbStaffName,
          fbStaffName: formData.fbStaffName
        }
        await createDutySchedule(createData)
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
.duty-schedule-view {
  padding: var(--spacing-md);
}

.view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: var(--radius-xl);
}

.header-actions {
  display: flex;
  gap: 12px;
  align-items: center;
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

.calendar-card {
  border-radius: var(--radius-xl);
}

.calendar-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}

.calendar-item {
  margin-bottom: 6px;
  padding: 8px;
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.08) 0%, rgba(64, 158, 255, 0.03) 100%);
  border-radius: 6px;
  border-left: 3px solid var(--el-color-primary);
  transition: all 0.2s;
}

.calendar-item:hover {
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.12) 0%, rgba(64, 158, 255, 0.05) 100%);
  transform: translateX(2px);
}

.duty-row {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 4px;
}

.duty-row:last-child {
  margin-bottom: 0;
}

.duty-label {
  font-size: 11px;
  font-weight: 600;
  color: var(--el-color-primary);
  min-width: 28px;
}

.duty-name {
  font-size: 12px;
  font-weight: 500;
  color: var(--text-primary);
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

.staff-tag {
  max-width: 100%;
  white-space: normal;
  word-break: break-word;
  line-height: 1.5;
  padding: 4px 12px;
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

