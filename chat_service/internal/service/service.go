package service

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/repository"

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(chatRepo repository.ChatRepository) ChatService {
	return &chatService{chatRepo}
}
