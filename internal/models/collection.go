package models

import (
	"time"
)

type ResourceCollection struct {
	ID          ULID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
