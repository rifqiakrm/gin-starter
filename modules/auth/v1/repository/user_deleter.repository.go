package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/common/interfaces"
	"gin-starter/entity"
)

// UserDeleterRepositoryUseCase defines the interface for soft-deleting user records.
type UserDeleterRepositoryUseCase interface {
	// Delete performs a soft-delete by updating deleted_by and deleted_at fields.
	Delete(ctx context.Context, id uuid.UUID, deletedBy string) error
}

// UserDeleterRepository is the GORM implementation of UserDeleterRepository.
type UserDeleterRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// NewUserDeleterRepository creates a new UserDeleterRepository.
func NewUserDeleterRepository(db *gorm.DB, cache interfaces.Cacheable) *UserDeleterRepository {
	return &UserDeleterRepository{
		db:    db,
		cache: cache,
	}
}

// Delete performs a soft-delete by updating deleted_by and deleted_at fields.
func (r *UserDeleterRepository) Delete(ctx context.Context, id uuid.UUID, deletedBy string) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_by": deletedBy,
			"updated_at": time.Now(),
			"deleted_at": time.Now(),
		}).Error; err != nil {
		return errors.Wrap(err, "[UserDeleterRepository-Delete] failed to delete user")
	}
	return nil
}
