package service

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

type ChatService interface {
	Create(history model.PrivateChatHistory) error
	Get(senderID, receiverID int64, limit, offset int) ([]model.PrivateChatHistory, error)
	Find(senderID, receiverID int64, limit, offset int, keyword string) ([]model.PrivateChatHistory, error)
	Revoke(id int64) error
	// 消息状态
	// MarkMessageAsRead()
	// GetUnreadCount()
	//
	// UpdateUserStatus()
	// GetUserStatus()
	//
	// SubscribeMessages()
}
