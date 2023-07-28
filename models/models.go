package models

import (
	"time"

	"github.com/google/uuid"
)

type General struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (s General) InitObject() {
	s.ID = uuid.New()
	now := time.Now().UTC()
	s.CreatedAt = &now
}
