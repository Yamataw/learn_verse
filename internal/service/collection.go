package service

import (
	"context"
	"learn_verse/internal/models"
	"learn_verse/internal/repository"
)

// CollectionService regroupe la logique métier

type CollectionServiceInterface interface {
	Create(ctx context.Context, collection models.ResourceCollection) (models.ResourceCollection, error)
	Get(ctx context.Context, id models.ULID) (models.ResourceCollection, error)
	List(ctx context.Context) ([]models.ResourceCollection, error)
	Update(ctx context.Context, collection models.ResourceCollection) (models.ResourceCollection, error)
	Delete(ctx context.Context, id models.ULID) error
}

type CollectionService struct {
	*BaseService[models.ResourceCollection, models.ULID, *repository.CollectionRepo]
}

// NewCollectionService crée un service pour ResourceCollection
func NewCollectionService(repo *repository.CollectionRepo) *CollectionService {
	return &CollectionService{
		BaseService: &BaseService[models.ResourceCollection, models.ULID, *repository.CollectionRepo]{Repo: repo},
	}
}
