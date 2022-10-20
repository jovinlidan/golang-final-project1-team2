package views

import (
	"time"
)

type Todo struct {
	ID   			uint 		`json:"id"`
	Title   		string 		`json:"title"`
	Description 	string 		`json:"description"`
	CompletedAt 	*time.Time 	`json:"completed_at"`
	
}
