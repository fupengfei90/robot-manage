<template>
  <el-dialog
    v-model="visible"
    title="分配权限"
    width="600px"
    :before-close="handleClose"
    class="custom-dialog"
  >
    <div v-if="role" class="role-info">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="角色名称">{{ role.name }}</el-descriptions-item>
        <el-descriptions-item label="角色描述">{{ role.description }}</el-descriptions-item>
      </el-descriptions>
    </div>

    <div class="permission-selection">
      <h4>选择权限</h4>
      <el-input
        v-model="searchText"
        placeholder="搜索权限"
        clearable
        style="margin-bottom: 12px"
      />
      <el-tree
        ref="treeRef"
        :data="permissionTree"
        :props="treeProps"
        :filter-node-method="filterNode"
        show-checkbox
        node-key="key"
        :default-checked-keys="selectedPermissionKeys"
        class="permission-tree"
      >
        <template #default="{ node, data }">
          <div class="tree-node">
            <div class="node-label">{{ data.intent }}{{ data.intent2 ? ':' + data.intent2 : '' }}</div>
            <div class="node-description">{{ data.description }}</div>
          </div>
        </template>
      </el-tree>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import type { ElTree } from 'element-plus'
import { 
  assignRolePermissions, 
  getPermissionTree, 
  type Role, 
  type PermissionTreeNode 
} from '@/api/rbac'

interface Props {
  modelValue: boolean
  role?: Role | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

interface TreeNode extends PermissionTreeNode {
  key: string
  children?: TreeNode[]
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const visible = ref(false)
const loading = ref(false)
const searchText = ref('')
const treeRef = ref<InstanceType<typeof ElTree>>()
const permissionTree = ref<TreeNode[]>([])
const selectedPermissionKeys = ref<string[]>([])

const treeProps = {
  children: 'children',
  label: 'intent'
}

const filterNode = (value: string, data: TreeNode) => {
  if (!value) return true
  const searchValue = value.toLowerCase()
  return (
    data.intent?.toLowerCase().includes(searchValue) ||
    data.intent2?.toLowerCase().includes(searchValue) ||
    data.description?.toLowerCase().includes(searchValue)
  )
}

const buildTreeData = (nodes: PermissionTreeNode[]): TreeNode[] => {
  return nodes.map(node => ({
    ...node,
    key: node.id ? `permission_${node.id}` : `intent_${node.intent}`,
    children: node.children ? buildTreeData(node.children) : undefined
  }))
}

const loadPermissionTree = async () => {
  try {
    const response = await getPermissionTree()
    permissionTree.value = buildTreeData(response)
  } catch (error) {
    console.error('加载权限树失败:', error)
    ElMessage.error('加载权限树失败')
  }
}

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val && props.role) {
    selectedPermissionKeys.value = props.role.permissions?.map(permission => `permission_${permission.id}`) || []
    searchText.value = ''
  }
})

watch(searchText, (val) => {
  treeRef.value?.filter(val)
})

watch(visible, (val) => {
  emit('update:modelValue', val)
})

const handleClose = () => {
  visible.value = false
}

const handleSubmit = async () => {
  if (!props.role || !treeRef.value) return
  
  try {
    loading.value = true
    
    // 获取选中的权限节点
    const checkedNodes = treeRef.value.getCheckedNodes(false, true) as TreeNode[]
    const permissionIds = checkedNodes
      .filter(node => node.id && node.key.startsWith('permission_'))
      .map(node => node.id!)
    
    await assignRolePermissions(props.role.id, { permission_ids: permissionIds })
    ElMessage.success('权限分配成功')
    visible.value = false
    emit('success')
  } catch (error) {
    console.error('权限分配失败:', error)
    ElMessage.error('权限分配失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadPermissionTree()
})
</script>

<style scoped>
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

.role-info {
  margin-bottom: var(--spacing-lg);
}

.permission-selection h4 {
  margin: var(--spacing-lg) 0 var(--spacing-md) 0;
  color: var(--text-primary);
  font-weight: 600;
}

.permission-tree {
  max-height: 400px;
  overflow-y: auto;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
}

.tree-node {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: var(--spacing-xs) var(--spacing-md) var(--spacing-xs) 0;
  line-height: 1.5;
}

.node-label {
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
  word-break: break-all;
}

.node-description {
  color: var(--text-secondary);
  font-size: 0.875rem;
  word-break: break-all;
}

.permission-tree :deep(.el-tree-node__content) {
  height: auto;
  min-height: 40px;
  padding: 4px 0;
  align-items: flex-start;
}

.permission-tree :deep(.el-checkbox) {
  margin-top: 8px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>