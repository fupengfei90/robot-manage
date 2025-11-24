import http from './http'

const request = {
  get: <T = any>(url: string, config?: any) => http.get<T>(url, config),
  post: <T = any>(url: string, data?: any, config?: any) => http.post<T>(url, data, config),
  put: <T = any>(url: string, data?: any, config?: any) => http.put<T>(url, data, config),
  delete: <T = any>(url: string, config?: any) => http.delete<T>(url, config)
}

export interface ScheduleTask {
  id: number
  name: string
  active: number
  cron: string
  workflow: string
  exec_token: string
  created_at: string
  updated_at: string
  deleted_at?: string
  cron_entry_id: number
  category: string
  dcn: string
  next_run_time?: string
  cron_desc: string
  status_text: string
}

export interface ScheduleTaskRequest {
  name: string
  active: number
  cron: string
  workflow: string
  exec_token: string
  category: string
  dcn: string
}

export interface ScheduleTaskQuery {
  name?: string
  active?: number
  category?: string
  dcn?: string
  start_time?: string
  end_time?: string
  page?: number
  page_size?: number
  order_by?: string
  order_dir?: string
}

export interface BatchOperationRequest {
  ids: number[]
  operation: 'enable' | 'disable' | 'delete'
}

export interface ScheduleTaskListResponse {
  list: ScheduleTask[]
  total: number
  page: number
  page_size: number
}

// 获取任务列表
export const getScheduleTaskList = (params: ScheduleTaskQuery) => {
  return request.get<ScheduleTaskListResponse>('/system/schedule-tasks', { params })
}

// 获取任务详情
export const getScheduleTaskById = (id: number) => {
  return request.get<ScheduleTask>(`/system/schedule-tasks/${id}`)
}

// 创建任务
export const createScheduleTask = (data: ScheduleTaskRequest) => {
  return request.post<ScheduleTask>('/system/schedule-tasks', data)
}

// 更新任务
export const updateScheduleTask = (id: number, data: ScheduleTaskRequest) => {
  return request.put<ScheduleTask>(`/system/schedule-tasks/${id}`, data)
}

// 删除任务
export const deleteScheduleTask = (id: number) => {
  return request.delete(`/system/schedule-tasks/${id}`)
}

// 更新任务状态
export const updateScheduleTaskStatus = (id: number, active: number) => {
  return request.put(`/system/schedule-tasks/${id}/status`, { active })
}

// 立即执行任务
export const executeScheduleTask = (id: number) => {
  return request.post(`/system/schedule-tasks/${id}/execute`)
}

// 批量操作
export const batchOperateScheduleTasks = (data: BatchOperationRequest) => {
  return request.post('/system/schedule-tasks/batch', data)
}

// 获取分类列表
export const getScheduleTaskCategories = () => {
  return request.get<string[]>('/system/schedule-tasks/categories')
}

// 获取数据中心列表
export const getScheduleTaskDCNs = () => {
  return request.get<string[]>('/system/schedule-tasks/dcns')
}