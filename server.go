package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"zz/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s) //todo
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}
}
