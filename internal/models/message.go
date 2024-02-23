package models

import (
	messagegrpc "message-service/pkg/v1/proto"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;"`
}

func ConvertToRpcMessage(m *Message) *messagegrpc.Message {
	createdAt := timestamppb.New(m.CreatedAt)
	return &messagegrpc.Message{
		Id:        int32(m.Id[1]),
		Message:   m.Message,
		CreatedAt: createdAt,
	}
}

func ConvertToRpcListMessage(listMessage []*Message) []*messagegrpc.Message {
	var result []*messagegrpc.Message
	for _, m := range listMessage {
		result = append(result, ConvertToRpcMessage(m))
	}
	return result
}
