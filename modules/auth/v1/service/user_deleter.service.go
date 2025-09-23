package service

import (
	"context"

	"github.com/google/uuid"

	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/auth/v1/repository"
)

// UserDeleter handles delete logic for User
type UserDeleter struct {
	cfg  config.Config
	repo repository.UserDeleterRepositoryUseCase
}

// UserDeleterUseCase defines the delete use case
type UserDeleterUseCase interface {
	// DeleteUserByID deletes a User by ID
	DeleteUserByID(ctx context.Context, id string) error
}

// NewUserDeleter returns a new UserDeleter
func NewUserDeleter(cfg config.Config, repo repository.UserDeleterRepositoryUseCase) *UserDeleter {
	return &UserDeleter{cfg: cfg, repo: repo}
}

// DeleteUserByID deletes a User by ID
func (uc *UserDeleter) DeleteUserByID(ctx context.Context, id string) error {
	if err := uc.repo.Delete(ctx, uuid.MustParse(id), "system"); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
