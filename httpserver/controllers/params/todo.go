package params

import "time"

type TodoCreateRequest struct {
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
}

type TodoUpdateRequest struct {
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
}

type TodoGetByIDRequest struct {
	ID int `validate:"required" json:"id"`
}

type TodoDeleteRequest struct {
	ID          int        `validate:"required" json:"id"`
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
}
