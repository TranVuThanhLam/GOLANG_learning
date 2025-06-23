package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "example.com/grpcbidi/pkg/chat"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Println("📥 Nhận từ client:", msg.Message)

		reply := &pb.ChatMessage{Message: "Server phản hồi: " + msg.Message}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Không lắng nghe:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})

	fmt.Println("🚀 Server đang chạy tại :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
