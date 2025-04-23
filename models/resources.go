package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Resource struct {
	ID           uuid.UUID
	CollectionID *uuid.UUID
	Type         string // "note" | "flashcard" | ...
	Title        string
	Content      json.RawMessage
	Metadata     json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
