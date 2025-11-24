<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑权限' : '新增权限'"
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
      <el-form-item label="一级意图" prop="intent">
        <el-select
          v-model="form.intent"
          placeholder="请选择或输入一级意图"
          filterable
          allow-create
          style="width: 100%"
        >
          <el-option
            v-for="intent in presetIntents"
            :key="intent.value"
            :label="intent.label"
            :value="intent.value"
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="二级意图" prop="intent2">
        <el-select
          v-model="form.intent2"
          placeholder="请选择或输入二级意图（可选）"
          filterable
          allow-create
          clearable
          style="width: 100%"
        >
          <el-option
            v-for="intent2 in presetIntent2s"
            :key="intent2.value"
            :label="intent2.label"
            :value="intent2.value"
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="权限描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="4"
          placeholder="请输入权限描述"
          maxlength="500"
          show-word-limit
        />
      </el-form-item>
      
      <el-alert
        title="权限说明"
        type="info"
        :closable="false"
        show-icon
      >
        <template #default>
          <ul style="margin: 0; padding-left: 20px;">
            <li>一级意图：表示系统模块，如 dashboard、cmdb、system 等</li>
            <li>二级意图：表示具体操作，如 read、write、delete、manage 等</li>
            <li>如果只设置一级意图，表示拥有该模块的所有权限</li>
            <li>如果设置二级意图，表示只拥有该模块的特定操作权限</li>
          </ul>
        </template>
      </el-alert>
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
import { createPermission, updatePermission, type Permission, type PermissionRequest } from '@/api/rbac'

interface Props {
  modelValue: boolean
  permission?: Permission | null
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

const form = reactive<PermissionRequest>({
  intent: '',
  intent2: '',
  description: ''
})

const rules: FormRules = {
  intent: [
    { required: true, message: '请输入一级意图', trigger: 'blur' },
    { max: 100, message: '一级意图不能超过100个字符', trigger: 'blur' }
  ],
  intent2: [
    { max: 100, message: '二级意图不能超过100个字符', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入权限描述', trigger: 'blur' },
    { max: 500, message: '权限描述不能超过500个字符', trigger: 'blur' }
  ]
}

// 预设的一级意图
const presetIntents = [
  { label: '仪表盘模块', value: 'dashboard' },
  { label: 'CMDB模块', value: 'cmdb' },
  { label: '智能交付模块', value: 'delivery' },
  { label: '数字员工模块', value: 'digital' },
  { label: '系统管理模块', value: 'system' }
]

// 预设的二级意图
const presetIntent2s = [
  { label: '查看权限', value: 'read' },
  { label: '编辑权限', value: 'write' },
  { label: '删除权限', value: 'delete' },
  { label: '管理权限', value: 'manage' },
  { label: '执行权限', value: 'execute' },
  { label: '导出权限', value: 'export' },
  { label: '导入权限', value: 'import' }
]

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    resetForm()
    if (props.permission) {
      isEdit.value = true
      Object.assign(form, {
        intent: props.permission.intent,
        intent2: props.permission.intent2 || '',
        description: props.permission.description
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
    intent: '',
    intent2: '',
    description: ''
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
    
    if (isEdit.value && props.permission) {
      await updatePermission(props.permission.id, form)
      ElMessage.success('权限更新成功')
    } else {
      await createPermission(form)
      ElMessage.success('权限创建成功')
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