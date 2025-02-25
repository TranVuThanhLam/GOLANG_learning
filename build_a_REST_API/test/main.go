package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler started")
	// Dữ liệu để truyền vào template
	data := struct {
		Name  string
		Title string
	}{
		Name:  "John", // Bạn có thể thay đổi tên ở đây
		Title: "Welcome Page",
	}

	// Tải template từ tệp HTML
	tmpl, err := template.ParseFiles("layout.tmpl", "greeting.tmpl")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	// Thực thi template và gửi kết quả HTML vào trình duyệt
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		log.Printf("Data: %+v\n", data)
	}
	log.Println("Handler finished")
}

func main() {
	// Đăng ký handler cho URL root "/"
	http.HandleFunc("/", handler)

	// Khởi động server
	log.Println("Server is starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
