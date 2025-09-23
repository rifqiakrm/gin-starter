package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/common/interfaces"
	"gin-starter/entity"
)

// UserUpdaterRepositoryUseCase defines the interface for updating user records.
type UserUpdaterRepositoryUseCase interface {
	// Update modifies an existing user in the database.
	Update(ctx context.Context, e *entity.User) error
}

// UserUpdaterRepository is the GORM implementation of UserUpdaterRepository.
type UserUpdaterRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// NewUserUpdaterRepository creates a new UserUpdaterRepository.
func NewUserUpdaterRepository(db *gorm.DB, cache interfaces.Cacheable) *UserUpdaterRepository {
	return &UserUpdaterRepository{
		db:    db,
		cache: cache,
	}
}

// Update modifies an existing user in the database.
func (r *UserUpdaterRepository) Update(ctx context.Context, e *entity.User) error {
	if err := r.db.WithContext(ctx).Save(e).Error; err != nil {
		return errors.Wrap(err, "[UserUpdaterRepository-Update] failed to update user")
	}
	return nil
}
