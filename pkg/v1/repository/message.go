package repository

import (
	"context"
	"message-service/internal/models"
	interfaces "message-service/pkg/v1"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB, init bool) interfaces.MessageInterface {
	if init {
		if err := db.AutoMigrate(&models.Message{}); err != nil {
			panic(err)
		}
	}
	return &MessageRepo{db: db}
}

func (d *MessageRepo) Create(ctx context.Context, message *models.Message) error {
	return d.db.WithContext(ctx).Create(message).Error
}

func (d *MessageRepo) GetById(ctx context.Context, messageId uuid.UUID) (*models.Message, error) {
	var message *models.Message
	err := d.db.WithContext(ctx).
		First(&message, messageId).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (d *MessageRepo) GetList(ctx context.Context) ([]*models.Message, error) { //limit int16, offset int16

	var messages []*models.Message

	err := d.db.WithContext(ctx).
		// Limit(int(limit)).
		// Offset(int(offset)).
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}
