package dao

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
)

func (cd *chatDao) CreateChatHistory(chistory model.PrivateChatHistory) error {
	return cd.DB.Create(chistory).Error
}

func (cd *chatDao) GetChatHistories(senderID, receiverID int64, limit, offset int) ([]model.PrivateChatHistory, error) {
	var histories []model.PrivateChatHistory
	err := cd.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		senderID, receiverID, receiverID, senderID).
		Order("timestamp").
		Limit(limit).
		Offset(offset).
		Find(&histories).Error
	return histories, err
}

func (cd *chatDao) FindChatHistory(senderID, receiverID int64, limit, offset int, keyword string) ([]model.PrivateChatHistory, error) {
	var histories []model.PrivateChatHistory
	err := cd.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		senderID, receiverID, receiverID, senderID).
		Where("content LIKE ?", "%"+keyword+"%").
		Order("timestamp").
		Limit(limit).
		Offset(offset).
		Find(&histories).Error
	return histories, err
}

func (cd *chatDao) RevokeChatHistory(id int64) error {
	return cd.DB.Model(&model.PrivateChatHistory{}).
		Where("id = ?", id).
		Update("status", model.Revoked).Error
}
