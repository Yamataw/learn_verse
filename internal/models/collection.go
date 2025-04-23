package models

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type ResourceCollection struct {
	ID          ulid.ULID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
