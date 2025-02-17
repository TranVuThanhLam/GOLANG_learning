package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

// Đọc danh sách Todo từ file JSON
func loadTodos() {
	file, err := os.ReadFile("todos.json") // Thay vì ioutil.ReadFile, dùng os.ReadFile
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	json.Unmarshal(file, &todos)
}

// Ghi danh sách Todo vào file JSON
func saveTodos() {
	file, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	err = os.WriteFile("todos.json", file, 0644) // Thay vì ioutil.WriteFile, dùng os.WriteFile
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// Lấy tất cả Todo
func getTodos(w http.ResponseWriter, r *http.Request) {
	loadTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Thêm Todo
func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	saveTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// Cập nhật Todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i] = updatedTodo
			saveTodos()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Todo not found"))
}

// Xóa Todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Todo deleted"))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Todo not found"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todos", addTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
