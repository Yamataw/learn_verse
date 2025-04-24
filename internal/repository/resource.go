package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learn_verse/internal/models"
)

// ResourceRepo offre l’accès SQL pour les resources
// et implémente Repository[*models.Resource, int64]
type ResourceRepo struct {
	db *sql.DB
}

func NewResourceRepo(db *sql.DB) *ResourceRepo {
	return &ResourceRepo{db: db}
}

// Assure qu'on implémente bien l'interface :
var _ Repository[*models.Resource, int64] = (*ResourceRepo)(nil)

func (r *ResourceRepo) Create(ctx context.Context, entity *models.Resource) (*models.Resource, error) {
	query := `
		INSERT INTO resources (collection_id, type, title, content, metadata)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRowContext(
		ctx, query,
		entity.CollectionID,
		entity.Type,
		entity.Title,
		entity.Content,
		maybeNull(entity.Metadata),
	).Scan(&entity.ID)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *ResourceRepo) GetByID(ctx context.Context, id int64) (*models.Resource, error) {
	query := `
		SELECT id, collection_id, type, title, content, metadata, created_at, updated_at
		FROM resources
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var res models.Resource
	var metadata sql.NullString
	if err := row.Scan(
		&res.ID,
		&res.CollectionID,
		&res.Type,
		&res.Title,
		&res.Content,
		&metadata,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("resource %d not found", id)
		}
		return nil, err
	}
	if metadata.Valid {
		res.Metadata = []byte(metadata.String)
	}
	return &res, nil
}

func (r *ResourceRepo) List(ctx context.Context) ([]*models.Resource, error) {
	query := `
		SELECT id, collection_id, type, title, content, metadata, created_at, updated_at
		FROM resources
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*models.Resource
	for rows.Next() {
		var res models.Resource
		var metadata sql.NullString
		if err := rows.Scan(
			&res.ID,
			&res.CollectionID,
			&res.Type,
			&res.Title,
			&res.Content,
			&metadata,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if metadata.Valid {
			res.Metadata = []byte(metadata.String)
		}
		list = append(list, &res)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ResourceRepo) Update(ctx context.Context, entity *models.Resource) (*models.Resource, error) {
	query := `
		UPDATE resources
		SET collection_id = $1, type = $2, title = $3, content = $4, metadata = $5, updated_at = NOW()
		WHERE id = $6
	`
	_, err := r.db.ExecContext(
		ctx, query,
		entity.CollectionID,
		entity.Type,
		entity.Title,
		entity.Content,
		maybeNull(entity.Metadata),
		entity.ID,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *ResourceRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM resources WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func maybeNull(b []byte) interface{} {
	if len(b) == 0 {
		return nil
	}
	return string(b)
}
