package service

import (
	"context"

	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/repository"
)

// UserCreator handles creation logic for User
type UserCreator struct {
	cfg  config.Config
	repo repository.UserCreatorRepositoryUseCase
}

// UserCreatorUseCase defines the creation use case
type UserCreatorUseCase interface {
	CreateUser(ctx context.Context, data *entity.User) (*entity.User, error)
}

// NewUserCreator returns a new UserCreator
func NewUserCreator(cfg config.Config, repo repository.UserCreatorRepositoryUseCase) *UserCreator {
	return &UserCreator{cfg: cfg, repo: repo}
}

// CreateUser creates a new User
func (uc *UserCreator) CreateUser(ctx context.Context, data *entity.User) (*entity.User, error) {
	if err := uc.repo.Create(ctx, data); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return data, nil
}
