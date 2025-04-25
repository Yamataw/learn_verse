package service

import (
	"learn_verse/internal/models"
	"learn_verse/internal/repository"
)

type UserService struct {
	*BaseService[models.User, models.ULID, *repository.UserRepo]
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		BaseService: &BaseService[models.User, models.ULID, *repository.UserRepo]{Repo: repo},
	}
}
