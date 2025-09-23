package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/common/interfaces"
	"gin-starter/entity"
)

// UserFinderRepositoryUseCase defines the interface for retrieving user records.
type UserFinderRepositoryUseCase interface {
	// FindByID retrieves a user by its ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// FindAll retrieves a list of user records with pagination.
	FindAll(ctx context.Context, limit, offset int) ([]*entity.User, error)
}

// UserFinderRepository is the GORM implementation of UserFinderRepository.
type UserFinderRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// NewUserFinderRepository creates a new UserFinderRepository.
func NewUserFinderRepository(db *gorm.DB, cache interfaces.Cacheable) *UserFinderRepository {
	return &UserFinderRepository{
		db:    db,
		cache: cache,
	}
}

// FindByID retrieves a user by its ID.
func (r *UserFinderRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var e entity.User
	if err := r.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserFinderRepository-FindByID] failed to find user")
	}
	return &e, nil
}

// FindAll retrieves a list of user records with pagination.
func (r *UserFinderRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	var list []*entity.User
	query := r.db.WithContext(ctx).Model(&entity.User{})
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[UserFinderRepository-FindAll] failed to find list of user")
	}
	return list, nil
}
