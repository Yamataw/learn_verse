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

// Assure que CollectionRepo implémente l'interface générique
var _ Repository[models.ResourceCollection, ulid.ULID] = (*CollectionRepo)(nil)

// NewCollectionRepo crée une nouvelle instance de CollectionRepo
func NewCollectionRepo(db *sql.DB) *CollectionRepo {
	return &CollectionRepo{db: db}
}

// Create insère une nouvelle collection et retourne l'entité créée
func (r *CollectionRepo) Create(ctx context.Context, entity models.ResourceCollection) (models.ResourceCollection, error) {
	query := `
	INSERT INTO resource_collections (name, description)
	VALUES ($1, $2)
	RETURNING id, name, description, created_at, updated_at
	`
	var coll models.ResourceCollection
	err := r.db.QueryRowContext(ctx, query, entity.Name, entity.Description).
		Scan(&coll.ID, &coll.Name, &coll.Description, &coll.CreatedAt, &coll.UpdatedAt)
	if err != nil {
		return models.ResourceCollection{}, fmt.Errorf("failed to create collection: %w", err)
	}
	return coll, nil
}

// GetByID récupère une collection par son ULID
func (r *CollectionRepo) GetByID(ctx context.Context, id ulid.ULID) (models.ResourceCollection, error) {
	var coll models.ResourceCollection
	query := `
	SELECT id, name, description, created_at, updated_at
	FROM resource_collections
	WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&coll.ID, &coll.Name, &coll.Description, &coll.CreatedAt, &coll.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ResourceCollection{}, fmt.Errorf("collection not found: %w", err)
		}
		return models.ResourceCollection{}, err
	}
	return coll, nil
}

// List renvoie toutes les collections non supprimées
func (r *CollectionRepo) List(ctx context.Context) ([]models.ResourceCollection, error) {
	query := `
	SELECT id, name, description, created_at, updated_at
	FROM resource_collections
	WHERE deleted_at IS NULL
	ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.ResourceCollection
	for rows.Next() {
		var coll models.ResourceCollection
		if err := rows.Scan(
			&coll.ID,
			&coll.Name,
			&coll.Description,
			&coll.CreatedAt,
			&coll.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, coll)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// Update modifie une collection existante et retourne l'entité mise à jour
func (r *CollectionRepo) Update(ctx context.Context, entity models.ResourceCollection) (models.ResourceCollection, error) {
	query := `
	UPDATE resource_collections
	SET name = $1,
	    description = $2,
	    updated_at = now()
	WHERE id = $3 AND deleted_at IS NULL
	RETURNING id, name, description, created_at, updated_at
	`
	var coll models.ResourceCollection
	err := r.db.QueryRowContext(ctx, query, entity.Name, entity.Description, entity.ID).
		Scan(&coll.ID, &coll.Name, &coll.Description, &coll.CreatedAt, &coll.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ResourceCollection{}, fmt.Errorf("collection not found or deleted: %w", err)
		}
		return models.ResourceCollection{}, err
	}
	return coll, nil
}

// Delete marque une collection comme supprimée (soft delete)
func (r *CollectionRepo) Delete(ctx context.Context, id ulid.ULID) error {
	query := `
	UPDATE resource_collections
	SET deleted_at = now()
	WHERE id = $1 AND deleted_at IS NULL
	`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("collection not found or already deleted: %s", id.String())
	}
	return nil
}
