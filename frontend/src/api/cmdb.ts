import http from './http'
import type {
  DutySchedule,
  DutyScheduleCreateRequest,
  DutyScheduleUpdateRequest,
  MilestoneEvent,
  MilestoneEventCreateRequest,
  MilestoneEventUpdateRequest,
  BatchTime,
  BatchTimeCreateRequest,
  BatchTimeUpdateRequest,
  System,
  Subsystem,
  WBCMDBInfo,
  WBCMDBInfoRequest,
  VBCMDBInfo,
  VBCMDBInfoRequest,
  ITSMPackageRecord,
  ITSMPackageRecordRequest,
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

// ========== 批量时间 API ==========

export const getBatchTimes = async (
  params?: {
    systemName?: string
    subsysName?: string
    batchTime?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<BatchTime>> => {
  const res = await http.get<PaginatedResponse<BatchTime>>(
    '/info/batch-times',
    { params }
  )
  return res
}

export const getBatchTimeById = async (id: number): Promise<BatchTime> => {
  const res = await http.get<BatchTime>(`/info/batch-times/${id}`)
  return res
}

export const createBatchTime = async (
  data: BatchTimeCreateRequest
): Promise<BatchTime> => {
  const res = await http.post<BatchTime>('/info/batch-times', data)
  return res
}

export const updateBatchTime = async (
  id: number,
  data: BatchTimeUpdateRequest
): Promise<BatchTime> => {
  const res = await http.put<BatchTime>(`/info/batch-times/${id}`, data)
  return res
}

export const deleteBatchTime = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/batch-times/${id}`)
}

export const getSystems = async (): Promise<System[]> => {
  const res = await http.get<System[]>('/info/systems')
  return res
}

export const getSubsystemsBySystemId = async (systemId: number): Promise<Subsystem[]> => {
  const res = await http.get<Subsystem[]>(`/info/systems/${systemId}/subsystems`)
  return res
}


// ========== WB CMDB API ==========

export const getWBCMDBList = async (
  params?: {
    systemName?: string
    environment?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<WBCMDBInfo>> => {
  const res = await http.get<PaginatedResponse<WBCMDBInfo>>(
    '/info/wb-cmdb',
    { params }
  )
  return res
}

export const getWBCMDBById = async (id: number): Promise<WBCMDBInfo> => {
  const res = await http.get<WBCMDBInfo>(`/info/wb-cmdb/${id}`)
  return res
}

export const createWBCMDB = async (
  data: WBCMDBInfoRequest
): Promise<WBCMDBInfo> => {
  const res = await http.post<WBCMDBInfo>('/info/wb-cmdb', data)
  return res
}

export const updateWBCMDB = async (
  id: number,
  data: WBCMDBInfoRequest
): Promise<WBCMDBInfo> => {
  const res = await http.put<WBCMDBInfo>(`/info/wb-cmdb/${id}`, data)
  return res
}

export const deleteWBCMDB = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/wb-cmdb/${id}`)
}

// ========== VB CMDB API ==========

export const getVBCMDBList = async (
  params?: {
    systemName?: string
    environment?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<VBCMDBInfo>> => {
  const res = await http.get<PaginatedResponse<VBCMDBInfo>>(
    '/info/vb-cmdb',
    { params }
  )
  return res
}

export const getVBCMDBById = async (id: number): Promise<VBCMDBInfo> => {
  const res = await http.get<VBCMDBInfo>(`/info/vb-cmdb/${id}`)
  return res
}

export const createVBCMDB = async (
  data: VBCMDBInfoRequest
): Promise<VBCMDBInfo> => {
  const res = await http.post<VBCMDBInfo>('/info/vb-cmdb', data)
  return res
}

export const updateVBCMDB = async (
  id: number,
  data: VBCMDBInfoRequest
): Promise<VBCMDBInfo> => {
  const res = await http.put<VBCMDBInfo>(`/info/vb-cmdb/${id}`, data)
  return res
}

export const deleteVBCMDB = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/vb-cmdb/${id}`)
}

// ========== ITSM出包记录 API ==========

export const getITSMPackageRecords = async (
  params?: {
    subsystem?: string
    status?: string
    environment?: string
    page?: number
    pageSize?: number
  }
): Promise<PaginatedResponse<ITSMPackageRecord>> => {
  const res = await http.get<PaginatedResponse<ITSMPackageRecord>>(
    '/info/itsm-packages',
    { params }
  )
  return res
}

export const getITSMPackageRecordById = async (id: number): Promise<ITSMPackageRecord> => {
  const res = await http.get<ITSMPackageRecord>(`/info/itsm-packages/${id}`)
  return res
}

export const createITSMPackageRecord = async (
  data: ITSMPackageRecordRequest
): Promise<ITSMPackageRecord> => {
  const res = await http.post<ITSMPackageRecord>('/info/itsm-packages', data)
  return res
}

export const updateITSMPackageRecord = async (
  id: number,
  data: ITSMPackageRecordRequest
): Promise<ITSMPackageRecord> => {
  const res = await http.put<ITSMPackageRecord>(`/info/itsm-packages/${id}`, data)
  return res
}

export const deleteITSMPackageRecord = async (id: number): Promise<void> => {
  await http.delete<void>(`/info/itsm-packages/${id}`)
}
