<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? t('system.task.editTask') : t('system.task.addTask')"
    width="800px"
    :before-close="handleClose"
    class="task-form-dialog"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="120px"
      size="default"
    >
      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item :label="t('system.task.taskName')" prop="name">
            <el-input
              v-model="form.name"
              :placeholder="t('system.task.searchPlaceholder')"
              maxlength="100"
              show-word-limit
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="t('system.task.status')" prop="active">
            <el-switch
              v-model="form.active"
              :active-value="1"
              :inactive-value="0"
              :active-text="t('common.enabled')"
              :inactive-text="t('common.disabled')"
            />
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item :label="t('system.task.category')" prop="category">
            <el-input
              v-model="form.category"
              :placeholder="t('system.task.selectCategory')"
              maxlength="32"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="t('system.task.dcn')" prop="dcn">
            <el-select
              v-model="form.dcn"
              :placeholder="t('system.task.selectDCN')"
              filterable
              allow-create
              style="width: 100%"
            >
              <el-option
                v-for="dcn in dcns"
                :key="dcn"
                :label="dcn"
                :value="dcn"
              />
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item :label="t('system.task.cronExpression')" prop="cron">
        <CronEditor v-model="form.cron" />
      </el-form-item>

      <el-form-item label="Workflow ID" prop="workflow">
        <el-input
            v-model="form.workflow"
            placeholder="e.g., bd35a471-a068-4cb1-8215-2c0003ba0382"
            maxlength="100"
        />
      </el-form-item>

      <el-form-item label="Exec Token" prop="exec_token">
        <el-input
            v-model="form.exec_token"
            placeholder="e.g., app-WGfAztV6fmzz60cZZ6p9aG74"
            maxlength="100"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">{{ t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          {{ isEdit ? t('common.save') : t('common.add') }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import CronEditor from './CronEditor.vue'
import { 
  createScheduleTask, 
  updateScheduleTask, 
  getScheduleTaskDCNs,
  type ScheduleTask,
  type ScheduleTaskRequest 
} from '@/api/schedule-task'
import { useCommon } from '@/composables/useCommon'

const { t } = useCommon()

interface Props {
  modelValue: boolean
  task?: ScheduleTask | null
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
const dcns = ref<string[]>([])

const isEdit = ref(false)

const form = reactive<ScheduleTaskRequest>({
  name: '',
  active: 0,
  cron: '0 0 10 * * *',
  workflow: '',
  exec_token: '',
  category: '',
  dcn: ''
})

const rules: FormRules = {
  name: [
    { required: true, message: 'Please enter task name', trigger: 'blur' },
    { max: 100, message: 'Task name cannot exceed 100 characters', trigger: 'blur' }
  ],
  cron: [
    { required: true, message: 'Please enter cron expression', trigger: 'blur' }
  ],
  workflow: [
    { required: true, message: 'Please enter workflow ID', trigger: 'blur' },
    { max: 100, message: 'Workflow ID cannot exceed 100 characters', trigger: 'blur' }
  ],
  exec_token: [
    { required: true, message: 'Please enter exec token', trigger: 'blur' },
    { max: 100, message: 'Exec token cannot exceed 100 characters', trigger: 'blur' }
  ],
  category: [
    { required: true, message: 'Please select category', trigger: 'change' },
    { max: 32, message: 'Category cannot exceed 32 characters', trigger: 'blur' }
  ],
  dcn: [
    { required: true, message: 'Please select DCN', trigger: 'change' },
    { max: 100, message: 'DCN cannot exceed 100 characters', trigger: 'blur' }
  ]
}

// 监听显示状态
watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    resetForm()
    if (props.task) {
      isEdit.value = true
      Object.assign(form, {
        name: props.task.name,
        active: props.task.active,
        cron: props.task.cron,
        workflow: props.task.workflow,
        exec_token: props.task.exec_token,
        category: props.task.category,
        dcn: props.task.dcn
      })
    } else {
      isEdit.value = false
    }
  }
})

watch(visible, (val) => {
  emit('update:modelValue', val)
})

// 重置表单
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  Object.assign(form, {
    name: '',
    active: 0,
    cron: '0 0 10 * * *',
    workflow: '',
    exec_token: '',
    category: '',
    dcn: ''
  })
}

// 关闭对话框
const handleClose = () => {
  visible.value = false
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    loading.value = true
    
    if (isEdit.value && props.task) {
      await updateScheduleTask(props.task.id, form)
      ElMessage.success(t('common.saveSuccess'))
    } else {
      await createScheduleTask(form)
      ElMessage.success(t('common.operationSuccess'))
    }
    
    visible.value = false
    emit('success')
  } catch (error) {
    console.error('提交失败:', error)
  } finally {
    loading.value = false
  }
}

// 加载数据中心
const loadOptions = async () => {
  try {
    const dcnsRes = await getScheduleTaskDCNs()
    dcns.value = dcnsRes.data
  } catch (error) {
    console.error('加载选项失败:', error)
  }
}

onMounted(() => {
  loadOptions()
})
</script>

<style scoped>
.task-form-dialog :deep(.el-dialog) {
  background: rgba(30, 30, 30, 0.95);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.task-form-dialog :deep(.el-dialog__header) {
  background: rgba(255, 255, 255, 0.05);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.task-form-dialog :deep(.el-form-item__label) {
  color: rgba(255, 255, 255, 0.9);
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>