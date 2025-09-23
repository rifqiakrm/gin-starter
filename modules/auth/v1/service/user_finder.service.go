package service

import (
	"context"

	"github.com/google/uuid"

	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/repository"
)

// UserFinder handles find logic for User
type UserFinder struct {
	cfg  config.Config
	repo repository.UserFinderRepositoryUseCase
}

// UserFinderUseCase defines the find use case
type UserFinderUseCase interface {
	// GetUserByID retrieves a User by ID
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}

// NewUserFinder returns a new UserFinder
func NewUserFinder(cfg config.Config, repo repository.UserFinderRepositoryUseCase) *UserFinder {
	return &UserFinder{cfg: cfg, repo: repo}
}

// GetUserByID retrieves a User by ID
func (uc *UserFinder) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	result, err := uc.repo.FindByID(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.ErrRecordNotFound.Error()
	}
	return result, nil
}
