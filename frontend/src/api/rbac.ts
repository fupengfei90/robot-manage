import http from './http'

// 用户相关接口
export interface User {
  id: number
  user_id: string
  name: string
  active: number
  p_id: string
  phone: string
  created_at: string
  updated_at: string
  roles: Role[]
  role_names: string[]
}

export interface UserRequest {
  user_id: string
  name: string
  active: number
  p_id?: string
  phone?: string
  role_ids: number[]
}

export interface UserQuery {
  user_id?: string
  name?: string
  active?: number
  role_id?: number
  start_time?: string
  end_time?: string
  page?: number
  page_size?: number
  order_by?: string
  order_dir?: string
}

// 角色相关接口
export interface Role {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
  users: User[]
  permissions: Permission[]
  user_count: number
  permission_count: number
  permission_names: string[]
}

export interface RoleRequest {
  name: string
  description: string
  permission_ids: number[]
}

export interface RoleQuery {
  name?: string
  page?: number
  page_size?: number
  order_by?: string
  order_dir?: string
}

// 权限相关接口
export interface Permission {
  id: number
  intent: string
  intent2: string
  description: string
  created_at: string
  updated_at: string
}

export interface PermissionRequest {
  intent: string
  intent2?: string
  description: string
}

export interface PermissionQuery {
  intent?: string
  intent2?: string
  page?: number
  page_size?: number
}

export interface PermissionTreeNode {
  intent: string
  description: string
  children?: PermissionTreeNode[]
  id?: number
  intent2?: string
}

// 批量操作
export interface BatchUserOperationRequest {
  user_ids: number[]
  operation: 'enable' | 'disable' | 'delete'
}

export interface UserRoleAssignRequest {
  role_ids: number[]
}

export interface RolePermissionAssignRequest {
  permission_ids: number[]
}

// 响应类型
export interface ListResponse<T> {
  list: T[]
  total: number
  page: number
  page_size: number
}

// 用户管理API
export const getUsers = (params: UserQuery) => {
  return http.get<ListResponse<User>>('/system/users', { params })
}

export const getUserById = (id: number) => {
  return http.get<User>(`/system/users/${id}`)
}

export const createUser = (data: UserRequest) => {
  return http.post<User>('/system/users', data)
}

export const updateUser = (id: number, data: UserRequest) => {
  return http.put<User>(`/system/users/${id}`, data)
}

export const deleteUser = (id: number) => {
  return http.delete(`/system/users/${id}`)
}

export const updateUserStatus = (id: number, active: number) => {
  return http.put(`/system/users/${id}/status`, { active })
}

export const batchOperateUsers = (data: BatchUserOperationRequest) => {
  return http.post('/system/users/batch', data)
}

export const assignUserRoles = (id: number, data: UserRoleAssignRequest) => {
  return http.post(`/system/users/${id}/roles`, data)
}

// 角色管理API
export const getRoles = (params: RoleQuery) => {
  return http.get<ListResponse<Role>>('/system/roles', { params })
}

export const getRoleById = (id: number) => {
  return http.get<Role>(`/system/roles/${id}`)
}

export const createRole = (data: RoleRequest) => {
  return http.post<Role>('/system/roles', data)
}

export const updateRole = (id: number, data: RoleRequest) => {
  return http.put<Role>(`/system/roles/${id}`, data)
}

export const deleteRole = (id: number) => {
  return http.delete(`/system/roles/${id}`)
}

export const assignRolePermissions = (id: number, data: RolePermissionAssignRequest) => {
  return http.post(`/system/roles/${id}/permissions`, data)
}

// 权限管理API
export const getPermissions = (params: PermissionQuery) => {
  return http.get<ListResponse<Permission>>('/system/permissions', { params })
}

export const getPermissionTree = () => {
  return http.get<PermissionTreeNode[]>('/system/permissions/tree')
}

export const getPermissionById = (id: number) => {
  return http.get<Permission>(`/system/permissions/${id}`)
}

export const createPermission = (data: PermissionRequest) => {
  return http.post<Permission>('/system/permissions', data)
}

export const updatePermission = (id: number, data: PermissionRequest) => {
  return http.put<Permission>(`/system/permissions/${id}`, data)
}

export const deletePermission = (id: number) => {
  return http.delete(`/system/permissions/${id}`)
}