package main

import (
	"context"
	"log"

	"github.com/jetexe/pbuf-example/api/messages"
	"github.com/jetexe/pbuf-example/api/services"
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
	defer conn.Close()

	client := services.NewMessagesAPIClient(conn)

	// send some
	account := "alice"
	for _, msg := range []string{"Hi", "What's up"} {
		// TODO err handling
		client.SendDirectMessages(context.Background(), &messages.Direct{
			Account: account,
			Text:    msg,
		})
	}

	// get some
	msgs, _ := client.GetDirectMessages(context.Background(), &services.Account{
		Account: account,
	})
	for _, msg := range msgs.Direct {
		log.Printf("%s", msg)
	}
}
