export interface DutySchedule {
  id: number
  dutyDate: string
  wbStaffName: string
  fbStaffName: string
  notified: boolean
  createdAt: string
}

export interface DutyScheduleCreateRequest {
  dutyDate: string
  wbStaffName: string
  fbStaffName: string
}

export interface DutyScheduleUpdateRequest {
  wbStaffName?: string
  fbStaffName?: string
  notified?: boolean
}

export interface MilestoneEvent {
  id: number
  eventDate: string
  dayOfWeek: string
  eventContent: string
  isActive: boolean
  createdAt: string
  updatedAt: string
}

export interface MilestoneEventCreateRequest {
  eventDate: string
  eventContent: string
}

export interface MilestoneEventUpdateRequest {
  eventDate?: string
  eventContent?: string
  isActive?: boolean
}

export interface PaginatedResponse<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

// 批量时间相关类型
export interface BatchTime {
  id: number
  systemName: string
  subsysName: string
  batchTime: string
  createdAt: string
  updatedAt: string
}

export interface BatchTimeCreateRequest {
  systemName: string
  subsysName: string
  batchTime: string
}

export interface BatchTimeUpdateRequest {
  systemName?: string
  subsysName?: string
  batchTime?: string
}

export interface System {
  id: number
  name: string
  code: string
  description: string
  owner: string
  status: string
  createdAt: string
  updatedAt: string
}

export interface Subsystem {
  id: number
  systemId: number
  name: string
  code: string
  description: string
  batchWindow: string
  createdAt: string
  updatedAt: string
}


// WB CMDB相关类型
export interface WBCMDBInfo {
  id: number
  systemName: string
  environment: string
  ipAddress: string
  port: string
  status: string
  owner: string
  remark: string
  createdAt: string
  updatedAt: string
}

export interface WBCMDBInfoRequest {
  systemName: string
  environment?: string
  ipAddress?: string
  port?: string
  status?: string
  owner?: string
  remark?: string
}

// VB CMDB相关类型
export interface VBCMDBInfo {
  id: number
  systemName: string
  environment: string
  ipAddress: string
  port: string
  status: string
  owner: string
  remark: string
  createdAt: string
  updatedAt: string
}

export interface VBCMDBInfoRequest {
  systemName: string
  environment?: string
  ipAddress?: string
  port?: string
  status?: string
  owner?: string
  remark?: string
}

// ITSM出包记录相关类型
export interface ITSMPackageRecord {
  id: number
  subsystem: string
  packageName: string
  requirementId: string
  itsmTicket: string
  status: string
  owner: string
  environment: string
  createdAt: string
  updatedAt: string
}

export interface ITSMPackageRecordRequest {
  subsystem: string
  packageName: string
  requirementId: string
  itsmTicket: string
  status: string
  owner: string
  environment: string
}
