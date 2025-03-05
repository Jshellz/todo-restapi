package models

import "time"

// Task структура самого таска с описание при запросе в json формате
type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
