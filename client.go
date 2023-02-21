package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"zz/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:%s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hello from client!",
	}
	responce, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Say Hello: %s", err)
	}
	log.Printf("Response from server:%s", responce.Body)
}
