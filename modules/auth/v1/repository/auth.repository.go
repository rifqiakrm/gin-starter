// Package repository provides functions for querying to database
package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gin-starter/common/helper"
	"gin-starter/entity"
)

// AuthRepository is a repository for auth
type AuthRepository struct {
	db *gorm.DB
}

// AuthRepositoryUseCase is a repository for auth
type AuthRepositoryUseCase interface {
	// GetUserByEmail finds a auth by email
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	// GetAdminByEmail finds a admin by email
	GetAdminByEmail(ctx context.Context, email string) (*entity.User, error)
	// UpdateOTP updates OTP
	UpdateOTP(ctx context.Context, user *entity.User, otp string) error
}

// NewAuthRepository returns a auth repository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

// GetUserByEmail finds a auth by email
func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Where("email = ?", email).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByEmail] email not found")
	}

	return result, nil
}

// UpdateOTP updates OTP
func (ar *AuthRepository) UpdateOTP(ctx context.Context, user *entity.User, otp string) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ar.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-ChangePassword] error when updating data")
			}
			if err := tx.Model(&entity.User{}).
				Where(`id = ?`, user.ID).Update("otp", helper.StringToNullString(otp)).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-Update] error when updating data User b")
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}
	return nil
}

// GetAdminByEmail finds a admin by email
func (ar *AuthRepository) GetAdminByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Joins("inner join auth.user_roles on auth.users.id=auth.user_roles.user_id").
		Where("email = ?", email).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetAdminByEmail] email not found")
	}

	return result, nil
}
