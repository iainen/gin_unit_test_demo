package entity

import "time"

// Model base model
type Model struct {
	ID        string     `mysql:"column:id;primary_key;size:36;" json:"id"`
	CreatedAt time.Time  `mysql:"column:created_at;index;" json:"-"`
	UpdatedAt time.Time  `mysql:"column:updated_at;index;" json:"-"`
	DeletedAt *time.Time `mysql:"column:deleted_at;index;" json:"-"`
}
