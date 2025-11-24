import http from './http'
import type {
  DutySchedule,
  DutyScheduleCreateRequest,
  DutyScheduleUpdateRequest,
  MilestoneEvent,
  MilestoneEventCreateRequest,
  MilestoneEventUpdateRequest,
  PaginatedResponse
} from '../types/cmdb'

// ========== 值班人 API ==========

export const getDutySchedules = async (
  params?: {
    startDate?: string
    endDate?: string
    staffName?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<DutySchedule>> => {
  const res = await http.get<PaginatedResponse<DutySchedule>>(
    '/info/duty/schedules',
    { params }
  )
  return res
}

export const getDutyScheduleById = async (id: number): Promise<DutySchedule> => {
  const res = await http.get<DutySchedule>(`/info/duty/schedules/${id}`)
  return res
}

export const createDutySchedule = async (
  data: DutyScheduleCreateRequest
): Promise<DutySchedule> => {
  const res = await http.post<DutySchedule>('/info/duty/schedules', data)
  return res
}

export const updateDutySchedule = async (
  id: number,
  data: DutyScheduleUpdateRequest
): Promise<DutySchedule> => {
  const res = await http.put<DutySchedule>(`/info/duty/schedules/${id}`, data)
  return res
}

export const deleteDutySchedule = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/duty/schedules/${id}`)
}

// ========== 大事记 API ==========

export const getMilestoneEvents = async (
  params?: {
    startDate?: string
    endDate?: string
    eventContent?: string
    isActive?: boolean
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<MilestoneEvent>> => {
  const res = await http.get<PaginatedResponse<MilestoneEvent>>(
    '/info/milestones',
    { params }
  )
  return res
}

export const getMilestoneEventById = async (id: number): Promise<MilestoneEvent> => {
  const res = await http.get<MilestoneEvent>(`/info/milestones/${id}`)
  return res
}

export const createMilestoneEvent = async (
  data: MilestoneEventCreateRequest
): Promise<MilestoneEvent> => {
  const res = await http.post<MilestoneEvent>('/info/milestones', data)
  return res
}

export const updateMilestoneEvent = async (
  id: number,
  data: MilestoneEventUpdateRequest
): Promise<MilestoneEvent> => {
  const res = await http.put<MilestoneEvent>(`/info/milestones/${id}`, data)
  return res
}

export const deleteMilestoneEvent = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/milestones/${id}`)
}

