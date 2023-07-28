package models

import "time"

type General struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (s General) InitObject() {
	s.ID = ""
	now := time.Now().UTC()
	s.CreatedAt = &now
}
