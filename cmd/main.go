package main

import (
	"ToDoList/internal/server"
	"ToDoList/pkg/database"
	_ "github.com/lib/pq"
)

func main() {
	router := server.InitRouter()
	database.Connect()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
