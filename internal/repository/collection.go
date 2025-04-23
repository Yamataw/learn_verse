package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"learn_verse/internal/models"
)

// CollectionRepo offre l’accès SQL pour les collections
type CollectionRepo struct {
	db *sql.DB
}

func NewCollectionRepo(db *sql.DB) *CollectionRepo {
	return &CollectionRepo{db: db}
}

func (r *CollectionRepo) Create(ctx context.Context, name string, desc *string) (*models.ResourceCollection, error) {
	id := uuid.New()
	query := `INSERT INTO resource_collections (id,name,description) VALUES ($1,$2,$3)`
	if _, err := r.db.ExecContext(ctx, query, id, name, desc); err != nil {
		return nil, err
	}
	return &models.ResourceCollection{ID: id, Name: name, Description: desc}, nil
}
