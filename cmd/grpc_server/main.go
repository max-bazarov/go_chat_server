package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/max-bazarov/go_chat_server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	chat_v1.UnimplementedChatAPIV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("%#v", req.GetChat())
	return &chat_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Received message: %s", req.GetMessage())
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
