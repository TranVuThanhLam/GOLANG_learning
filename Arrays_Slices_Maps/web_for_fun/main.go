package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// ToDo represents a single to-do item
type ToDo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos []ToDo

// Home handler for rendering the To-Do list
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Todo List</title>
    <!-- Thêm Bootstrap CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
        }
        .todo-container {
            max-width: 600px;
            margin: 50px auto;
            padding: 20px;
            background-color: white;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }
        .completed {
            text-decoration: line-through;
        }
    </style>
</head>
<body>

<div class="todo-container">
    <h1 class="text-center">My Todo List</h1>

    <ul class="list-group">
        {{range .}}
        <li class="list-group-item d-flex justify-content-between align-items-center {{if .Done}}completed{{end}}">
            <span>{{.Task}}</span>

            <div class="btn-group">
                <form action="/toggle/{{.ID}}" method="post" style="display:inline;">
                    <button type="submit" class="btn btn-sm {{if .Done}}btn-warning{{else}}btn-success{{end}}">
                        {{if .Done}}Undo{{else}}Done{{end}}
                    </button>
                </form>
                <form action="/delete/{{.ID}}" method="post" style="display:inline;">
                    <button type="submit" class="btn btn-sm btn-danger">Delete</button>
                </form>
            </div>
        </li>
        {{end}}
    </ul>

    <form action="/add" method="post" class="mt-3">
        <div class="input-group">
            <input type="text" name="task" class="form-control" required placeholder="New Task">
            <button type="submit" class="btn btn-primary">Add</button>
        </div>
    </form>
</div>

<!-- Thêm Bootstrap JS và Popper.js -->
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.6/dist/umd/popper.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.min.js"></script>

</body>
</html>


	`)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	// Render the template with the ToDo list data
	tmpl.Execute(w, todos)
}

// Add handler for adding a new To-Do item
func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.FormValue("task")

	// Create a new to-do item
	id := fmt.Sprintf("%d", len(todos)+1)
	todo := ToDo{
		ID:   id,
		Task: task,
		Done: false,
	}

	// Add the new to-do item to the list
	todos = append(todos, todo)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Toggle handler for marking a To-Do item as done/undone
func Toggle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find the To-Do item by ID and toggle its done status
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = !todos[i].Done
			break
		}
	}

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete handler for deleting a To-Do item
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Remove the To-Do item by ID
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Add routes
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/add", Add).Methods("POST")
	r.HandleFunc("/toggle/{id}", Toggle).Methods("POST")
	r.HandleFunc("/delete/{id}", Delete).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", r)
}
