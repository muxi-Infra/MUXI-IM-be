package model

import (
	"encoding/json"
	"time"
)

// 消息状态
const (
	Sent = iota
	Read
	Revoked
)

type PrivateChatHistory struct {
	Id          int64
	SenderUid   int64
	ReceiverUid int64
	QuoteReply  int64
	Content     json.RawMessage
	Extensions  json.RawMessage // 可能是额外的元数据
	Timestamp   int64
	Status      int // 消息状态

	CreatedAt time.Time
	UpdatedAt time.Time
}


type MessageContent struct {
	Kind    int // At, PlainText, Image
	Content string
}

func NewContentFromSlice(contents []MessageContent) (json.RawMessage, error) {
	b, err := json.Marshal(contents)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b), nil
}

func NewExtensionsFromMap(m map[string]string) (json.RawMessage, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b), nil
}

func ParseContent(b json.RawMessage) ([]MessageContent, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var out []MessageContent
	err := json.Unmarshal(b, &out)
	return out, err
}

func ParseExtensions(b json.RawMessage) (map[string]string, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var out map[string]string
	err := json.Unmarshal(b, &out)
	return out, err
}
