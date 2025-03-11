package controllers

import (
	"net/http"
	"todolist/config"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách To-Do
func GetTodos(context *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, status FROM todos")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, todo)
	}

	context.JSON(http.StatusOK, todos)
}

// Thêm To-Do mới
func CreateTodo(context *gin.Context) {
	var todo models.Todo
	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", todo.Title, todo.Status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todo.ID = uint(id)

	context.JSON(http.StatusOK, todo)
}

// Cập nhật To-Do
func UpdateTodo(context *gin.Context) {
	var todo models.Todo
	id := context.Param("id")

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE todos SET title = ?, status = ? WHERE id = ?", todo.Title, todo.Status, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, todo)
}

// Xóa To-Do
func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	_, err := config.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "To-Do deleted"})
}
