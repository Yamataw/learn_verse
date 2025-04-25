package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"learn_verse/internal/models"
)

// UserRepo offre l’accès SQL pour les utilisateurs
type UserRepo struct {
	db *sql.DB
}

// Assure que UserRepo implémente l'interface générique
var _ Repository[models.User, models.ULID] = (*UserRepo)(nil)

// NewUserRepo crée une nouvelle instance de UserRepo
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create insère un nouvel utilisateur et retourne l'entité créée
func (r *UserRepo) Create(ctx context.Context, entity models.User) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(entity.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `
	INSERT INTO users (username, email, password_hash)
	VALUES ($1, $2, $3)
	RETURNING id, username, email, password_hash
	`
	var user models.User
	err = r.db.QueryRowContext(ctx, query,
		entity.Username,
		entity.Email,
		hash,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

// GetByID récupère un utilisateur par son ULID
func (r *UserRepo) GetByID(ctx context.Context, id models.ULID) (models.User, error) {
	var user models.User
	query := `
	SELECT id, username, email, password_hash
	FROM users
	WHERE id = $1
	`
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found: %w", err)
		}
		return models.User{}, err
	}
	return user, nil
}

// List renvoie tous les utilisateurs
func (r *UserRepo) List(ctx context.Context) ([]models.User, error) {
	query := `
	SELECT id, username, email, password_hash
	FROM users
	ORDER BY username
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
		); err != nil {
			return nil, err
		}
		out = append(out, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// Update modifie un utilisateur existant et retourne l'entité mise à jour
func (r *UserRepo) Update(ctx context.Context, entity models.User) (models.User, error) {
	query := `
	UPDATE users
	SET username = $1,
	    email = $2,
	    password_hash = $3
	WHERE id = $4
	RETURNING id, username, email, password_hash
	`
	var user models.User
	err := r.db.QueryRowContext(ctx, query,
		entity.Username,
		entity.Email,
		entity.PasswordHash,
		entity.ID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found: %w", err)
		}
		return models.User{}, err
	}
	return user, nil
}

// Delete supprime un utilisateur (hard delete)
func (r *UserRepo) Delete(ctx context.Context, id models.ULID) error {

	query := `
	DELETE FROM users
	WHERE id = $1
	`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("user not found: %s", id)
	}
	return nil
}
