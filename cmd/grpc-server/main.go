package main

import (
	"context"
	"fmt"
	"log"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"github.com/jetexe/pbuf-example/api/messages"
	"github.com/jetexe/pbuf-example/api/services"
)

type Server struct {
	queue map[string][]string // [account]text
	services.UnimplementedMessagesAPIServer
}

func NewServer() *Server {
	return &Server{
		queue: make(map[string][]string),
	}
}

func (s *Server) GetDirectMessages(ctx context.Context, in *services.Account) (*services.Directs, error) {
	// TODO handle errors and nils and concurency
	msgs, ok := s.queue[in.Account]
	if !ok {
		return nil, nil
	}
	answer := &services.Directs{
		Direct: make([]*messages.Direct, len(msgs)),
	}
	for i, message := range msgs {
		answer.Direct[i] = &messages.Direct{
			Account: in.Account, // TODO must be from
			Text:    message,
		}
	}

	return answer, nil
}

func (s *Server) SendDirectMessages(ctx context.Context, in *messages.Direct) (*empty.Empty, error) {
	// TODO handle errors and nils and concurency
	s.queue[in.Account] = append(s.queue[in.Account], in.Text)

	return nil, nil
}

// TODO streaming version

const port = 8080

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	services.RegisterMessagesAPIServer(grpcServer, NewServer())
	grpcServer.Serve(lis)
}
