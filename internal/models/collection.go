package models

import (
	"github.com/google/uuid"
	"time"
)

type ResourceCollection struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
}
