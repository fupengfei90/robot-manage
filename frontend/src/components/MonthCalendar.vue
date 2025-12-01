<template>
  <div class="month-calendar">
    <div class="calendar-header">
      <el-button @click="prevMonth" :icon="ArrowLeft" circle />
      <span class="current-month">{{ currentMonthText }}</span>
      <el-button @click="nextMonth" :icon="ArrowRight" circle />
    </div>
    <div class="calendar-grid">
      <div v-for="day in weekDays" :key="day" class="weekday-header">{{ day }}</div>
      <div
        v-for="date in calendarDates"
        :key="date.key"
        :class="['calendar-cell', { 'other-month': !date.isCurrentMonth, 'today': date.isToday }]"
      >
        <div class="date-number">{{ date.day }}</div>
        <div class="date-content">
          <slot :date="date.fullDate" :items="getItemsForDate(date.fullDate)" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'

interface CalendarDate {
  day: number
  fullDate: string
  isCurrentMonth: boolean
  isToday: boolean
  key: string
}

interface Props {
  items?: Array<{ date: string; [key: string]: any }>
}

const props = defineProps<Props>()

const currentDate = ref(new Date())
const weekDays = ['日', '一', '二', '三', '四', '五', '六']

const currentMonthText = computed(() => {
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth() + 1
  return `${year}年${month}月`
})

const calendarDates = computed(() => {
  const year = currentDate.value.getFullYear()
  const month = currentDate.value.getMonth()
  const firstDay = new Date(year, month, 1)
  const lastDay = new Date(year, month + 1, 0)
  const prevMonthLastDay = new Date(year, month, 0)
  
  const dates: CalendarDate[] = []
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  
  // 上月日期
  const firstDayOfWeek = firstDay.getDay()
  for (let i = firstDayOfWeek - 1; i >= 0; i--) {
    const day = prevMonthLastDay.getDate() - i
    const date = new Date(year, month - 1, day)
    dates.push({
      day,
      fullDate: formatDate(date),
      isCurrentMonth: false,
      isToday: false,
      key: `prev-${day}`
    })
  }
  
  // 当月日期
  for (let day = 1; day <= lastDay.getDate(); day++) {
    const date = new Date(year, month, day)
    dates.push({
      day,
      fullDate: formatDate(date),
      isCurrentMonth: true,
      isToday: date.getTime() === today.getTime(),
      key: `current-${day}`
    })
  }
  
  // 下月日期
  const remainingCells = 42 - dates.length
  for (let day = 1; day <= remainingCells; day++) {
    const date = new Date(year, month + 1, day)
    dates.push({
      day,
      fullDate: formatDate(date),
      isCurrentMonth: false,
      isToday: false,
      key: `next-${day}`
    })
  }
  
  return dates
})

const getItemsForDate = (date: string) => {
  if (!props.items) return []
  return props.items.filter(item => item.date === date)
}

const formatDate = (date: Date) => {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const prevMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() - 1, 1)
}

const nextMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() + 1, 1)
}

defineExpose({ currentDate })
</script>

<style scoped>
.month-calendar {
  width: 100%;
}

.calendar-header {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  margin-bottom: 20px;
}

.current-month {
  font-size: 18px;
  font-weight: 600;
  min-width: 120px;
  text-align: center;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1px;
  background: var(--border-color);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  overflow: hidden;
}

.weekday-header {
  background: var(--bg-card);
  padding: 12px;
  text-align: center;
  font-weight: 600;
  color: var(--text-primary);
}

.calendar-cell {
  background: var(--bg-card);
  min-height: 120px;
  padding: 8px;
  position: relative;
  transition: all 0.2s;
}

.calendar-cell:hover {
  background: var(--bg-glass-hover);
}

.calendar-cell.other-month {
  opacity: 0.4;
}

.calendar-cell.today {
  background: rgba(64, 158, 255, 0.1);
}

.date-number {
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 4px;
  color: var(--text-primary);
}

.calendar-cell.today .date-number {
  color: var(--el-color-primary);
}

.date-content {
  font-size: 12px;
  overflow-y: auto;
  max-height: 95px;
}
</style>
