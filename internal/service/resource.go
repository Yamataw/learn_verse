package service

import (
	"context"
	"learn_verse/internal/models"
	"learn_verse/internal/repository"
)

// ResourceService regroupe la logique métier
type ResourceService struct {
	repo *repository.ResourceRepo
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
