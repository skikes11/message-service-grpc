package grpc_server

import (
	messagegrpc "message-service/pkg/v1/proto"
	"message-service/pkg/v1/usecase"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func MustMakeGrpcServerRun(
	listenAddress string,
	messageUsecase *usecase.MessageUseCase,
) {
	messageServer := NewGrpcMessageServer(messageUsecase)

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(10 * 1024 * 1024),
		grpc.MaxSendMsgSize(10 * 1024 * 1024),
	}
	server := grpc.NewServer(options...)
	healthServer := health.NewServer()
	healthServer.SetServingStatus(messagegrpc.MessageService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	messagegrpc.RegisterMessageServiceServer(server, messageServer)
	healthpb.RegisterHealthServer(server, healthServer)

	log.Info().Msgf("Starting grpc server at %s", listenAddress)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}

}
