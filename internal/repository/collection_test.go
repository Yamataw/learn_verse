package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"learn_verse/internal/models"
	"testing"
	"time"
)

func TestCollectionRepo_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewCollectionRepo(db)
	now := time.Now()
	desc := "test desc"

	expected := models.ResourceCollection{
		ID:          models.ULID(ulid.Make()),
		Name:        "Ma Collection",
		Description: &desc,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mock.ExpectQuery(`(?i)INSERT\s+INTO\s+resource_collections\s*\(\s*name\s*,\s*description\s*\)\s*VALUES\s*\(\s*\$1\s*,\s*\$2\s*\)\s*RETURNING\s+id\s*,\s*name\s*,\s*description\s*,\s*created_at\s*,\s*updated_at`).
		WithArgs("Ma Collection", &desc).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(expected.ID, "Ma Collection", desc, now, now))

	got, err := repo.Create(context.Background(), models.ResourceCollection{
		Name:        expected.Name,
		Description: expected.Description,
	})

	assert.NoError(t, err)
	assert.Equal(t, expected.Name, got.Name)
	assert.Equal(t, expected.Description, got.Description)
}

func TestCollectionRepo_GetByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	id := models.ULID(ulid.Make())
	now := time.Now()
	name := "Collection A"
	desc := "desc"

	mock.ExpectQuery(`SELECT id, name, description`).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(id, name, desc, now, now))

	result, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, desc, *result.Description)
}

func TestCollectionRepo_GetByID_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	id := models.ULID(ulid.Make())

	mock.ExpectQuery(`SELECT id, name, description`).
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	_, err := repo.GetByID(context.Background(), id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "collection not found")
}

func TestCollectionRepo_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	now := time.Now()
	desc := "listed"
	mock.ExpectQuery(`SELECT id, name, description`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(models.ULID(ulid.Make()), "Col A", desc, now, now).
			AddRow(models.ULID(ulid.Make()), "Col B", desc, now, now))

	result, err := repo.List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestCollectionRepo_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	id := models.ULID(ulid.Make())
	now := time.Now()
	desc := "updated"

	mock.ExpectQuery(`UPDATE resource_collections`).
		WithArgs("Updated Name", &desc, id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
			AddRow(id, "Updated Name", desc, now, now))

	result, err := repo.Update(context.Background(), models.ResourceCollection{
		ID:          id,
		Name:        "Updated Name",
		Description: &desc,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", result.Name)
}

func TestCollectionRepo_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	id := models.ULID(ulid.Make())
	mock.ExpectExec(`(?i)UPDATE resource_collections\s+SET\s+deleted_at\s*=\s*now\(\)\s+WHERE\s+id\s*=\s*\$1\s+AND\s+deleted_at\s+IS\s+NULL`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(context.Background(), id)
	assert.NoError(t, err)
}

func TestCollectionRepo_Delete_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewCollectionRepo(db)

	id := models.ULID(ulid.Make())
	mock.ExpectExec(`(?i)UPDATE resource_collections\s+SET\s+deleted_at\s*=\s*now\(\)\s+WHERE\s+id\s*=\s*\$1\s+AND\s+deleted_at\s+IS\s+NULL`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(context.Background(), id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "collection not found")
}
