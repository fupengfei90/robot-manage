<template>
  <div class="cron-editor">
    <el-input
      v-model="cronValue"
      placeholder="请输入cron表达式，如：0 0 10 * * *"
      @input="handleCronChange"
      class="cron-input"
    >
      <template #append>
        <el-button @click="showBuilder = !showBuilder" type="primary">
          {{ showBuilder ? '隐藏' : '可视化' }}
        </el-button>
      </template>
    </el-input>
    
    <div v-if="cronDescription" class="cron-description">
      <el-text type="info" size="small">
        <el-icon><InfoFilled /></el-icon>
        {{ cronDescription }}
      </el-text>
    </div>
    
    <div v-if="cronError" class="cron-error">
      <el-text type="danger" size="small">
        <el-icon><WarningFilled /></el-icon>
        {{ cronError }}
      </el-text>
    </div>

    <el-collapse-transition>
      <div v-show="showBuilder" class="cron-builder">
        <el-card shadow="never" class="builder-card">
          <template #header>
            <span>Cron表达式生成器</span>
          </template>
          
          <el-form :model="cronForm" label-width="80px" size="small">
            <el-row :gutter="16">
              <el-col :span="8">
                <el-form-item label="秒">
                  <el-select v-model="cronForm.second" @change="buildCron">
                    <el-option label="每秒" value="*" />
                    <el-option label="0秒" value="0" />
                    <el-option v-for="i in 59" :key="i" :label="`${i}秒`" :value="i.toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="分钟">
                  <el-select v-model="cronForm.minute" @change="buildCron">
                    <el-option label="每分钟" value="*" />
                    <el-option label="0分" value="0" />
                    <el-option v-for="i in 59" :key="i" :label="`${i}分`" :value="i.toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="小时">
                  <el-select v-model="cronForm.hour" @change="buildCron">
                    <el-option label="每小时" value="*" />
                    <el-option v-for="i in 23" :key="i" :label="`${i}点`" :value="i.toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-row :gutter="16">
              <el-col :span="8">
                <el-form-item label="日期">
                  <el-select v-model="cronForm.day" @change="buildCron">
                    <el-option label="每天" value="*" />
                    <el-option v-for="i in 31" :key="i" :label="`${i}日`" :value="i.toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="月份">
                  <el-select v-model="cronForm.month" @change="buildCron">
                    <el-option label="每月" value="*" />
                    <el-option v-for="(month, i) in months" :key="i" :label="month" :value="(i + 1).toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="星期">
                  <el-select v-model="cronForm.weekday" @change="buildCron" multiple collapse-tags>
                    <el-option label="每天" value="*" />
                    <el-option v-for="(day, i) in weekdays" :key="i" :label="day" :value="i.toString()" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-form-item>
              <el-space>
                <el-button @click="setPreset('daily')" size="small">每天执行</el-button>
                <el-button @click="setPreset('weekly')" size="small">每周执行</el-button>
                <el-button @click="setPreset('workday')" size="small">工作日执行</el-button>
                <el-button @click="setPreset('monthly')" size="small">每月执行</el-button>
              </el-space>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
    </el-collapse-transition>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import { InfoFilled, WarningFilled } from '@element-plus/icons-vue'

interface Props {
  modelValue: string
}

interface Emits {
  (e: 'update:modelValue', value: string): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const cronValue = ref(props.modelValue)
const showBuilder = ref(false)
const cronError = ref('')

const cronForm = ref({
  second: '0',
  minute: '0',
  hour: '10',
  day: '*',
  month: '*',
  weekday: '*'
})

const months = [
  '1月', '2月', '3月', '4月', '5月', '6月',
  '7月', '8月', '9月', '10月', '11月', '12月'
]

const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']

// 计算cron描述
const cronDescription = computed(() => {
  if (!cronValue.value || cronError.value) return ''
  return parseCronDescription(cronValue.value)
})

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  cronValue.value = newVal
  parseCronToForm(newVal)
})

// 处理cron表达式变化
const handleCronChange = (value: string) => {
  cronValue.value = value
  validateCron(value)
  emit('update:modelValue', value)
}

// 验证cron表达式
const validateCron = (cron: string) => {
  cronError.value = ''
  if (!cron) return
  
  const parts = cron.trim().split(/\s+/)
  if (parts.length !== 6) {
    cronError.value = 'cron表达式必须包含6个字段（秒 分 时 日 月 周）'
    return
  }
  
  // 简单验证每个字段
  const [second, minute, hour, day, month, weekday] = parts
  
  if (!isValidCronField(second, 0, 59)) {
    cronError.value = '秒字段无效（0-59）'
    return
  }
  if (!isValidCronField(minute, 0, 59)) {
    cronError.value = '分钟字段无效（0-59）'
    return
  }
  if (!isValidCronField(hour, 0, 23)) {
    cronError.value = '小时字段无效（0-23）'
    return
  }
  if (!isValidCronField(day, 1, 31)) {
    cronError.value = '日期字段无效（1-31）'
    return
  }
  if (!isValidCronField(month, 1, 12)) {
    cronError.value = '月份字段无效（1-12）'
    return
  }
  if (!isValidCronField(weekday, 0, 6)) {
    cronError.value = '星期字段无效（0-6）'
    return
  }
}

// 验证cron字段
const isValidCronField = (field: string, min: number, max: number): boolean => {
  if (field === '*') return true
  
  // 处理逗号分隔的值
  if (field.includes(',')) {
    const values = field.split(',')
    return values.every(v => {
      const num = parseInt(v.trim())
      return !isNaN(num) && num >= min && num <= max
    })
  }
  
  // 处理范围
  if (field.includes('-')) {
    const [start, end] = field.split('-').map(v => parseInt(v.trim()))
    return !isNaN(start) && !isNaN(end) && start >= min && end <= max && start <= end
  }
  
  // 处理步长
  if (field.includes('/')) {
    const [range, step] = field.split('/')
    const stepNum = parseInt(step)
    if (isNaN(stepNum) || stepNum <= 0) return false
    
    if (range === '*') return true
    return isValidCronField(range, min, max)
  }
  
  // 单个数值
  const num = parseInt(field)
  return !isNaN(num) && num >= min && num <= max
}

// 构建cron表达式
const buildCron = () => {
  const { second, minute, hour, day, month, weekday } = cronForm.value
  
  let weekdayValue = weekday
  if (Array.isArray(weekday) && weekday.length > 0) {
    if (weekday.includes('*')) {
      weekdayValue = '*'
    } else {
      weekdayValue = weekday.join(',')
    }
  }
  
  const cron = `${second} ${minute} ${hour} ${day} ${month} ${weekdayValue}`
  cronValue.value = cron
  validateCron(cron)
  emit('update:modelValue', cron)
}

// 解析cron表达式到表单
const parseCronToForm = (cron: string) => {
  if (!cron) return
  
  const parts = cron.trim().split(/\s+/)
  if (parts.length === 6) {
    cronForm.value = {
      second: parts[0],
      minute: parts[1],
      hour: parts[2],
      day: parts[3],
      month: parts[4],
      weekday: parts[5].includes(',') ? parts[5].split(',') : parts[5]
    }
  }
}

// 设置预设值
const setPreset = (type: string) => {
  switch (type) {
    case 'daily':
      cronForm.value = { second: '0', minute: '0', hour: '10', day: '*', month: '*', weekday: '*' }
      break
    case 'weekly':
      cronForm.value = { second: '0', minute: '0', hour: '10', day: '*', month: '*', weekday: '1' }
      break
    case 'workday':
      cronForm.value = { second: '0', minute: '0', hour: '10', day: '*', month: '*', weekday: ['1', '2', '3', '4', '5'] }
      break
    case 'monthly':
      cronForm.value = { second: '0', minute: '0', hour: '10', day: '1', month: '*', weekday: '*' }
      break
  }
  buildCron()
}

// 解析cron描述
const parseCronDescription = (cron: string): string => {
  const parts = cron.trim().split(/\s+/)
  if (parts.length !== 6) return '无效的cron表达式'
  
  const [second, minute, hour, day, month, weekday] = parts
  
  let desc = ''
  
  // 处理时间
  if (hour !== '*' && minute !== '*') {
    desc += `${hour}:${minute.padStart(2, '0')}`
  } else if (hour !== '*') {
    desc += `${hour}点`
  } else if (minute !== '*') {
    desc += `每小时${minute}分`
  } else {
    desc += '每分钟'
  }
  
  // 处理星期
  if (weekday !== '*') {
    const days = weekday.split(',').map(d => {
      const dayMap: Record<string, string> = {
        '0': '周日', '1': '周一', '2': '周二', '3': '周三',
        '4': '周四', '5': '周五', '6': '周六'
      }
      return dayMap[d] || d
    })
    desc += ` ${days.join(',')}`
  } else if (day !== '*') {
    desc += ` 每月${day}日`
  } else {
    desc += ' 每天'
  }
  
  return desc
}

// 初始化
parseCronToForm(props.modelValue)
validateCron(props.modelValue)
</script>

<style scoped>
.cron-editor {
  width: 100%;
}

.cron-input {
  margin-bottom: 8px;
}

.cron-description {
  margin-bottom: 8px;
  padding: 8px 12px;
  background: var(--el-color-info-light-9);
  border-radius: 4px;
  border-left: 3px solid var(--el-color-info);
}

.cron-error {
  margin-bottom: 8px;
  padding: 8px 12px;
  background: var(--el-color-danger-light-9);
  border-radius: 4px;
  border-left: 3px solid var(--el-color-danger);
}

.cron-builder {
  margin-top: 16px;
}

.builder-card {
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.builder-card :deep(.el-card__header) {
  background: rgba(255, 255, 255, 0.05);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}
</style>