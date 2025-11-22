package repository

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/repository/dao"
)

// 私聊相关接口
type ChatRepository interface {
	CreateChatHistory(chistory model.PrivateChatHistory) error

	GetChatHistories(senderID, receiverID int64, limit, offset int) ([]model.PrivateChatHistory, error)
	// 主动搜索聊天历史记录
	FindChatHistory(senderID, receiverID int64, limit, offset int, keyword string) ([]model.PrivateChatHistory, error)

	RevokeChatHistory(id int64) error
}

// 群聊相关接口
type GroupChatRepository interface {
	CreateChatHistory(gchistory model.GroupChatHistory) error
	GetChatHistories() ([]model.GroupChatHistory, error)
	// 主动搜索聊天历史记录
	FindChatHistory() ([]model.GroupChatHistory, error)
}

type chatRepo struct {
	dao.ChatDao
}

type groupRepo struct {
	dao.GroupChatDao
}

func NewChatRepository(dao dao.ChatDao) ChatRepository {
	return &chatRepo{dao}
}

func NewGroupChatRepository() GroupChatRepository {
	return &groupRepo{}
}
