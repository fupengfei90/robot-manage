<template>
  <el-dialog
    v-model="visible"
    title="分配角色"
    width="500px"
    :before-close="handleClose"
    class="custom-dialog"
  >
    <div v-if="user" class="user-info">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="用户ID">{{ user.user_id }}</el-descriptions-item>
        <el-descriptions-item label="用户姓名">{{ user.name }}</el-descriptions-item>
      </el-descriptions>
    </div>

    <div class="role-selection">
      <h4>选择角色</h4>
      <el-checkbox-group v-model="selectedRoleIds" class="role-list">
        <div v-for="role in roles" :key="role.id" class="role-item">
          <el-checkbox :value="role.id" class="role-checkbox">
            <div class="role-content">
              <div class="role-name">{{ role.name }}</div>
              <div class="role-description">{{ role.description }}</div>
              <div class="role-stats">
                用户数: {{ role.user_count }} | 权限数: {{ role.permission_count }}
              </div>
            </div>
          </el-checkbox>
        </div>
      </el-checkbox-group>
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
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { assignUserRoles, type User, type Role } from '@/api/rbac'

interface Props {
  modelValue: boolean
  user?: User | null
  roles: Role[]
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const visible = ref(false)
const loading = ref(false)
const selectedRoleIds = ref<number[]>([])

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val && props.user) {
    selectedRoleIds.value = props.user.roles?.map(role => role.id) || []
  }
})

watch(visible, (val) => {
  emit('update:modelValue', val)
})

const handleClose = () => {
  visible.value = false
}

const handleSubmit = async () => {
  if (!props.user) return
  
  try {
    loading.value = true
    await assignUserRoles(props.user.id, { role_ids: selectedRoleIds.value })
    ElMessage.success('角色分配成功')
    visible.value = false
    emit('success')
  } catch (error) {
    console.error('角色分配失败:', error)
    ElMessage.error('角色分配失败')
  } finally {
    loading.value = false
  }
}
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

.user-info {
  margin-bottom: var(--spacing-lg);
}

.role-selection h4 {
  margin: var(--spacing-lg) 0 var(--spacing-md) 0;
  color: var(--text-primary);
  font-weight: 600;
}

.role-list {
  max-height: 300px;
  overflow-y: auto;
}

.role-item {
  margin-bottom: var(--spacing-md);
  padding: var(--spacing-md);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  background: rgba(255, 255, 255, 0.02);
}

.role-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.role-checkbox {
  width: 100%;
}

.role-content {
  margin-left: var(--spacing-md);
}

.role-name {
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.role-description {
  color: var(--text-secondary);
  font-size: 0.875rem;
  margin-bottom: var(--spacing-xs);
}

.role-stats {
  color: var(--text-muted);
  font-size: 0.75rem;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>