package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"

	pb "example.com/grpcbidichatUPGRADE/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	stream         pb.ChatService_ChatClient
	streamInit     sync.Once
	streamReady    = make(chan struct{})
	messageHistory []pb.ChatMessage
	historyMutex   sync.Mutex
	userName       string
)

func main() {
	http.HandleFunc("/", requireLogin(chatPageHandler))
	http.HandleFunc("/login", loginPageHandler)
	http.HandleFunc("/send", requireLogin(sendHandler))
	http.HandleFunc("/messages", requireLogin(messagesHandler))

	fmt.Println("🌐 HTTP server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Middleware: redirect nếu chưa nhập tên
func requireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userName = cookie.Value
		streamInit.Do(func() {
			go connectToGRPC(userName)
		})
		<-streamReady // đảm bảo stream đã sẵn sàng
		next(w, r)
	}
}

// Kết nối tới gRPC stream
func connectToGRPC(name string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	client := pb.NewChatServiceClient(conn)

	stream, err = client.Chat(context.Background())
	if err != nil {
		log.Fatalf("❌ Failed to create stream: %v", err)
	}

	// Gửi tên người dùng cho server
	_ = stream.Send(&pb.ChatMessage{From: name})

	// Bắt đầu nhận message
	go receiveMessages()

	close(streamReady)
}

func receiveMessages() {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("🔚 Server closed stream")
			return
		}
		if err != nil {
			log.Printf("❌ Receive error: %v", err)
			continue
		}
		if msg.From == userName || msg.To == userName {
			historyMutex.Lock()
			messageHistory = append(messageHistory, *msg)
			historyMutex.Unlock()
		}
	}
}

func chatPageHandler(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<!DOCTYPE html>
	<html><head><title>Chat</title></head><body>
	<h2>Chào {{.Name}} 👋</h2>
	<form action="/send" method="post">
		Gửi đến: <input name="to"><br>
		Nội dung: <textarea name="message"></textarea><br>
		<input type="submit" value="Gửi">
	</form>
	<hr>
	<div id="messages">Đang tải tin nhắn...</div>
	<script>
		function load() {
			fetch("/messages")
			.then(r => r.text())
			.then(html => {
				document.getElementById("messages").innerHTML = html;
			});
		}
		setInterval(load, 2000);
		load();
	</script>
	</body></html>
	`
	t := template.Must(template.New("chat").Parse(tpl))
	t.Execute(w, struct{ Name string }{userName})
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("username")
		if name != "" {
			http.SetCookie(w, &http.Cookie{
				Name:  "username",
				Value: name,
				Path:  "/",
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	tpl := `
	<!DOCTYPE html>
	<html>
	<head><title>Đăng nhập</title></head>
	<body>
		<h2>Đăng nhập</h2>
		<form method="POST">
			Tên bạn: <input name="username">
			<input type="submit" value="Vào chat">
		</form>
	</body>
	</html>
	`
	t := template.Must(template.New("login").Parse(tpl))
	t.Execute(w, nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	to := r.FormValue("to")
	msg := r.FormValue("message")
	if to == "" || msg == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := stream.Send(&pb.ChatMessage{From: userName, To: to, Message: msg})
	if err != nil {
		http.Error(w, "❌ Send failed", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	historyMutex.Lock()
	defer historyMutex.Unlock()
	for _, msg := range messageHistory {
		fmt.Fprintf(w, "<p><strong>%s ➡ %s:</strong> %s</p>", msg.From, msg.To, msg.Message)
	}
}
