package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Grumlebob/ThomasShowcase/protos"
)

var userId int32

func main() {
	// Create a virtual RPC Client Connection on port 9080
	var conn *grpc.ClientConn
	context, cancelFunction := context.WithTimeout(context.Background(), time.Second*200) //standard er 5
	defer cancelFunction()
	// IPv4:port = "172.30.48.1:9080"
	conn, err := grpc.DialContext(context, ":9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	client := pb.NewChatServiceClient(conn)

	sendMsg(client, context)
}

func sendMsg(client pb.ChatServiceClient, context context.Context) {
	msg := "hello boy"
	chatMessage := &pb.ChatMessage{Message: msg}

	user, err := client.PublishMessage(context, chatMessage)
	if err != nil {
		log.Fatalf("Error when calling server: %s", err)
	}
	fmt.Println("Response from server: ", user.Message)

}
