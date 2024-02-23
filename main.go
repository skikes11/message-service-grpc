package main

import (
	"fmt"
	"log"
	"message-service/initializers"
	"message-service/internal/database"
	"message-service/pkg/v1/repository"
	"message-service/pkg/v1/usecase"
	grpc_server "message-service/server/grpc"

	"gorm.io/gorm"
)

var (
	config initializers.Config
	DB     *gorm.DB

	messageUseCase *usecase.MessageUseCase
)

func init() {
	var err error
	config, err = initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	DB = database.ConnectDB(&config)

	messageRepo := repository.NewMessageRepo(DB, true)

	messageUseCase = usecase.NewMessageUseCase(DB, messageRepo)

}

func main() {

	grpc_server.MustMakeGrpcServerRun(fmt.Sprintf(":%s", config.RPCPort), messageUseCase)

}
