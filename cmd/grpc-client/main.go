package main

import (
	"context"
	"log"

	"github.com/jetexe/pbuf-example/internal/pkg/grpc/messages/v1"
	services "github.com/jetexe/pbuf-example/internal/pkg/grpc/services/v1"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close() //nolint:errcheck

	client := services.NewMessagesAPIServiceClient(conn)

	// send some
	account := "alice"

	for _, msg := range []string{"Hi", "What's up"} {
		_, err := client.SendDirectMessages(context.Background(), &services.SendDirectMessagesRequest{
			Message: &messages.Direct{
				Account: account,
				Text:    msg,
			},
		})

		if err != nil {
			log.Println("send message error", err)
		}
	}

	// get some
	msgs, _ := client.GetDirectMessages(context.Background(), &services.GetDirectMessagesRequest{
		Account: account,
	})
	for _, msg := range msgs.Direct {
		log.Printf("%s", msg)
	}
}
