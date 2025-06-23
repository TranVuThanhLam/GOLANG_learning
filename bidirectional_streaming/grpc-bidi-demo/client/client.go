package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "example.com/grpcbidi/pkg/chat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Không kết nối được:", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatal("Lỗi khi tạo stream:", err)
	}

	// Nhận tin nhắn từ server song song
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Println("Lỗi nhận:", err)
				return
			}
			fmt.Println("📨", res.Message)
		}
	}()

	// Gửi tin nhắn
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Bạn: ")
		text, _ := reader.ReadString('\n')
		stream.Send(&pb.ChatMessage{Message: text})
		time.Sleep(200 * time.Millisecond)
	}
}
