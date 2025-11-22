package service

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
)

const (
	DefaultLimit  = 20
	DefaultOffset = 0
)

func (cs *chatService) Create(history model.PrivateChatHistory) error {
	return cs.repo.CreateChatHistory(history)
}

func (cs *chatService) Get(senderID, receiverID int64, limit, offset int) ([]model.PrivateChatHistory, error) {
	if limit <= 0 {
		limit = DefaultLimit
	}
	if offset <= 0 {
		offset = DefaultOffset
	}
	return cs.repo.GetChatHistories(senderID, receiverID, limit, offset)
}

func (cs *chatService) Find(senderID, receiverID int64, limit, offset int, keyword string) ([]model.PrivateChatHistory, error) {
	if limit <= 0 {
		limit = DefaultLimit
	}
	if offset <= 0 {
		offset = DefaultOffset
	}
	return cs.repo.FindChatHistory(senderID, receiverID, limit, offset, keyword)
}

func (cs *chatService) Revoke(id int64) error {
	return cs.repo.RevokeChatHistory(id)
}
