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

