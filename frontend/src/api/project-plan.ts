import http from './http'

const request = {
  get: <T = any>(url: string, config?: any) => http.get<T>(url, config),
  post: <T = any>(url: string, data?: any, config?: any) => http.post<T>(url, data, config),
  put: <T = any>(url: string, data?: any, config?: any) => http.put<T>(url, data, config),
  delete: <T = any>(url: string, config?: any) => http.delete<T>(url, config)
}

export interface ProjectModule {
  id: number
  name: string
  icon: string
  sort_order: number
  created_at: string
  updated_at: string
}

export interface ProjectItem {
  id: number
  module_id: number
  name: string
  status: 'not_started' | 'in_progress' | 'completed' | 'blocked'
  priority: 'low' | 'medium' | 'high' | 'urgent'
  owner: string
  estimated_time: string
  notes: string
  sort_order: number
  created_at: string
  updated_at: string
}

export interface ModuleWithItems extends ProjectModule {
  items: ProjectItem[]
}

export const getModulesWithItems = () => {
  return request.get<ModuleWithItems[]>('/project-plan/modules')
}

export const createModule = (data: Partial<ProjectModule>) => {
  return request.post<ProjectModule>('/project-plan/modules', data)
}

export const updateModule = (id: number, data: Partial<ProjectModule>) => {
  return request.put<ProjectModule>(`/project-plan/modules/${id}`, data)
}

export const deleteModule = (id: number) => {
  return request.delete(`/project-plan/modules/${id}`)
}

export const createItem = (data: Partial<ProjectItem>) => {
  return request.post<ProjectItem>('/project-plan/items', data)
}

export const updateItem = (id: number, data: Partial<ProjectItem>) => {
  return request.put<ProjectItem>(`/project-plan/items/${id}`, data)
}

export const deleteItem = (id: number) => {
  return request.delete(`/project-plan/items/${id}`)
}

export const batchUpdateItems = (items: ProjectItem[]) => {
  return request.post('/project-plan/items/batch', items)
}
