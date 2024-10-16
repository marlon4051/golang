package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=5"`
	Status      string `json:"status" validate:"required,oneof=pending working completed"`
	UserID      int    `json:"user_id"`
}
