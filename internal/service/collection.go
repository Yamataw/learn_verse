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

func (s *CollectionService) Create(ctx context.Context, name string, desc *string) (*models.ResourceCollection, error) {
	// ici on pourrait valider, vérifier droits, etc.
	return s.repo.Create(ctx, name, desc)
}
