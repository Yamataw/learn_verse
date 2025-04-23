package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/oklog/ulid/v2"

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
	query := `INSERT INTO resource_collections (name,description) VALUES ($1,$2)`
	if _, err := r.db.ExecContext(ctx, query, name, desc); err != nil {
		return nil, err
	}
	return &models.ResourceCollection{Name: name, Description: desc}, nil
}

func (r *CollectionRepo) FindByID(ctx context.Context, id ulid.ULID) (*models.ResourceCollection, error) {
	var coll models.ResourceCollection
	q := `SELECT id,name,description,created_at,updated_at FROM resource_collections WHERE id=$1`
	if err := r.db.QueryRowContext(ctx, q, id).Scan(
		&coll.ID, &coll.Name, &coll.Description, &coll.CreatedAt, &coll.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &coll, nil
}

// List renvoie toutes les collections
func (r *CollectionRepo) List(ctx context.Context) ([]models.ResourceCollection, error) {
	q := `SELECT id,name,description,created_at,updated_at FROM resource_collections`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []models.ResourceCollection
	for rows.Next() {
		var (
			idStr string
			coll  models.ResourceCollection
		)
		if err := rows.Scan(
			&idStr,
			&coll.Name,
			&coll.Description,
			&coll.CreatedAt,
			&coll.UpdatedAt,
		); err != nil {
			return nil, err
		}
		u, err := ulid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("invalid ULID %q: %w", idStr, err)
		}
		coll.ID = u
		out = append(out, coll)
	}

	return out, nil
}
