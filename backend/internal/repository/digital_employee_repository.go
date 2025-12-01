package repository

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/example/robot-manage/internal/model"
)

type DigitalEmployeeRepository struct {
	db *gorm.DB
}

func NewDigitalEmployeeRepository(db *gorm.DB) *DigitalEmployeeRepository {
	return &DigitalEmployeeRepository{db: db}
}

// ========== 历史会话记录相关 ==========

// GetMessageRecords 获取会话记录列表
func (r *DigitalEmployeeRepository) GetMessageRecords(query *model.MessageRecordQuery) ([]*model.MessageRecord, int64, error) {
	var records []*model.MessageRecord
	var total int64

	db := r.db.Model(&model.MessageRecord{})

	// 排除测试数据
	db = db.Where("msg_id NOT LIKE ?", "test-%")

	// 条件筛选
	if query.MsgID != "" {
		db = db.Where("msg_id = ?", query.MsgID)
	}
	if query.ConversationID != "" {
		db = db.Where("conversation_id LIKE ?", "%"+query.ConversationID+"%")
	}
	if query.RequestUserID != "" {
		db = db.Where("request_user_id LIKE ?", "%"+query.RequestUserID+"%")
	}
	if query.RequestUserName != "" {
		db = db.Where("request_user_name LIKE ?", "%"+query.RequestUserName+"%")
	}
	if query.Debug != "" {
		db = db.Where("debug = ?", query.Debug)
	}
	if query.Type != "" {
		types := strings.Split(query.Type, ",")
		var conditions []string
		var args []interface{}
		for _, t := range types {
			t = strings.TrimSpace(t)
			if t == "null" {
				conditions = append(conditions, "type IS NULL")
			} else if t != "" {
				conditions = append(conditions, "type = ?")
				args = append(args, t)
			}
		}
		if len(conditions) > 0 {
			db = db.Where(strings.Join(conditions, " OR "), args...)
		}
	}
	if query.IsValid != nil {
		db = db.Where("is_valid = ?", *query.IsValid)
	}
	if query.StartTime != "" {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("created_at <= ?", query.EndTime)
	}
	if query.Keyword != "" {
		db = db.Where("input LIKE ? OR output LIKE ?", "%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "created_at"
	}
	orderDir := strings.ToUpper(query.OrderDir)
	if orderDir != "ASC" && orderDir != "DESC" {
		orderDir = "DESC"
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", orderBy, orderDir)).
		Offset(offset).Limit(query.PageSize).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// GetMessageRecordByID 根据ID获取会话记录
func (r *DigitalEmployeeRepository) GetMessageRecordByID(msgID string) (*model.MessageRecord, error) {
	var record model.MessageRecord
	if err := r.db.Where("msg_id = ?", msgID).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// GetMessagesByConversationID 根据会话ID获取所有消息
func (r *DigitalEmployeeRepository) GetMessagesByConversationID(conversationID string) ([]*model.MessageRecord, error) {
	var records []*model.MessageRecord
	if err := r.db.Where("conversation_id = ? AND msg_id NOT LIKE ?", conversationID, "test-%").
		Order("created_at ASC").Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// CreateMessageRecord 创建会话记录
func (r *DigitalEmployeeRepository) CreateMessageRecord(record *model.MessageRecord) error {
	now := time.Now()
	record.CreatedAt = &now
	record.UpdatedAt = &now
	return r.db.Create(record).Error
}

// UpdateMessageRecord 更新会话记录
func (r *DigitalEmployeeRepository) UpdateMessageRecord(msgID string, record *model.MessageRecord) error {
	now := time.Now()
	record.UpdatedAt = &now
	return r.db.Model(&model.MessageRecord{}).Where("msg_id = ?", msgID).Updates(record).Error
}

// DeleteMessageRecord 删除会话记录
func (r *DigitalEmployeeRepository) DeleteMessageRecord(msgID string) error {
	return r.db.Where("msg_id = ?", msgID).Delete(&model.MessageRecord{}).Error
}

// BatchDeleteMessageRecords 批量删除会话记录
func (r *DigitalEmployeeRepository) BatchDeleteMessageRecords(msgIDs []string) error {
	return r.db.Where("msg_id IN ?", msgIDs).Delete(&model.MessageRecord{}).Error
}

// BatchUpdateDebugStatus 批量更新调试状态
func (r *DigitalEmployeeRepository) BatchUpdateDebugStatus(msgIDs []string, debug string) error {
	now := time.Now()
	return r.db.Model(&model.MessageRecord{}).
		Where("msg_id IN ?", msgIDs).
		Updates(map[string]interface{}{
			"debug":      debug,
			"updated_at": now,
		}).Error
}

// GetConversationGroups 获取会话分组
func (r *DigitalEmployeeRepository) GetConversationGroups(limit int) ([]*model.ConversationGroup, error) {
	var groups []*model.ConversationGroup

	err := r.db.Model(&model.MessageRecord{}).
		Where("msg_id NOT LIKE ?", "test-%").
		Select("conversation_id, request_user_id, request_user_name, COUNT(*) as message_count, MIN(created_at) as first_message_at, MAX(created_at) as last_message_at").
		Group("conversation_id, request_user_id, request_user_name").
		Order("last_message_at DESC").
		Limit(limit).
		Scan(&groups).Error

	return groups, err
}

// GetMessageStatistics 获取会话统计
func (r *DigitalEmployeeRepository) GetMessageStatistics() (*model.MessageStatistics, error) {
	stats := &model.MessageStatistics{}

	// 总消息数
	r.db.Model(&model.MessageRecord{}).Where("msg_id NOT LIKE ?", "test-%").Count(&stats.TotalMessages)

	// 总会话数
	r.db.Model(&model.MessageRecord{}).
		Where("msg_id NOT LIKE ?", "test-%").
		Distinct("conversation_id").
		Count(&stats.TotalConversations)

	// 总用户数
	r.db.Model(&model.MessageRecord{}).
		Where("msg_id NOT LIKE ?", "test-%").
		Distinct("request_user_id").
		Count(&stats.TotalUsers)

	// 调试消息数
	r.db.Model(&model.MessageRecord{}).
		Where("msg_id NOT LIKE ? AND debug IS NOT NULL AND debug != ''", "test-%").
		Count(&stats.DebugMessages)

	// 类型分布
	var typeDistribution []model.TypeDistribution
	r.db.Model(&model.MessageRecord{}).
		Select("type, COUNT(*) as count").
		Where("msg_id NOT LIKE ? AND type IS NOT NULL AND type != ''", "test-%").
		Group("type").
		Order("count DESC").
		Scan(&typeDistribution)
	stats.TypeDistribution = typeDistribution

	// 今日消息数
	today := time.Now().Format("2006-01-02")
	r.db.Model(&model.MessageRecord{}).
		Where("msg_id NOT LIKE ? AND DATE(created_at) = ?", "test-%", today).
		Count(&stats.TodayMessages)

	// 最近7天趋势
	var trendData []model.TrendDataPoint
	r.db.Model(&model.MessageRecord{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("msg_id NOT LIKE ? AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)", "test-%").
		Group("DATE(created_at)").
		Order("date ASC").
		Scan(&trendData)
	stats.TrendData = trendData

	// Top 10 活跃用户
	var topUsers []model.UserActivityData
	r.db.Model(&model.MessageRecord{}).
		Select("request_user_id as user_id, request_user_name as user_name, COUNT(*) as count").
		Where("msg_id NOT LIKE ?", "test-%").
		Group("request_user_id, request_user_name").
		Order("count DESC").
		Limit(10).
		Scan(&topUsers)
	stats.TopUsers = topUsers

	return stats, nil
}

// ========== 服务回传记录相关 ==========

// GetExportRecords 获取导出记录列表
func (r *DigitalEmployeeRepository) GetExportRecords(query *model.ExportRecordQuery) ([]*model.ExportRecord, int64, error) {
	var records []*model.ExportRecord
	var total int64

	db := r.db.Model(&model.ExportRecord{})

	// 条件筛选
	if query.FileName != "" {
		db = db.Where("file_name LIKE ?", "%"+query.FileName+"%")
	}
	if query.EmailSubject != "" {
		db = db.Where("email_subject LIKE ?", "%"+query.EmailSubject+"%")
	}
	if query.EmailStatus != "" {
		db = db.Where("email_status = ?", query.EmailStatus)
	}
	if query.StartDate != "" {
		db = db.Where("start_date >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("end_date <= ?", query.EndDate)
	}
	if query.StartTime != "" {
		db = db.Where("export_time >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("export_time <= ?", query.EndTime)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "export_time"
	}
	orderDir := strings.ToUpper(query.OrderDir)
	if orderDir != "ASC" && orderDir != "DESC" {
		orderDir = "DESC"
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", orderBy, orderDir)).
		Offset(offset).Limit(query.PageSize).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// GetExportRecordByID 根据ID获取导出记录
func (r *DigitalEmployeeRepository) GetExportRecordByID(id uint) (*model.ExportRecord, error) {
	var record model.ExportRecord
	if err := r.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateExportRecord 创建导出记录
func (r *DigitalEmployeeRepository) CreateExportRecord(record *model.ExportRecord) error {
	now := time.Now()
	record.ExportTime = now
	record.CreatedAt = &now
	record.UpdatedAt = &now
	return r.db.Create(record).Error
}

// UpdateExportRecord 更新导出记录
func (r *DigitalEmployeeRepository) UpdateExportRecord(id uint, record *model.ExportRecord) error {
	now := time.Now()
	record.UpdatedAt = &now
	return r.db.Model(&model.ExportRecord{}).Where("id = ?", id).Updates(record).Error
}

// UpdateEmailStatus 更新邮件状态
func (r *DigitalEmployeeRepository) UpdateEmailStatus(id uint, status string, sendTime *time.Time, errorReason string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"email_status": status,
		"updated_at":   now,
	}
	if sendTime != nil {
		updates["email_send_time"] = sendTime
	}
	if errorReason != "" {
		updates["error_reason"] = errorReason
	}
	return r.db.Model(&model.ExportRecord{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteExportRecord 删除导出记录
func (r *DigitalEmployeeRepository) DeleteExportRecord(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.ExportRecord{}).Error
}

// BatchDeleteExportRecords 批量删除导出记录
func (r *DigitalEmployeeRepository) BatchDeleteExportRecords(ids []uint) error {
	return r.db.Where("id IN ?", ids).Delete(&model.ExportRecord{}).Error
}

// GetExportStatistics 获取导出统计
func (r *DigitalEmployeeRepository) GetExportStatistics() (*model.ExportStatistics, error) {
	stats := &model.ExportStatistics{}

	// 总导出数
	r.db.Model(&model.ExportRecord{}).Count(&stats.TotalExports)

	// 成功数
	r.db.Model(&model.ExportRecord{}).
		Where("email_status = ?", "success").
		Count(&stats.SuccessCount)

	// 失败数
	r.db.Model(&model.ExportRecord{}).
		Where("email_status = ?", "failed").
		Count(&stats.FailedCount)

	// 待发送数
	r.db.Model(&model.ExportRecord{}).
		Where("email_status IS NULL OR email_status = '' OR email_status = 'pending'").
		Count(&stats.PendingCount)

	// 成功率
	if stats.TotalExports > 0 {
		stats.SuccessRate = float64(stats.SuccessCount) / float64(stats.TotalExports) * 100
	}

	// 总记录数
	r.db.Model(&model.ExportRecord{}).
		Select("COALESCE(SUM(total_records), 0)").
		Scan(&stats.TotalRecords)

	// 最近7天趋势
	var trendData []model.TrendDataPoint
	r.db.Model(&model.ExportRecord{}).
		Select("DATE(export_time) as date, COUNT(*) as count").
		Where("export_time >= DATE_SUB(NOW(), INTERVAL 7 DAY)").
		Group("DATE(export_time)").
		Order("date ASC").
		Scan(&trendData)
	stats.TrendData = trendData

	return stats, nil
}
