package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	pb "example.com/grpcbidichatUPGRADE/pkg/chat"
	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	streams map[string]pb.ChatService_ChatServer // tên người dùng → stream
	mu      sync.Mutex
}

func newChatServer() *chatServer {
	return &chatServer{
		streams: make(map[string]pb.ChatService_ChatServer),
	}
}

func (s *chatServer) Chat(stream pb.ChatService_ChatServer) error {
	var userName string

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			s.removeUser(userName)
			log.Printf("🔌 %s disconnected", userName)
			return nil
		}
		if err != nil {
			log.Printf("❌ Recv error: %v", err)
			s.removeUser(userName)
			return err
		}

		// Lần đầu: coi như đăng ký tên
		if userName == "" {
			userName = msg.GetFrom()
			s.mu.Lock()
			s.streams[userName] = stream
			s.mu.Unlock()
			log.Printf("✅ User connected: %s", userName)
			continue
		}

		log.Printf("💬 %s ➡ %s: %s", msg.GetFrom(), msg.GetTo(), msg.GetMessage())

		// Gửi lại cho người nhận
		s.mu.Lock()
		receiverStream, ok := s.streams[msg.GetTo()]
		s.mu.Unlock()

		if ok {
			// Gửi cho người nhận (nếu khác người gửi thì gửi lại cho người gửi)
			if msg.To != msg.From {
				_ = receiverStream.Send(msg) // Gửi cho người nhận
				_ = stream.Send(msg)         // Gửi lại cho người gửi để thấy tin đã gửi
			} else {
				_ = stream.Send(msg) // Gửi cho chính mình
			}
		} else {
			// Nếu người nhận không online thì vẫn gửi lại cho người gửi để biết
			_ = stream.Send(&pb.ChatMessage{
				From:    "Server",
				To:      msg.From,
				Message: fmt.Sprintf("⚠️ Người dùng '%s' chưa online hoặc không tồn tại.", msg.To),
			})
		}

	}
}

func (s *chatServer) removeUser(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.streams, name)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, newChatServer())

	log.Println("🚀 gRPC server running at :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
