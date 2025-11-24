package service

import (
	"fmt"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/repository"
)

type DigitalEmployeeService struct {
	repo *repository.DigitalEmployeeRepository
}

func NewDigitalEmployeeService(repo *repository.DigitalEmployeeRepository) *DigitalEmployeeService {
	return &DigitalEmployeeService{repo: repo}
}

// ========== 历史会话记录相关 ==========

// GetMessageRecords 获取会话记录列表
func (s *DigitalEmployeeService) GetMessageRecords(query *model.MessageRecordQuery) ([]*model.MessageRecord, int64, error) {
	return s.repo.GetMessageRecords(query)
}

// GetMessageRecordByID 根据ID获取会话记录
func (s *DigitalEmployeeService) GetMessageRecordByID(msgID string) (*model.MessageRecord, error) {
	return s.repo.GetMessageRecordByID(msgID)
}

// GetMessagesByConversationID 根据会话ID获取所有消息
func (s *DigitalEmployeeService) GetMessagesByConversationID(conversationID string) ([]*model.MessageRecord, error) {
	return s.repo.GetMessagesByConversationID(conversationID)
}

// CreateMessageRecord 创建会话记录
func (s *DigitalEmployeeService) CreateMessageRecord(req *model.MessageRecordRequest) (*model.MessageRecord, error) {
	record := &model.MessageRecord{
		MsgID:           req.MsgID,
		Input:           req.Input,
		Output:          req.Output,
		ConversationID:  req.ConversationID,
		RequestUserID:   req.RequestUserID,
		RequestUserName: req.RequestUserName,
		ExtraData:       req.ExtraData,
		Debug:           req.Debug,
		Type:            req.Type,
		IsValid:         req.IsValid,
	}

	if err := s.repo.CreateMessageRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateMessageRecord 更新会话记录
func (s *DigitalEmployeeService) UpdateMessageRecord(msgID string, req *model.MessageRecordRequest) (*model.MessageRecord, error) {
	existing, err := s.repo.GetMessageRecordByID(msgID)
	if err != nil {
		return nil, fmt.Errorf("会话记录不存在: %w", err)
	}

	existing.Input = req.Input
	existing.Output = req.Output
	existing.ConversationID = req.ConversationID
	existing.RequestUserID = req.RequestUserID
	existing.RequestUserName = req.RequestUserName
	existing.ExtraData = req.ExtraData
	existing.Debug = req.Debug
	existing.Type = req.Type
	existing.IsValid = req.IsValid

	if err := s.repo.UpdateMessageRecord(msgID, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

// DeleteMessageRecord 删除会话记录
func (s *DigitalEmployeeService) DeleteMessageRecord(msgID string) error {
	if _, err := s.repo.GetMessageRecordByID(msgID); err != nil {
		return fmt.Errorf("会话记录不存在: %w", err)
	}
	return s.repo.DeleteMessageRecord(msgID)
}

// BatchOperateMessageRecords 批量操作会话记录
func (s *DigitalEmployeeService) BatchOperateMessageRecords(req *model.BatchMessageOperationRequest) error {
	switch req.Operation {
	case "delete":
		return s.repo.BatchDeleteMessageRecords(req.MsgIDs)
	case "mark_debug":
		return s.repo.BatchUpdateDebugStatus(req.MsgIDs, "1")
	case "unmark_debug":
		return s.repo.BatchUpdateDebugStatus(req.MsgIDs, "")
	default:
		return fmt.Errorf("不支持的操作: %s", req.Operation)
	}
}

// GetConversationGroups 获取会话分组
func (s *DigitalEmployeeService) GetConversationGroups(limit int) ([]*model.ConversationGroup, error) {
	groups, err := s.repo.GetConversationGroups(limit)
	if err != nil {
		return nil, err
	}

	// 为每个分组加载消息
	for _, group := range groups {
		messages, err := s.repo.GetMessagesByConversationID(group.ConversationID)
		if err != nil {
			continue
		}
		group.Messages = make([]model.MessageRecord, len(messages))
		for i, msg := range messages {
			group.Messages[i] = *msg
		}
	}

	return groups, nil
}

// GetMessageStatistics 获取会话统计
func (s *DigitalEmployeeService) GetMessageStatistics() (*model.MessageStatistics, error) {
	return s.repo.GetMessageStatistics()
}

// ========== 服务回传记录相关 ==========

// GetExportRecords 获取导出记录列表
func (s *DigitalEmployeeService) GetExportRecords(query *model.ExportRecordQuery) ([]*model.ExportRecord, int64, error) {
	return s.repo.GetExportRecords(query)
}

// GetExportRecordByID 根据ID获取导出记录
func (s *DigitalEmployeeService) GetExportRecordByID(id uint) (*model.ExportRecord, error) {
	return s.repo.GetExportRecordByID(id)
}

// CreateExportRecord 创建导出记录
func (s *DigitalEmployeeService) CreateExportRecord(req *model.ExportRecordRequest) (*model.ExportRecord, error) {
	record := &model.ExportRecord{
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		TotalRecords:    req.TotalRecords,
		FileName:        req.FileName,
		FilePath:        req.FilePath,
		EmailSubject:    req.EmailSubject,
		EmailRecipients: req.EmailRecipients,
		EmailCC:         req.EmailCC,
		EmailBCC:        req.EmailBCC,
		EmailStatus:     "pending",
	}

	if err := s.repo.CreateExportRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateExportRecord 更新导出记录
func (s *DigitalEmployeeService) UpdateExportRecord(id uint, req *model.ExportRecordRequest) (*model.ExportRecord, error) {
	existing, err := s.repo.GetExportRecordByID(id)
	if err != nil {
		return nil, fmt.Errorf("导出记录不存在: %w", err)
	}

	existing.StartDate = req.StartDate
	existing.EndDate = req.EndDate
	existing.TotalRecords = req.TotalRecords
	existing.FileName = req.FileName
	existing.FilePath = req.FilePath
	existing.EmailSubject = req.EmailSubject
	existing.EmailRecipients = req.EmailRecipients
	existing.EmailCC = req.EmailCC
	existing.EmailBCC = req.EmailBCC

	if err := s.repo.UpdateExportRecord(id, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

// DeleteExportRecord 删除导出记录
func (s *DigitalEmployeeService) DeleteExportRecord(id uint) error {
	if _, err := s.repo.GetExportRecordByID(id); err != nil {
		return fmt.Errorf("导出记录不存在: %w", err)
	}
	return s.repo.DeleteExportRecord(id)
}

// BatchOperateExportRecords 批量操作导出记录
func (s *DigitalEmployeeService) BatchOperateExportRecords(req *model.BatchExportOperationRequest) error {
	switch req.Operation {
	case "delete":
		return s.repo.BatchDeleteExportRecords(req.IDs)
	case "retry":
		// 重试发送邮件逻辑
		for _, id := range req.IDs {
			record, err := s.repo.GetExportRecordByID(id)
			if err != nil {
				continue
			}
			// 这里应该调用邮件发送服务
			// 暂时只更新状态为pending
			s.repo.UpdateEmailStatus(id, "pending", nil, "")
			_ = record
		}
		return nil
	default:
		return fmt.Errorf("不支持的操作: %s", req.Operation)
	}
}

// RetryExportEmail 重试发送邮件
func (s *DigitalEmployeeService) RetryExportEmail(id uint) error {
	record, err := s.repo.GetExportRecordByID(id)
	if err != nil {
		return fmt.Errorf("导出记录不存在: %w", err)
	}

	// 这里应该调用邮件发送服务
	// 暂时只更新状态
	_ = record
	return s.repo.UpdateEmailStatus(id, "pending", nil, "")
}

// GetExportStatistics 获取导出统计
func (s *DigitalEmployeeService) GetExportStatistics() (*model.ExportStatistics, error) {
	return s.repo.GetExportStatistics()
}
