package service

import (
	"context"
)

// Définition de l’interface minimale que doit implémenter votre repo
type CRUDRepo[T any, ID comparable] interface {
	Create(ctx context.Context, e T) (T, error)
	GetByID(ctx context.Context, id ID) (T, error)
	List(ctx context.Context) ([]T, error)
	Update(ctx context.Context, e T) (T, error)
	Delete(ctx context.Context, id ID) error
}

// BaseService implémente les méthodes CRUD en déléguant au repo
type BaseService[T any, ID comparable, R CRUDRepo[T, ID]] struct {
	Repo R
}

func (s *BaseService[T, ID, R]) Create(ctx context.Context, e T) (T, error) {
	return s.Repo.Create(ctx, e)
}

func (s *BaseService[T, ID, R]) Get(ctx context.Context, id ID) (T, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *BaseService[T, ID, R]) List(ctx context.Context) ([]T, error) {
	return s.Repo.List(ctx)
}

func (s *BaseService[T, ID, R]) Update(ctx context.Context, e T) (T, error) {
	return s.Repo.Update(ctx, e)
}

func (s *BaseService[T, ID, R]) Delete(ctx context.Context, id ID) error {
	return s.Repo.Delete(ctx, id)
}
