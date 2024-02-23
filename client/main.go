package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	messagegrpc "message-service/pkg/v1/proto"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultMessage = "Xin chao moi nguoi"
)

var (
	message = flag.String("message", defaultMessage, "the message")
)

func main() {
	conn, err := grpc.Dial("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer conn.Close()

	client := messagegrpc.NewMessageServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := client.SendMessage(ctx, &messagegrpc.SendMessageResquest{Message: *message})

	if err != nil {
		log.Fatalf("err while send message: %v", err)
	}

	fmt.Println("", r.GetMessage())

	//req get messages list

	r1, err1 := client.GetMessageList(ctx, &messagegrpc.GetMessageListRequest{})

	if err1 != nil {
		log.Fatalf("err while send message: %v", err)
	}
	fmt.Println("", r1.GetMessages())

}
