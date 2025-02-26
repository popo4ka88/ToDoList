package models

import "time"

type ToDoModel struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Completed bool      `db:"completed" json:"completed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type ToDo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
