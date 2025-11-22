package repository

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

func (gr *groupRepo) CreateChatHistory(gchistory model.GroupChatHistory) error {
	return gr.CreateGroupChatHistory(gchistory)
}
func (gr *groupRepo) GetChatHistories() ([]model.GroupChatHistory, error) {
	return gr.GetGroupChatHistories()
}
func (gr *groupRepo) FindChatHistory() ([]model.GroupChatHistory, error) {
	return gr.FindGroupChatHistory()
}
