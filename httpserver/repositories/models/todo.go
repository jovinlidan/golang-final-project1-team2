package models

import "time"

type Todo struct {
	ID       	uint `json:"id" gorm:"primary_key;autoIncrement"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
}
