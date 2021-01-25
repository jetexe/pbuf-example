package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jetexe/pbuf-example/internal/pkg/grpc/messages/v1"
	"github.com/jetexe/pbuf-example/internal/pkg/grpc/services/v1"
)

// Server structure contains all needed.
type Server struct {
	services.UnimplementedMessagesAPIServiceServer
	queue map[string][]string // [account]text
}

// NewServer instance.
func NewServer() *Server {
	return &Server{
		queue: make(map[string][]string),
	}
}

// GetDirectMessages handler.
func (s *Server) GetDirectMessages(
	ctx context.Context,
	in *services.GetDirectMessagesRequest,
) (*services.GetDirectMessagesResponse, error) {
	// TODO handle errors and nils and concurency
	msgs, ok := s.queue[in.Account]
	if !ok {
		return nil, nil
	}

	answer := &services.GetDirectMessagesResponse{
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

// SendDirectMessages handler.
func (s *Server) SendDirectMessages(
	ctx context.Context,
	in *services.SendDirectMessagesRequest,
) (*services.SendDirectMessagesResponse, error) {
	// TODO handle errors and nils and concurency
	s.queue[in.Message.Account] = append(s.queue[in.Message.Account], in.Message.Text)

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
	services.RegisterMessagesAPIServiceServer(grpcServer, NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
