package repository

import "github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"

func (cr *chatRepo) CreateChatHistory(chistory model.PrivateChatHistory) error {
	chistory.Status = model.Sent
	err := cr.ChatDao.CreateChatHistory(chistory)
	if err != nil {
		return err
	}
	return nil
}
func (cr *chatRepo) GetChatHistories(senderID, receiverID int64, limit, offset int) ([]model.PrivateChatHistory, error) {
	histories, err := cr.ChatDao.GetChatHistories(senderID, receiverID, limit, offset)
	if err != nil {
		return nil, err
	}
	return histories, nil
}
func (cr *chatRepo) FindChatHistory(senderID, receiverID int64, limit, offset int, keyword string) ([]model.PrivateChatHistory, error) {
	histories, err := cr.ChatDao.FindChatHistory(senderID, receiverID, limit, offset, keyword)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (cr *chatRepo) RevokeChatHistory(id int64) error {
	err := cr.ChatDao.RevokeChatHistory(id)
	if err != nil {
		return err
	}
	return nil
}
