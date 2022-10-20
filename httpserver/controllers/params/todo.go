package params

import "time"

type TodoCreateRequest struct {
	Title string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
}

