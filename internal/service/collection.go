package service

import (
	"context"
	"github.com/oklog/ulid/v2"
	"learn_verse/internal/models"
	"learn_verse/internal/repository"
)

// CollectionService regroupe la logique métier
type CollectionService struct {
	repo *repository.CollectionRepo
}

func NewCollectionService(repo *repository.CollectionRepo) *CollectionService {
	return &CollectionService{repo: repo}
}

func (s *CollectionService) Create(ctx context.Context, name string, desc *string) (*models.ResourceCollection, error) {
	// ici on pourrait valider, vérifier droits, etc.
	return s.repo.Create(ctx, name, desc)
}

// Get récupère une collection par ID
func (s *CollectionService) Get(ctx context.Context, id ulid.ULID) (*models.ResourceCollection, error) {
	return s.repo.FindByID(ctx, id)
}

// List renvoie toutes les collections
func (s *CollectionService) List(ctx context.Context) ([]models.ResourceCollection, error) {
	return s.repo.List(ctx)
}
