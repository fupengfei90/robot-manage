<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑任务' : '新增任务'"
    width="800px"
    :before-close="handleClose"
    class="task-form-dialog"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
      size="default"
    >
      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item label="任务名称" prop="name">
            <el-input
              v-model="form.name"
              placeholder="请输入任务名称"
              maxlength="100"
              show-word-limit
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="任务状态" prop="active">
            <el-switch
              v-model="form.active"
              :active-value="1"
              :inactive-value="0"
              active-text="启用"
              inactive-text="停用"
            />
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item label="任务分类" prop="category">
            <el-select
              v-model="form.category"
              placeholder="请选择任务分类"
              filterable
              allow-create
              style="width: 100%"
            >
              <el-option
                v-for="category in categories"
                :key="category"
                :label="getCategoryLabel(category)"
                :value="category"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="数据中心" prop="dcn">
            <el-select
              v-model="form.dcn"
              placeholder="请选择数据中心"
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

      <el-form-item label="Cron表达式" prop="cron">
        <CronEditor v-model="form.cron" />
      </el-form-item>

      <el-form-item label="工作流ID" prop="workflow">
        <el-input
          v-model="form.workflow"
          placeholder="请输入工作流ID，如：bd35a471-a068-4cb1-8215-2c0003ba0382"
          maxlength="100"
        />
      </el-form-item>

      <el-form-item label="执行Token" prop="exec_token">
        <el-input
          v-model="form.exec_token"
          placeholder="请输入执行token，如：app-WGfAztV6fmzz60cZZ6p9aG74"
          maxlength="100"
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
import { ref, reactive, watch, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import CronEditor from './CronEditor.vue'
import { 
  createScheduleTask, 
  updateScheduleTask, 
  getScheduleTaskCategories, 
  getScheduleTaskDCNs,
  type ScheduleTask,
  type ScheduleTaskRequest 
} from '@/api/schedule-task'

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
const categories = ref<string[]>([])
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
    { required: true, message: '请输入任务名称', trigger: 'blur' },
    { max: 100, message: '任务名称不能超过100个字符', trigger: 'blur' }
  ],
  cron: [
    { required: true, message: '请输入cron表达式', trigger: 'blur' }
  ],
  workflow: [
    { required: true, message: '请输入工作流ID', trigger: 'blur' },
    { max: 100, message: '工作流ID不能超过100个字符', trigger: 'blur' }
  ],
  exec_token: [
    { required: true, message: '请输入执行token', trigger: 'blur' },
    { max: 100, message: '执行token不能超过100个字符', trigger: 'blur' }
  ],
  category: [
    { required: true, message: '请选择任务分类', trigger: 'change' },
    { max: 32, message: '任务分类不能超过32个字符', trigger: 'blur' }
  ],
  dcn: [
    { required: true, message: '请选择数据中心', trigger: 'change' },
    { max: 100, message: '数据中心不能超过100个字符', trigger: 'blur' }
  ]
}

// 分类标签映射
const categoryLabels: Record<string, string> = {
  'rota': '值班提醒',
  'weekly-report': '周报生成',
  'version-align': '版本通知'
}

const getCategoryLabel = (category: string) => {
  return categoryLabels[category] || category
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
      ElMessage.success('任务更新成功')
    } else {
      await createScheduleTask(form)
      ElMessage.success('任务创建成功')
    }
    
    visible.value = false
    emit('success')
  } catch (error) {
    console.error('提交失败:', error)
  } finally {
    loading.value = false
  }
}

// 加载分类和数据中心
const loadOptions = async () => {
  try {
    const [categoriesRes, dcnsRes] = await Promise.all([
      getScheduleTaskCategories(),
      getScheduleTaskDCNs()
    ])
    categories.value = categoriesRes.data
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