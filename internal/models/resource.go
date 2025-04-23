package models

import (
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"time"
)

type Resource struct {
	ID           ulid.ULID
	CollectionID *ulid.ULID
	Type         string // "note" | "flashcard" | ...
	Title        string
	Content      json.RawMessage
	Metadata     json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
