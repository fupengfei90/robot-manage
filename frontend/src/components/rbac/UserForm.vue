<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑用户' : '新增用户'"
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
      <el-form-item label="用户ID" prop="user_id">
        <el-input
          v-model="form.user_id"
          placeholder="请输入用户ID"
          maxlength="100"
          show-word-limit
          :disabled="isEdit"
        />
      </el-form-item>
      
      <el-form-item label="用户姓名" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入用户姓名"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="P账号" prop="p_id">
        <el-input
          v-model="form.p_id"
          placeholder="请输入P账号"
          maxlength="100"
        />
      </el-form-item>
      
      <el-form-item label="手机号" prop="phone">
        <el-input
          v-model="form.phone"
          placeholder="请输入手机号"
          maxlength="20"
        />
      </el-form-item>
      
      <el-form-item label="状态" prop="active">
        <el-switch
          v-model="form.active"
          :active-value="1"
          :inactive-value="0"
          active-text="启用"
          inactive-text="停用"
        />
      </el-form-item>
      
      <el-form-item label="分配角色" prop="role_ids">
        <el-select
          v-model="form.role_ids"
          placeholder="请选择角色"
          multiple
          collapse-tags
          collapse-tags-tooltip
          style="width: 100%"
        >
          <el-option
            v-for="role in roles"
            :key="role.id"
            :label="role.name"
            :value="role.id"
          >
            <span>{{ role.name }}</span>
            <span style="float: right; color: var(--el-text-color-secondary); font-size: 13px">
              {{ role.description }}
            </span>
          </el-option>
        </el-select>
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
import { createUser, updateUser, type User, type UserRequest, type Role } from '@/api/rbac'

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
const formRef = ref<FormInstance>()
const isEdit = ref(false)

const form = reactive<UserRequest>({
  user_id: '',
  name: '',
  active: 1,
  p_id: '',
  phone: '',
  role_ids: []
})

const rules: FormRules = {
  user_id: [
    { required: true, message: '请输入用户ID', trigger: 'blur' },
    { max: 100, message: '用户ID不能超过100个字符', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入用户姓名', trigger: 'blur' },
    { max: 100, message: '用户姓名不能超过100个字符', trigger: 'blur' }
  ],
  p_id: [
    { max: 100, message: 'P账号不能超过100个字符', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }
  ]
}

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    resetForm()
    if (props.user) {
      isEdit.value = true
      Object.assign(form, {
        user_id: props.user.user_id,
        name: props.user.name,
        active: props.user.active,
        p_id: props.user.p_id || '',
        phone: props.user.phone || '',
        role_ids: props.user.roles?.map(role => role.id) || []
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
    user_id: '',
    name: '',
    active: 1,
    p_id: '',
    phone: '',
    role_ids: []
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
    
    if (isEdit.value && props.user) {
      await updateUser(props.user.id, form)
      ElMessage.success('用户更新成功')
    } else {
      await createUser(form)
      ElMessage.success('用户创建成功')
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