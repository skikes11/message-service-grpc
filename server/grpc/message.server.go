package grpc_server

import (
	"context"
	"message-service/internal/models"
	messagegrpc "message-service/pkg/v1/proto"
	"message-service/pkg/v1/usecase"
)

type MessageServer struct {
	messageUsecase *usecase.MessageUseCase
	messagegrpc.UnimplementedMessageServiceServer
}

func NewGrpcMessageServer(messageUsecase *usecase.MessageUseCase) *MessageServer {
	return &MessageServer{
		messageUsecase: messageUsecase,
	}
}

func (s *MessageServer) SendMessage(ctx context.Context, req *messagegrpc.SendMessageResquest) (*messagegrpc.SendMessageResponse, error) {
	_, err := s.messageUsecase.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	return &messagegrpc.SendMessageResponse{
		Message: "Send message successfully",
	}, nil
}

func (s *MessageServer) GetMessageList(ctx context.Context, req *messagegrpc.GetMessageListRequest) (*messagegrpc.GetMessageListResponse, error) {
	messages, err := s.messageUsecase.GetList(ctx)

	if err != nil {
		return nil, err
	}

	return &messagegrpc.GetMessageListResponse{
		Messages: models.ConvertToRpcListMessage(messages),
	}, nil
}
