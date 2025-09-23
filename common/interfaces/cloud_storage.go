// Package interfaces provide interfaces
//
//revive:disable:var-naming
package interfaces

import (
	"context"
	"mime/multipart"
)

// CloudStorageUseCase define interface for Cloud Storage
type CloudStorageUseCase interface {
	Upload(ctx context.Context, f *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, key string) error
}
