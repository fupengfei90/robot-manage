<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑角色' : '新增角色'"
    width="600px"
    :before-close="handleClose"
    class="custom-dialog"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
      size="default"
    >
      <el-form-item label="角色名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入角色名称"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="角色描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="4"
          placeholder="请输入角色描述"
          maxlength="500"
          show-word-limit
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          {{ isEdit ? '更新' : '创建' }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { createRole, updateRole, type Role, type RoleRequest } from '@/api/rbac'

interface Props {
  modelValue: boolean
  role?: Role | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const visible = ref(false)
const loading = ref(false)
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive<RoleRequest>({
  name: '',
  description: '',
  permission_ids: []
})

const rules: FormRules = {
  name: [
    { required: true, message: '请输入角色名称', trigger: 'blur' },
    { max: 100, message: '角色名称不能超过100个字符', trigger: 'blur' }
  ],
  description: [
    { max: 500, message: '角色描述不能超过500个字符', trigger: 'blur' }
  ]
}

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    resetForm()
    if (props.role) {
      isEdit.value = true
      Object.assign(form, {
        name: props.role.name,
        description: props.role.description,
        permission_ids: props.role.permissions?.map(permission => permission.id) || []
      })
    } else {
      isEdit.value = false
    }
  }
})

watch(visible, (val) => {
  emit('update:modelValue', val)
})

const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  Object.assign(form, {
    name: '',
    description: '',
    permission_ids: []
  })
}

const handleClose = () => {
  visible.value = false
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    loading.value = true
    
    if (isEdit.value && props.role) {
      await updateRole(props.role.id, form)
      ElMessage.success('角色更新成功')
    } else {
      await createRole(form)
      ElMessage.success('角色创建成功')
    }
    
    visible.value = false
    emit('success')
  } catch (error) {
    console.error('提交失败:', error)
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

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>