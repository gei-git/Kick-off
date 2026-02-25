package model

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description,omitempty"`
	Priority    string     `gorm:"default:medium" json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Done        bool       `gorm:"default:false" json:"done"`
	UserID      uint       `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
