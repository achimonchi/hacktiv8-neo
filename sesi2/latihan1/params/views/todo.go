package views

import "time"

type CreateTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoSuccessSwag struct {
	Status  int               `json:"status" example:"201"`
	Message string            `json:"message" example:"CREATE TODO SUCCESS"`
	Payload CreateTodoPayload `json:"payload"`
}
