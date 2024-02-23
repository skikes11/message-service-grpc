package usecase

import (
	"context"
	"message-service/internal/models"
	interfaces "message-service/pkg/v1"
	messagegrpc "message-service/pkg/v1/proto"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type MessageUseCase struct {
	db *gorm.DB

	message interfaces.MessageInterface
}

func NewMessageUseCase(db *gorm.DB, message interfaces.MessageInterface) *MessageUseCase {
	return &MessageUseCase{
		db:      db,
		message: message,
	}
}

func (m *MessageUseCase) Create(ctx context.Context, req *messagegrpc.SendMessageResquest) (*models.Message, error) {
	message := &models.Message{
		Id:        uuid.New(),
		Message:   req.GetMessage(),
		CreatedAt: time.Now(),
	}

	if err := m.message.Create(ctx, message); err != nil {
		return nil, status.Newf(codes.Internal, "failed to create dataset: %v", err).Err()
	}

	return message, nil
}

func (m *MessageUseCase) GetById(ctx context.Context, messageId uuid.UUID) (*models.Message, error) {
	return m.message.GetById(ctx, messageId)
}

func (m *MessageUseCase) GetList(ctx context.Context) ([]*models.Message, error) {
	return m.message.GetList(ctx)
}
