package service

import (
	"context"

	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/repository"
)

// UserUpdater handles update logic for User
type UserUpdater struct {
	cfg  config.Config
	repo repository.UserUpdaterRepositoryUseCase
}

// UserUpdaterUseCase defines the update use case
type UserUpdaterUseCase interface {
	UpdateUser(ctx context.Context, data *entity.User) (*entity.User, error)
}

// NewUserUpdater returns a new UserUpdater
func NewUserUpdater(cfg config.Config, repo repository.UserUpdaterRepositoryUseCase) *UserUpdater {
	return &UserUpdater{cfg: cfg, repo: repo}
}

// UpdateUser updates an existing User
func (uc *UserUpdater) UpdateUser(ctx context.Context, data *entity.User) (*entity.User, error) {
	if err := uc.repo.Update(ctx, data); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return data, nil
}
