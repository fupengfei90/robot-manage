<template>
  <div class="permission-management">
    <!-- 操作栏 -->
    <div class="action-bar">
      <div class="action-left">
        <el-button type="primary" @click="handleCreate">
          <span>➕</span> 新增权限
        </el-button>
        <el-button @click="loadData">
          <span>🔄</span> 刷新
        </el-button>
      </div>
    </div>

    <!-- 权限树形表格 -->
    <el-card shadow="never" class="table-card glass-effect animate-fade-in-up">
      <el-table
        v-loading="loading"
        :data="treeData"
        row-key="key"
        :tree-props="{ children: 'children' }"
        stripe
        class="custom-table"
        style="width: 100%"
      >
        <el-table-column prop="intent" label="一级意图" width="200" />
        <el-table-column prop="intent2" label="二级意图" width="200">
          <template #default="{ row }">
            <span v-if="row.intent2">{{ row.intent2 }}</span>
            <el-tag v-else type="info" size="small">全部权限</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="权限描述" min-width="300" show-overflow-tooltip />
        <el-table-column prop="created_at" label="创建时间" width="170">
          <template #default="{ row }">
            <span v-if="row.created_at">{{ formatDateTime(row.created_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <template v-if="row.id">
              <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
              <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
            </template>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 权限表单对话框 -->
    <PermissionForm
      v-model="formVisible"
      :permission="currentPermission"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import PermissionForm from './PermissionForm.vue'
import {
  getPermissions,
  deletePermission,
  type Permission
} from '@/api/rbac'

interface TreePermission extends Permission {
  key: string
  children?: TreePermission[]
}

const loading = ref(false)
const treeData = ref<TreePermission[]>([])
const formVisible = ref(false)
const currentPermission = ref<Permission | null>(null)

const formatDateTime = (dateTime: string) => {
  if (!dateTime) return ''
  return new Date(dateTime).toLocaleString('zh-CN')
}

const buildTreeData = (permissions: Permission[]): TreePermission[] => {
  const intentMap = new Map<string, TreePermission>()
  
  // 按一级意图分组
  permissions.forEach(permission => {
    const key = `${permission.intent}_${permission.id}`
    const treePermission: TreePermission = {
      ...permission,
      key,
      children: []
    }
    
    if (!permission.intent2) {
      // 一级意图权限
      intentMap.set(permission.intent, treePermission)
    } else {
      // 二级意图权限
      let parentIntent = intentMap.get(permission.intent)
      if (!parentIntent) {
        // 创建虚拟的一级意图节点
        parentIntent = {
          id: 0,
          intent: permission.intent,
          intent2: '',
          description: `${permission.intent}模块`,
          created_at: '',
          updated_at: '',
          key: `intent_${permission.intent}`,
          children: []
        }
        intentMap.set(permission.intent, parentIntent)
      }
      parentIntent.children!.push(treePermission)
    }
  })
  
  return Array.from(intentMap.values())
}

const loadData = async () => {
  try {
    loading.value = true
    const response = await getPermissions({ page: 1, page_size: 1000 })
    treeData.value = buildTreeData(response.list)
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败')
  } finally {
    loading.value = false
  }
}

const handleCreate = () => {
  currentPermission.value = null
  formVisible.value = true
}

const handleEdit = (permission: Permission) => {
  currentPermission.value = permission
  formVisible.value = true
}

const handleDelete = async (permission: Permission) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除权限"${permission.intent}${permission.intent2 ? ':' + permission.intent2 : ''}"吗？删除后无法恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await deletePermission(permission.id)
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
.permission-management {
  padding: 0;
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

.table-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}
</style>