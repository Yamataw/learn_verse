package service

import (
	"context"
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

func (s *CollectionService) Create(ctx context.Context, collection models.ResourceCollection) (models.ResourceCollection, error) {
	// ici on pourrait valider, vérifier droits, etc.
	return s.repo.Create(ctx, collection)
}

// Get récupère une collection par ID
func (s *CollectionService) Get(ctx context.Context, id models.ULID) (models.ResourceCollection, error) {
	return s.repo.GetByID(ctx, id)
}

// List renvoie toutes les collections
func (s *CollectionService) List(ctx context.Context) ([]models.ResourceCollection, error) {
	return s.repo.List(ctx)
}
