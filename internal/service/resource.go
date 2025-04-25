package service

import (
	"context"
	"encoding/json"
	"learn_verse/internal/models"
	"learn_verse/internal/repository"
)

// ResourceService regroupe la logique métier
type ResourceService struct {
	repo *repository.ResourceRepo
}

type ResourceServiceInterface interface {
	Create(ctx context.Context, collectionID *models.ULID, typ string, title string, content, metadata json.RawMessage) (models.Resource, error)
}

func NewResourceService(repo *repository.ResourceRepo) *ResourceService {
	return &ResourceService{repo: repo}
}

func (s *ResourceService) Create(ctx context.Context, collID *models.ULID, t, title string, content, meta []byte) (*models.Resource, error) {
	res := &models.Resource{
		CollectionID: collID,
		Type:         t,
		Title:        title,
		Content:      content,
		Metadata:     meta,
	}
	return s.repo.Create(ctx, res)
}
