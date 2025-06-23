package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "example.com/grpcbidichat/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Send messages in goroutine
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("👤 You: ")
			if !scanner.Scan() {
				break
			}
			text := scanner.Text()
			stream.Send(&pb.ChatMessage{Message: text})
		}
	}()

	// Receive messages
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Receive error: %v", err)
		}
		fmt.Printf("💬 Server: %s\n", in.Message)
		time.Sleep(100 * time.Millisecond)
	}
}
