<template>
  <div class="vb-cmdb-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">VB CMDB信息 <el-tag size="small" type="warning">TODO</el-tag></h1>
        <p class="view-subtitle">VB系统配置信息管理</p>
      </div>
      <el-button type="primary" @click="handleCreate" class="create-btn hover-lift">
        <span>➕</span> 新增配置
      </el-button>
    </div>

    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filterForm" class="filter-form">
        <el-form-item label="系统名称">
          <el-input v-model="filterForm.systemName" placeholder="搜索系统名称" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item label="环境">
          <el-input v-model="filterForm.environment" placeholder="搜索环境" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table :data="tableData" v-loading="loading" stripe class="custom-table" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="systemName" label="系统名称" min-width="150" />
        <el-table-column prop="environment" label="环境" width="120" />
        <el-table-column prop="ipAddress" label="IP地址" width="150" />
        <el-table-column prop="port" label="端口" width="100" />
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column prop="owner" label="负责人" width="120" />
        <el-table-column prop="remark" label="备注" min-width="200" show-overflow-tooltip />
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" class="custom-dialog">
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="120px">
        <el-form-item label="系统名称" prop="systemName">
          <el-input v-model="formData.systemName" placeholder="请输入系统名称" />
        </el-form-item>
        <el-form-item label="环境" prop="environment">
          <el-input v-model="formData.environment" placeholder="请输入环境" />
        </el-form-item>
        <el-form-item label="IP地址" prop="ipAddress">
          <el-input v-model="formData.ipAddress" placeholder="请输入IP地址" />
        </el-form-item>
        <el-form-item label="端口" prop="port">
          <el-input v-model="formData.port" placeholder="请输入端口" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-input v-model="formData.status" placeholder="请输入状态" />
        </el-form-item>
        <el-form-item label="负责人" prop="owner">
          <el-input v-model="formData.owner" placeholder="请输入负责人" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitting">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { getVBCMDBList, createVBCMDB, updateVBCMDB, deleteVBCMDB } from '../../api/cmdb'
import type { VBCMDBInfo, VBCMDBInfoRequest } from '../../types/cmdb'

const loading = ref(false)
const submitting = ref(false)
const dialogVisible = ref(false)
const formRef = ref<FormInstance>()
const currentEditId = ref<number | null>(null)

const filterForm = reactive({
  systemName: '',
  environment: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

const tableData = ref<VBCMDBInfo[]>([])

const formData = reactive<VBCMDBInfoRequest>({
  systemName: '',
  environment: '',
  ipAddress: '',
  port: '',
  status: '',
  owner: '',
  remark: ''
})

const formRules: FormRules = {
  systemName: [{ required: true, message: '请输入系统名称', trigger: 'blur' }]
}

const dialogTitle = computed(() => currentEditId.value ? '编辑VB CMDB信息' : '新增VB CMDB信息')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await getVBCMDBList({
      systemName: filterForm.systemName || undefined,
      environment: filterForm.environment || undefined,
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
  filterForm.environment = ''
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
  Object.assign(formData, {
    systemName: '',
    environment: '',
    ipAddress: '',
    port: '',
    status: '',
    owner: '',
    remark: ''
  })
  formRef.value?.clearValidate()
  dialogVisible.value = true
}

const handleEdit = (row: VBCMDBInfo) => {
  currentEditId.value = row.id
  Object.assign(formData, {
    systemName: row.systemName,
    environment: row.environment,
    ipAddress: row.ipAddress,
    port: row.port,
    status: row.status,
    owner: row.owner,
    remark: row.remark
  })
  dialogVisible.value = true
}

const handleDelete = async (row: VBCMDBInfo) => {
  try {
    await ElMessageBox.confirm(`确定要删除系统"${row.systemName}"的配置吗？`, '确认删除', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteVBCMDB(row.id)
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
        await updateVBCMDB(currentEditId.value, formData)
        ElMessage.success('更新成功')
      } else {
        await createVBCMDB(formData)
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
.vb-cmdb-view {
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

.view-title :deep(.el-tag) {
  background-color: #fef3c7 !important;
  color: #92400e !important;
  border-color: #fbbf24 !important;
  vertical-align: middle;
  margin-left: 8px;
  -webkit-text-fill-color: #92400e !important;
}
</style>
