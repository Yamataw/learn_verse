package models

import (
	"encoding/json"
	"time"
)

type Resource struct {
	ID           ULID
	CollectionID *ULID
	Type         string // "note" | "flashcard" | ...
	Title        string
	Content      json.RawMessage
	Metadata     json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
