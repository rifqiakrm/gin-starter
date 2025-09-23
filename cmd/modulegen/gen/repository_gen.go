package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// RepositoryData holds template data for repository generation.
type RepositoryData struct {
	Schema      string
	Version     string
	EntityLower string
	EntityUpper string
}

var repositoryTemplates = map[string]string{
	"creator": `package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gin-starter/entity"
)

// {{.EntityUpper}}CreatorRepositoryUseCase defines the interface for creating {{.EntityLower}} records.
type {{.EntityUpper}}CreatorRepositoryUseCase interface {
	// Create inserts a new {{.EntityLower}} into the database.
	Create(ctx context.Context, e *entity.{{.EntityUpper}}) error
}

// {{.EntityUpper}}CreatorRepository is the GORM implementation of {{.EntityUpper}}CreatorRepository.
type {{.EntityUpper}}CreatorRepository struct {
	db *gorm.DB
	cache interfaces.Cacheable
}

// New{{.EntityUpper}}CreatorRepository creates a new {{.EntityUpper}}CreatorRepository.
func New{{.EntityUpper}}CreatorRepository(db *gorm.DB, cache interfaces.Cacheable) *{{.EntityUpper}}CreatorRepository {
	return &{{.EntityUpper}}CreatorRepository{
		db: db,
		cache: cache,
	}
}

// Create inserts a new {{.EntityLower}} into the database.
func (r *{{.EntityUpper}}CreatorRepository) Create(ctx context.Context, e *entity.{{.EntityUpper}}) error {
	if err := r.db.WithContext(ctx).Create(e).Error; err != nil {
		return errors.Wrap(err, "[{{.EntityUpper}}CreatorRepository-Create] failed to create {{.EntityLower}}")
	}
	return nil
}
`,

	"finder": `package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gin-starter/entity"
)

// {{.EntityUpper}}FinderRepositoryUseCase defines the interface for retrieving {{.EntityLower}} records.
type {{.EntityUpper}}FinderRepositoryUseCase interface {
	// FindByID retrieves a {{.EntityLower}} by its ID.
	FindByID(ctx context.Context, id uuid.UUID) (*entity.{{.EntityUpper}}, error)
	// FindAll retrieves a list of {{.EntityLower}} records with pagination.
	FindAll(ctx context.Context, limit, offset int) ([]*entity.{{.EntityUpper}}, error)
}

// {{.EntityUpper}}FinderRepository is the GORM implementation of {{.EntityUpper}}FinderRepository.
type {{.EntityUpper}}FinderRepository struct {
	db *gorm.DB
	cache interfaces.Cacheable
}

// New{{.EntityUpper}}FinderRepository creates a new {{.EntityUpper}}FinderRepository.
func New{{.EntityUpper}}FinderRepository(db *gorm.DB, cache interfaces.Cacheable) *{{.EntityUpper}}FinderRepository {
	return &{{.EntityUpper}}FinderRepository{
		db: db,
		cache: cache,
	}
}

// FindByID retrieves a {{.EntityLower}} by its ID.
func (r *{{.EntityUpper}}FinderRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.{{.EntityUpper}}, error) {
	var e entity.{{.EntityUpper}}
	if err := r.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[{{.EntityUpper}}FinderRepository-FindByID] failed to find {{.EntityLower}}")
	}
	return &e, nil
}

// FindAll retrieves a list of {{.EntityLower}} records with pagination.
func (r *{{.EntityUpper}}FinderRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.{{.EntityUpper}}, error) {
	var list []*entity.{{.EntityUpper}}
	query := r.db.WithContext(ctx).Model(&entity.{{.EntityUpper}}{})
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[{{.EntityUpper}}FinderRepository-FindAll] failed to find list of {{.EntityLower}}")
	}
	return list, nil
}
`,

	"updater": `package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gin-starter/entity"
)

// {{.EntityUpper}}UpdaterRepositoryUseCase defines the interface for updating {{.EntityLower}} records.
type {{.EntityUpper}}UpdaterRepositoryUseCase interface {
	// Update modifies an existing {{.EntityLower}} in the database.
	Update(ctx context.Context, e *entity.{{.EntityUpper}}) error
}

// {{.EntityUpper}}UpdaterRepository is the GORM implementation of {{.EntityUpper}}UpdaterRepository.
type {{.EntityUpper}}UpdaterRepository struct {
	db *gorm.DB
	cache interfaces.Cacheable
}

// New{{.EntityUpper}}UpdaterRepository creates a new {{.EntityUpper}}UpdaterRepository.
func New{{.EntityUpper}}UpdaterRepository(db *gorm.DB, cache interfaces.Cacheable) *{{.EntityUpper}}UpdaterRepository {
	return &{{.EntityUpper}}UpdaterRepository{
		db: db,
		cache: cache,
	}
}

// Update modifies an existing {{.EntityLower}} in the database.
func (r *{{.EntityUpper}}UpdaterRepository) Update(ctx context.Context, e *entity.{{.EntityUpper}}) error {
	if err := r.db.WithContext(ctx).Save(e).Error; err != nil {
		return errors.Wrap(err, "[{{.EntityUpper}}UpdaterRepository-Update] failed to update {{.EntityLower}}")
	}
	return nil
}
`,

	"deleter": `package repository

import (
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gin-starter/entity"
)

// {{.EntityUpper}}DeleterRepositoryUseCase defines the interface for soft-deleting {{.EntityLower}} records.
type {{.EntityUpper}}DeleterRepositoryUseCase interface {
	// Delete performs a soft-delete by updating deleted_by and deleted_at fields.
	Delete(ctx context.Context, id uuid.UUID, deletedBy string) error
}

// {{.EntityUpper}}DeleterRepository is the GORM implementation of {{.EntityUpper}}DeleterRepository.
type {{.EntityUpper}}DeleterRepository struct {
	db *gorm.DB
	cache interfaces.Cacheable
}

// New{{.EntityUpper}}DeleterRepository creates a new {{.EntityUpper}}DeleterRepository.
func New{{.EntityUpper}}DeleterRepository(db *gorm.DB, cache interfaces.Cacheable) *{{.EntityUpper}}DeleterRepository {
	return &{{.EntityUpper}}DeleterRepository{
		db: db,
		cache: cache,
	}
}

// Delete performs a soft-delete by updating deleted_by and deleted_at fields.
func (r *{{.EntityUpper}}DeleterRepository) Delete(ctx context.Context, id uuid.UUID, deletedBy string) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.{{.EntityUpper}}{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_by": deletedBy,
			"updated_at": time.Now(),
			"deleted_at": time.Now(),
		}).Error; err != nil {
		return errors.Wrap(err, "[{{.EntityUpper}}DeleterRepository-Delete] failed to delete {{.EntityLower}}")
	}
	return nil
}
`,
}

// GenerateRepositories creates repository code files (creator, finder, updater, deleter)
// for the given schema, version, and entity name. The entity name is automatically
// normalized to singular form for cleaner naming conventions.
func GenerateRepositories(schema, version, entity string) error {
	// Normalize entity to singular
	singular := singularize(entity)

	data := RepositoryData{
		Schema:      schema,
		Version:     version,
		EntityLower: strings.ToLower(singular),
		EntityUpper: strings.Title(singular),
	}

	dir := filepath.Join("modules", schema, version, "repository")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for action, tmpl := range repositoryTemplates {
		t, err := template.New(action).Parse(tmpl)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return err
		}

		filename := filepath.Join(dir, fmt.Sprintf("%s_%s.repository.go", data.EntityLower, action))
		if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
			return err
		}
		fmt.Println("✅ Generated:", filename)
	}
	return nil
}
