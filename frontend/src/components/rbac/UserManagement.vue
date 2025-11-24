<template>
  <div class="user-management">
    <!-- 搜索筛选 -->
    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filters" class="filter-form">
        <el-form-item label="用户ID">
          <el-input
            v-model="filters.user_id"
            placeholder="请输入用户ID"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          />
        </el-form-item>
        <el-form-item label="用户姓名">
          <el-input
            v-model="filters.name"
            placeholder="请输入用户姓名"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="filters.active" placeholder="请选择状态" clearable style="width: 120px">
            <el-option label="启用" :value="1" />
            <el-option label="停用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="filters.role_id" placeholder="请选择角色" clearable style="width: 150px">
            <el-option
              v-for="role in allRoles"
              :key="role.id"
              :label="role.name"
              :value="role.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 操作栏 -->
    <div class="action-bar">
      <div class="action-left">
        <el-button type="primary" @click="handleCreate">
          <span>➕</span> 新增用户
        </el-button>
        <el-dropdown @command="handleBatchOperation" :disabled="!selectedUsers.length">
          <el-button :disabled="!selectedUsers.length">
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
      <div class="action-right">
        <el-text type="info" size="small">
          已选择 {{ selectedUsers.length }} 项
        </el-text>
      </div>
    </div>

    <!-- 数据表格 -->
    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        v-loading="loading"
        :data="tableData"
        @selection-change="handleSelectionChange"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="user_id" label="用户ID" width="150" show-overflow-tooltip />
        <el-table-column prop="name" label="用户姓名" width="120" show-overflow-tooltip />
        <el-table-column prop="p_id" label="P账号" width="120" show-overflow-tooltip />
        <el-table-column prop="phone" label="手机号" width="130" />
        <el-table-column prop="active" label="状态" width="100">
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
        <el-table-column prop="role_names" label="角色" min-width="200">
          <template #default="{ row }">
            <el-tag
              v-for="roleName in row.role_names"
              :key="roleName"
              type="info"
              size="small"
              class="role-tag"
            >
              {{ roleName }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="170">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
            <el-button link type="primary" @click="handleAssignRoles(row)">分配角色</el-button>
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
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 用户表单对话框 -->
    <UserForm
      v-model="formVisible"
      :user="currentUser"
      :roles="allRoles"
      @success="handleFormSuccess"
    />

    <!-- 角色分配对话框 -->
    <UserRoleAssign
      v-model="roleAssignVisible"
      :user="currentUser"
      :roles="allRoles"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ArrowDown } from '@element-plus/icons-vue'
import UserForm from './UserForm.vue'
import UserRoleAssign from './UserRoleAssign.vue'
import {
  getUsers,
  deleteUser,
  updateUserStatus,
  batchOperateUsers,
  getRoles,
  type User,
  type UserQuery,
  type Role,
  type BatchUserOperationRequest
} from '@/api/rbac'

const loading = ref(false)
const tableData = ref<User[]>([])
const selectedUsers = ref<User[]>([])
const formVisible = ref(false)
const roleAssignVisible = ref(false)
const currentUser = ref<User | null>(null)
const allRoles = ref<Role[]>([])

const statusLoading = ref<Record<number, boolean>>({})

const filters = reactive<UserQuery>({
  user_id: '',
  name: '',
  active: undefined,
  role_id: undefined
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const formatDateTime = (dateTime: string) => {
  if (!dateTime) return ''
  return new Date(dateTime).toLocaleString('zh-CN')
}

const loadData = async () => {
  try {
    loading.value = true
    const params = {
      ...filters,
      page: pagination.page,
      page_size: pagination.pageSize
    }
    
    const response = await getUsers(params)
    tableData.value = response.list
    pagination.total = response.total
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败')
  } finally {
    loading.value = false
  }
}

const loadRoles = async () => {
  try {
    const response = await getRoles({ page: 1, page_size: 1000 })
    allRoles.value = response.list
  } catch (error) {
    console.error('加载角色失败:', error)
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  Object.assign(filters, {
    user_id: '',
    name: '',
    active: undefined,
    role_id: undefined
  })
  handleSearch()
}

const handlePageChange = (page: number) => {
  pagination.page = page
  loadData()
}

const handlePageSizeChange = (pageSize: number) => {
  pagination.pageSize = pageSize
  pagination.page = 1
  loadData()
}

const handleSelectionChange = (selection: User[]) => {
  selectedUsers.value = selection
}

const handleCreate = () => {
  currentUser.value = null
  formVisible.value = true
}

const handleEdit = (user: User) => {
  currentUser.value = user
  formVisible.value = true
}

const handleAssignRoles = (user: User) => {
  currentUser.value = user
  roleAssignVisible.value = true
}

const handleDelete = async (user: User) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除用户"${user.name}"吗？删除后无法恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await deleteUser(user.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

const handleStatusChange = async (user: User, active: number) => {
  try {
    statusLoading.value[user.id] = true
    await updateUserStatus(user.id, active)
    user.active = active
    ElMessage.success(`用户已${active === 1 ? '启用' : '停用'}`)
  } catch (error) {
    console.error('状态更新失败:', error)
    ElMessage.error('状态更新失败')
  } finally {
    statusLoading.value[user.id] = false
  }
}

const handleBatchOperation = async (command: string) => {
  if (!selectedUsers.value.length) return
  
  const operationMap: Record<string, string> = {
    enable: '启用',
    disable: '停用',
    delete: '删除'
  }
  
  const operation = operationMap[command]
  if (!operation) return
  
  try {
    await ElMessageBox.confirm(
      `确定要${operation}选中的 ${selectedUsers.value.length} 个用户吗？`,
      `确认${operation}`,
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: command === 'delete' ? 'warning' : 'info'
      }
    )
    
    const userIds = selectedUsers.value.map(user => user.id)
    const request: BatchUserOperationRequest = {
      user_ids: userIds,
      operation: command as 'enable' | 'disable' | 'delete'
    }
    
    await batchOperateUsers(request)
    ElMessage.success(`批量${operation}成功`)
    selectedUsers.value = []
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(`批量${operation}失败:`, error)
      ElMessage.error(`批量${operation}失败`)
    }
  }
}

const handleFormSuccess = () => {
  loadData()
}

onMounted(() => {
  loadData()
  loadRoles()
})
</script>

<style scoped>
.user-management {
  padding: 0;
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

.action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
  padding: var(--spacing-md) var(--spacing-lg);
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-xl);
  backdrop-filter: blur(10px);
}

.action-left {
  display: flex;
  gap: var(--spacing-md);
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

.role-tag {
  margin-right: var(--spacing-xs);
  margin-bottom: var(--spacing-xs);
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
</style>