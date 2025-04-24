package repository

import "context"

type Repository[T any, ID comparable] interface {
	Create(ctx context.Context, entity T) (T, error)
	GetByID(ctx context.Context, id ID) (T, error)
	List(ctx context.Context) ([]T, error)
	Update(ctx context.Context, entity T) (T, error)
	Delete(ctx context.Context, id ID) error
}
