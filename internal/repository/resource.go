package repository

import (
	"context"
	"database/sql"

	"learn_verse/internal/models"
)

// ResourceRepo offre l’accès SQL pour les resources
type ResourceRepo struct {
	db *sql.DB
}

func NewResourceRepo(db *sql.DB) *ResourceRepo {
	return &ResourceRepo{db: db}
}

func (r *ResourceRepo) Create(ctx context.Context, res *models.Resource) (*models.Resource, error) {
	query := `INSERT INTO resources (collection_id,type,title,content,metadata) VALUES ($1,$2,$3,$4,$5)`
	if _, err := r.db.ExecContext(ctx, query, res.CollectionID, res.Type, res.Title, res.Content, maybeNull(res.Metadata)); err != nil {
		return nil, err
	}
	return res, nil
}

func maybeNull(b []byte) interface{} {
	if len(b) == 0 {
		return nil
	}
	return b
}
