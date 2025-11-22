package grpc

import (
	context "context"

	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/model"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/service"
)

type ChatRpcServer struct {
	Chat service.ChatService
	UnimplementedChatServiceServer
}

func NewChatRpcServer(ser service.ChatService) *ChatRpcServer {
	return &ChatRpcServer{Chat: ser}
}

func (s *ChatRpcServer) Register(server *kgrpc.Server) {
	RegisterChatServiceServer(server, s)
}

func (s *ChatRpcServer) CreateChatHistory(ctx context.Context, req *SendPrivateMessageRequest) (*SendPrivateMessageResponse, error) {
	content := req.Message.Contents
	con := make([]model.MessageContent, 0, len(content))
	for _, c := range content {
		con = append(con, model.MessageContent{Kind: int(c.Kind), Content: c.Content})
	}

	cb, err := model.NewContentFromSlice(con)
	if err != nil {
		return nil, err
	}
	eb, err := model.NewExtensionsFromMap(req.Message.Extensions)
	if err != nil {
		return nil, err
	}

	message := model.PrivateChatHistory{
		SenderUid:   req.Message.SenderUid,
		ReceiverUid: req.Message.ReceiverUid,
		QuoteReply:  req.Message.QuoteReply,
		Content:     cb,
		Extensions:  eb,
	}
	if err := s.Chat.Create(message); err != nil {
		return nil, err
	}
	return &SendPrivateMessageResponse{
		Id: message.Id,
	}, nil
}

func (s *ChatRpcServer) GetChatHistory(ctx context.Context, req *GetPrivateChatHistoryRequest) (*GetPrivateChatHistoryResponse, error) {
	histories, err := s.Chat.Get(req.SenderUid, req.ReceiverUid, -1, -1)
	if err != nil {
		return nil, err
	}
	var msgs []*PrivateChatMessage
	for _, history := range histories {
		cs, err := model.ParseContent(history.Content)
	if err != nil {
		return nil, err
	}
		content := make([]*MessageContent, 0, len(cs))
	for _, c := range cs {
		content = append(content, &MessageContent{Kind: int32(c.Kind), Content: c.Content})
	}

		ext, err := model.ParseExtensions(history.Extensions)
	if err != nil {
		return nil, err
	}
		msg := &PrivateChatMessage{
			Id:          history.Id,
			SenderUid:   history.SenderUid,
			ReceiverUid: history.ReceiverUid,
			QuoteReply:  history.QuoteReply,
			Timestamp:   history.Timestamp,
			Contents:    content,
			Status:      int32(history.Status),
			Extensions:  ext,
		}
		msgs = append(msgs, msg)
	}
	return &GetPrivateChatHistoryResponse{
		Messages: msgs,
	}, nil
}

func (s *ChatRpcServer) FindChatHistory(ctx context.Context, req *SearchPrivateMessagesRequest) (*SearchPrivateMessagesResponse, error) {
	histories, err := s.Chat.Get(req.SenderUid, req.ReceiverUid, -1, -1)
	if err != nil {
		return nil, err
	}
	var msgs []*PrivateChatMessage
	for _, history := range histories {
		cs, err := model.ParseContent(history.Content)
	if err != nil {
		return nil, err
	}
		content := make([]*MessageContent, 0, len(cs))
	for _, c := range cs {
		content = append(content, &MessageContent{Kind: int32(c.Kind), Content: c.Content})
	}

		ext, err := model.ParseExtensions(history.Extensions)
	if err != nil {
		return nil, err
	}
		msg := &PrivateChatMessage{
			Id:          history.Id,
			SenderUid:   history.SenderUid,
			ReceiverUid: history.ReceiverUid,
			QuoteReply:  history.QuoteReply,
			Timestamp:   history.Timestamp,
			Contents:    content,
			Status:      int32(history.Status),
			Extensions:  ext,
		}
		msgs = append(msgs, msg)
	}
	return &SearchPrivateMessagesResponse{
		Messages: msgs,
	}, nil
}

func (s *ChatRpcServer) RevokeChatHistory(ctx context.Context, req *RevokePrivateMessageRequest) (*RevokePrivateMessageResponse, error) {
	err := s.Chat.Revoke(req.Id)
	if err != nil {
		return nil, err
	}
	return &RevokePrivateMessageResponse{}, nil
}
