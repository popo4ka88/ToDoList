package handlers

import (
	"ToDoList/models"
	"ToDoList/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateTodo(c *gin.Context) {
	var t models.ToDo

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is required",
		})
		return
	}

	tm := models.ToDoModel{
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	var id int
	err := database.DB.QueryRow(`
		INSERT INTO todos (title, completed, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, tm.Title, tm.Completed, tm.CreatedAt).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save todo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"todo_id": id,
	})
}

func FetchTodos(c *gin.Context) {
	var todos []models.ToDoModel

	rows, err := database.DB.Query("SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch todos",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.ToDoModel
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to scan todo",
				"error":   err.Error(),
			})
			return
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error after scanning todos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	var t models.ToDo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is required",
		})
		return
	}

	_, err = database.DB.Exec(`
		UPDATE todos
		SET title = $1, completed = $2
		WHERE id = $3
	`, t.Title, t.Completed, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update todo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully",
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	_, err = database.DB.Exec(`DELETE FROM todos WHERE id = $1`, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete todo",
			"error":   err.Error(),
		})
		return
	}

}
