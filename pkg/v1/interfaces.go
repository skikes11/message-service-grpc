package v1

import (
	"context"
	"message-service/internal/models"

	"github.com/google/uuid"
)

type MessageInterface interface {
	Create(context context.Context, message *models.Message) error
	GetById(context context.Context, id uuid.UUID) (*models.Message, error)
	GetList(context context.Context) ([]*models.Message, error)
}
