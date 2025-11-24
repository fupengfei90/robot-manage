<template>
  <div class="role-management">
    <!-- 搜索筛选 -->
    <el-card shadow="never" class="filter-card glass-effect animate-fade-in">
      <el-form :inline="true" :model="filters" class="filter-form">
        <el-form-item label="角色名称">
          <el-input
            v-model="filters.name"
            placeholder="请输入角色名称"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          />
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
          <span>➕</span> 新增角色
        </el-button>
      </div>
    </div>

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
        <el-table-column prop="name" label="角色名称" width="150" show-overflow-tooltip />
        <el-table-column prop="description" label="角色描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="user_count" label="用户数量" width="100">
          <template #default="{ row }">
            <el-tag type="info" size="small">{{ row.user_count }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="permission_count" label="权限数量" width="100">
          <template #default="{ row }">
            <el-tag type="success" size="small">{{ row.permission_count }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="permission_names" label="权限列表" min-width="250">
          <template #default="{ row }">
            <el-tag
              v-for="permissionName in row.permission_names.slice(0, 3)"
              :key="permissionName"
              type="warning"
              size="small"
              class="permission-tag"
            >
              {{ permissionName }}
            </el-tag>
            <el-tag
              v-if="row.permission_names.length > 3"
              type="info"
              size="small"
              class="permission-tag"
            >
              +{{ row.permission_names.length - 3 }}
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
            <el-button link type="primary" @click="handleAssignPermissions(row)">分配权限</el-button>
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

    <!-- 角色表单对话框 -->
    <RoleForm
      v-model="formVisible"
      :role="currentRole"
      @success="handleFormSuccess"
    />

    <!-- 权限分配对话框 -->
    <RolePermissionAssign
      v-model="permissionAssignVisible"
      :role="currentRole"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import RoleForm from './RoleForm.vue'
import RolePermissionAssign from './RolePermissionAssign.vue'
import {
  getRoles,
  deleteRole,
  type Role,
  type RoleQuery
} from '@/api/rbac'

const loading = ref(false)
const tableData = ref<Role[]>([])
const formVisible = ref(false)
const permissionAssignVisible = ref(false)
const currentRole = ref<Role | null>(null)

const filters = reactive<RoleQuery>({
  name: ''
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
    
    const response = await getRoles(params)
    tableData.value = response.list
    pagination.total = response.total
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  loadData()
}

const handleReset = () => {
  Object.assign(filters, {
    name: ''
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

const handleCreate = () => {
  currentRole.value = null
  formVisible.value = true
}

const handleEdit = (role: Role) => {
  currentRole.value = role
  formVisible.value = true
}

const handleAssignPermissions = (role: Role) => {
  currentRole.value = role
  permissionAssignVisible.value = true
}

const handleDelete = async (role: Role) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除角色"${role.name}"吗？删除后无法恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await deleteRole(role.id)
    ElMessage.success('删除成功')
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

const handleFormSuccess = () => {
  loadData()
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.role-management {
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

.permission-tag {
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