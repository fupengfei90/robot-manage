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
import type { ApiResponse } from '../types/api'

// ========== 值班排班 API ==========

export const getDutySchedules = async (
  params?: {
    startDate?: string
    endDate?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<DutySchedule>> => {
  const res = await http.get<ApiResponse<PaginatedResponse<DutySchedule>>, ApiResponse<PaginatedResponse<DutySchedule>>>(
    '/info/duty/schedules',
    { params }
  )
  return res.data
}

export const getDutyScheduleById = async (id: number): Promise<DutySchedule> => {
  const res = await http.get<ApiResponse<DutySchedule>, ApiResponse<DutySchedule>>(
    `/info/duty/schedules/${id}`
  )
  return res.data
}

export const createDutySchedule = async (
  data: DutyScheduleCreateRequest
): Promise<DutySchedule> => {
  const res = await http.post<ApiResponse<DutySchedule>, ApiResponse<DutySchedule>>(
    '/info/duty/schedules',
    data
  )
  return res.data
}

export const updateDutySchedule = async (
  id: number,
  data: DutyScheduleUpdateRequest
): Promise<DutySchedule> => {
  const res = await http.put<ApiResponse<DutySchedule>, ApiResponse<DutySchedule>>(
    `/info/duty/schedules/${id}`,
    data
  )
  return res.data
}

export const deleteDutySchedule = async (id: number): Promise<void> => {
  await http.delete<ApiResponse<void>>(`/info/duty/schedules/${id}`)
}

// ========== 大事记 API ==========

export const getMilestoneEvents = async (
  params?: {
    startDate?: string
    endDate?: string
    isActive?: boolean
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<MilestoneEvent>> => {
  const res = await http.get<ApiResponse<PaginatedResponse<MilestoneEvent>>, ApiResponse<PaginatedResponse<MilestoneEvent>>>(
    '/info/milestones',
    { params }
  )
  return res.data
}

export const getMilestoneEventById = async (id: number): Promise<MilestoneEvent> => {
  const res = await http.get<ApiResponse<MilestoneEvent>, ApiResponse<MilestoneEvent>>(
    `/info/milestones/${id}`
  )
  return res.data
}

export const createMilestoneEvent = async (
  data: MilestoneEventCreateRequest
): Promise<MilestoneEvent> => {
  const res = await http.post<ApiResponse<MilestoneEvent>, ApiResponse<MilestoneEvent>>(
    '/info/milestones',
    data
  )
  return res.data
}

export const updateMilestoneEvent = async (
  id: number,
  data: MilestoneEventUpdateRequest
): Promise<MilestoneEvent> => {
  const res = await http.put<ApiResponse<MilestoneEvent>, ApiResponse<MilestoneEvent>>(
    `/info/milestones/${id}`,
    data
  )
  return res.data
}

export const deleteMilestoneEvent = async (id: number): Promise<void> => {
  await http.delete<ApiResponse<void>>(`/info/milestones/${id}`)
}

