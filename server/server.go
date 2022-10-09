package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Grumlebob/ThomasShowcase/protos"

	"google.golang.org/grpc"
)

type Server struct {
	pb.ChatServiceServer
	messageChannels map[int32]chan *pb.ChatMessage
}

func main() {
	// Create listener tcp on port 9080
	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &Server{
		messageChannels: make(map[int32]chan *pb.ChatMessage),
	})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (s *Server) PublishMessage(ctx context.Context, clientMessage *pb.ChatMessage) (*pb.ChatMessage, error) {

	fmt.Println("Server got message from client with msg, ", clientMessage.Message, "and time:", time.Now())
	return &pb.ChatMessage{Message: "server greets back:"}, nil

}
