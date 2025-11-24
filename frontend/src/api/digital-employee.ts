import http from './http'

// ========== 历史会话记录 API ==========

export interface MessageRecord {
  msg_id: string
  input: string
  output: string
  conversation_id: string
  created_at: string
  request_user_id: string
  request_user_name: string
  extra_data: string
  updated_at: string
  debug: string
  type: string
  is_valid: number
}

export interface MessageRecordQuery {
  msg_id?: string
  conversation_id?: string
  request_user_id?: string
  request_user_name?: string
  debug?: string
  type?: string
  is_valid?: number
  start_time?: string
  end_time?: string
  keyword?: string
  page?: number
  page_size?: number
  order_by?: string
  order_dir?: string
}

export interface MessageRecordRequest {
  msg_id: string
  input?: string
  output?: string
  conversation_id?: string
  request_user_id?: string
  request_user_name?: string
  extra_data?: string
  debug?: string
}

export interface ConversationGroup {
  conversation_id: string
  request_user_id: string
  request_user_name: string
  message_count: number
  first_message_at: string
  last_message_at: string
  messages: MessageRecord[]
}

export interface MessageStatistics {
  total_messages: number
  total_conversations: number
  total_users: number
  debug_messages: number
  today_messages: number
  type_distribution: Array<{ type: string; count: number }>
  trend_data: Array<{ date: string; count: number }>
  top_users: Array<{ user_id: string; user_name: string; count: number }>
}

// 获取会话记录列表
export const getMessageRecords = (params: MessageRecordQuery) => {
  return http.get('/digital/message-records', { params })
}

// 获取会话记录详情
export const getMessageRecordById = (msgId: string) => {
  return http.get(`/digital/message-records/${msgId}`)
}

// 根据会话ID获取所有消息
export const getMessagesByConversationId = (conversationId: string) => {
  return http.get(`/digital/conversations/${conversationId}/messages`)
}

// 创建会话记录
export const createMessageRecord = (data: MessageRecordRequest) => {
  return http.post('/digital/message-records', data)
}

// 更新会话记录
export const updateMessageRecord = (msgId: string, data: MessageRecordRequest) => {
  return http.put(`/digital/message-records/${msgId}`, data)
}

// 删除会话记录
export const deleteMessageRecord = (msgId: string) => {
  return http.delete(`/digital/message-records/${msgId}`)
}

// 批量操作会话记录
export const batchOperateMessageRecords = (data: {
  msg_ids: string[]
  operation: 'delete' | 'mark_debug' | 'unmark_debug'
}) => {
  return http.post('/digital/message-records/batch', data)
}

// 获取会话分组
export const getConversationGroups = (limit = 20) => {
  return http.get('/digital/conversation-groups', { params: { limit } })
}

// 获取会话统计
export const getMessageStatistics = () => {
  return http.get('/digital/message-statistics')
}

// ========== 服务回传记录 API ==========

export interface ExportRecord {
  id: number
  export_time: string
  start_date: string
  end_date: string
  total_records: number
  file_name: string
  file_path: string
  email_send_time?: string
  email_subject?: string
  email_status?: string
  email_recipients?: string
  email_cc?: string
  email_bcc?: string
  error_reason?: string
  created_at: string
  updated_at: string
}

export interface ExportRecordQuery {
  file_name?: string
  email_subject?: string
  email_status?: string
  start_date?: string
  end_date?: string
  start_time?: string
  end_time?: string
  page?: number
  page_size?: number
  order_by?: string
  order_dir?: string
}

export interface ExportRecordRequest {
  start_date: string
  end_date: string
  total_records: number
  file_name: string
  file_path: string
  email_subject?: string
  email_recipients?: string
  email_cc?: string
  email_bcc?: string
}

export interface ExportStatistics {
  total_exports: number
  success_count: number
  failed_count: number
  pending_count: number
  success_rate: number
  total_records: number
  trend_data: Array<{ date: string; count: number }>
}

// 获取导出记录列表
export const getExportRecords = (params: ExportRecordQuery) => {
  return http.get('/digital/export-records', { params })
}

// 获取导出记录详情
export const getExportRecordById = (id: number) => {
  return http.get(`/digital/export-records/${id}`)
}

// 创建导出记录
export const createExportRecord = (data: ExportRecordRequest) => {
  return http.post('/digital/export-records', data)
}

// 更新导出记录
export const updateExportRecord = (id: number, data: ExportRecordRequest) => {
  return http.put(`/digital/export-records/${id}`, data)
}

// 删除导出记录
export const deleteExportRecord = (id: number) => {
  return http.delete(`/digital/export-records/${id}`)
}

// 批量操作导出记录
export const batchOperateExportRecords = (data: {
  ids: number[]
  operation: 'delete' | 'retry'
}) => {
  return http.post('/digital/export-records/batch', data)
}

// 重试发送邮件
export const retryExportEmail = (id: number) => {
  return http.post(`/digital/export-records/${id}/retry`)
}

// 获取导出统计
export const getExportStatistics = () => {
  return http.get('/digital/export-statistics')
}
