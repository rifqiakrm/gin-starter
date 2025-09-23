package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/common/interfaces"
	"gin-starter/entity"
)

// UserCreatorRepositoryUseCase defines the interface for creating user records.
type UserCreatorRepositoryUseCase interface {
	// Create inserts a new user into the database.
	Create(ctx context.Context, e *entity.User) error
}

// UserCreatorRepository is the GORM implementation of UserCreatorRepository.
type UserCreatorRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// NewUserCreatorRepository creates a new UserCreatorRepository.
func NewUserCreatorRepository(db *gorm.DB, cache interfaces.Cacheable) *UserCreatorRepository {
	return &UserCreatorRepository{
		db:    db,
		cache: cache,
	}
}

// Create inserts a new user into the database.
func (r *UserCreatorRepository) Create(ctx context.Context, e *entity.User) error {
	if err := r.db.WithContext(ctx).Create(e).Error; err != nil {
		return errors.Wrap(err, "[UserCreatorRepository-Create] failed to create user")
	}
	return nil
}
