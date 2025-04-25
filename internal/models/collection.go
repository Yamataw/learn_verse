package models

import (
	"time"
)

type ResourceCollection struct {
	ID          ULID      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// si vous voulez pouvoir renvoyer null au lieu de "0001-01-01T00:00:00Z"
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
