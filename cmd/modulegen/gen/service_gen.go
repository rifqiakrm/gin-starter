package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ServiceData holds template data for service generation
type ServiceData struct {
	Schema      string
	Version     string
	EntityLower string
	EntityUpper string
}

// singularize does a very naive plural-to-singular conversion
func singularize(word string) string {
	word = strings.ToLower(word)
	if strings.HasSuffix(word, "ies") {
		// e.g. companies -> company
		return word[:len(word)-3] + "y"
	}
	if strings.HasSuffix(word, "ses") {
		// e.g. addresses -> address
		return word[:len(word)-2]
	}
	if strings.HasSuffix(word, "s") && len(word) > 1 {
		// generic: users -> user
		return word[:len(word)-1]
	}
	return word
}

var serviceTemplates = map[string]string{
	"creator": `package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/repository"
)

// {{.EntityUpper}}Creator handles creation logic for {{.EntityUpper}}
type {{.EntityUpper}}Creator struct {
	cfg  config.Config
	repo repository.{{.EntityUpper}}CreatorRepositoryUseCase
}

// {{.EntityUpper}}CreatorUseCase defines the creation use case
type {{.EntityUpper}}CreatorUseCase interface {
	Create{{.EntityUpper}}(ctx context.Context, data *entity.{{.EntityUpper}}) (*entity.{{.EntityUpper}}, error)
}

// New{{.EntityUpper}}Creator returns a new {{.EntityUpper}}Creator
func New{{.EntityUpper}}Creator(cfg config.Config, repo repository.{{.EntityUpper}}CreatorRepositoryUseCase) *{{.EntityUpper}}Creator {
	return &{{.EntityUpper}}Creator{cfg: cfg, repo: repo}
}

// Create{{.EntityUpper}} creates a new {{.EntityUpper}}
func (uc *{{.EntityUpper}}Creator) Create{{.EntityUpper}}(ctx context.Context, data *entity.{{.EntityUpper}}) (*entity.{{.EntityUpper}}, error) {
	if err := uc.repo.Create(ctx, data); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return data, nil
}
`,

	"finder": `package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/repository"
	"github.com/google/uuid"
)

// {{.EntityUpper}}Finder handles find logic for {{.EntityUpper}}
type {{.EntityUpper}}Finder struct {
	cfg  config.Config
	repo repository.{{.EntityUpper}}FinderRepositoryUseCase
}

// {{.EntityUpper}}FinderUseCase defines the find use case
type {{.EntityUpper}}FinderUseCase interface {
	// Get{{.EntityUpper}}ByID retrieves a {{.EntityUpper}} by ID
	Get{{.EntityUpper}}ByID(ctx context.Context, id string) (*entity.{{.EntityUpper}}, error)
}

// New{{.EntityUpper}}Finder returns a new {{.EntityUpper}}Finder
func New{{.EntityUpper}}Finder(cfg config.Config, repo repository.{{.EntityUpper}}FinderRepositoryUseCase) *{{.EntityUpper}}Finder {
	return &{{.EntityUpper}}Finder{cfg: cfg, repo: repo}
}

// Get{{.EntityUpper}}ByID retrieves a {{.EntityUpper}} by ID
func (uc *{{.EntityUpper}}Finder) Get{{.EntityUpper}}ByID(ctx context.Context, id string) (*entity.{{.EntityUpper}}, error) {
	result, err := uc.repo.FindByID(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.ErrRecordNotFound.Error()
	}
	return result, nil
}
`,

	"updater": `package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/repository"
)

// {{.EntityUpper}}Updater handles update logic for {{.EntityUpper}}
type {{.EntityUpper}}Updater struct {
	cfg  config.Config
	repo repository.{{.EntityUpper}}UpdaterRepositoryUseCase
}

// {{.EntityUpper}}UpdaterUseCase defines the update use case
type {{.EntityUpper}}UpdaterUseCase interface {
	Update{{.EntityUpper}}(ctx context.Context, data *entity.{{.EntityUpper}}) (*entity.{{.EntityUpper}}, error)
}

// New{{.EntityUpper}}Updater returns a new {{.EntityUpper}}Updater
func New{{.EntityUpper}}Updater(cfg config.Config, repo repository.{{.EntityUpper}}UpdaterRepositoryUseCase) *{{.EntityUpper}}Updater {
	return &{{.EntityUpper}}Updater{cfg: cfg, repo: repo}
}

// Update{{.EntityUpper}} updates an existing {{.EntityUpper}}
func (uc *{{.EntityUpper}}Updater) Update{{.EntityUpper}}(ctx context.Context, data *entity.{{.EntityUpper}}) (*entity.{{.EntityUpper}}, error) {
	if err := uc.repo.Update(ctx, data); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return data, nil
}
`,

	"deleter": `package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/repository"
	"github.com/google/uuid"
)

// {{.EntityUpper}}Deleter handles delete logic for {{.EntityUpper}}
type {{.EntityUpper}}Deleter struct {
	cfg  config.Config
	repo repository.{{.EntityUpper}}DeleterRepositoryUseCase
}

// {{.EntityUpper}}DeleterUseCase defines the delete use case
type {{.EntityUpper}}DeleterUseCase interface {
	// Delete{{.EntityUpper}}ByID deletes a {{.EntityUpper}} by ID
	Delete{{.EntityUpper}}ByID(ctx context.Context, id string) error
}

// New{{.EntityUpper}}Deleter returns a new {{.EntityUpper}}Deleter
func New{{.EntityUpper}}Deleter(cfg config.Config, repo repository.{{.EntityUpper}}DeleterRepositoryUseCase) *{{.EntityUpper}}Deleter {
	return &{{.EntityUpper}}Deleter{cfg: cfg, repo: repo}
}

// Delete{{.EntityUpper}}ByID deletes a {{.EntityUpper}} by ID
func (uc *{{.EntityUpper}}Deleter) Delete{{.EntityUpper}}ByID(ctx context.Context, id string) error {
	if err := uc.repo.Delete(ctx, uuid.MustParse(id), "system"); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
`,
}

func GenerateServices(schema, version, entity string) error {
	// Normalize to singular
	singular := singularize(entity)

	data := ServiceData{
		Schema:      schema,
		Version:     version,
		EntityLower: strings.ToLower(singular),
		EntityUpper: strings.Title(singular),
	}

	dir := filepath.Join("modules", schema, version, "service")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for action, tmpl := range serviceTemplates {
		t, err := template.New(action).Parse(tmpl)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return err
		}

		filename := filepath.Join(dir, fmt.Sprintf("%s_%s.service.go", data.EntityLower, action))
		if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
			return err
		}
		fmt.Println("✅ Generated:", filename)
	}
	return nil
}
