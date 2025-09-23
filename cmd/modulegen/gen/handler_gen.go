package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type HandlerData struct {
	Schema      string
	Version     string
	EntityLower string
	EntityUpper string
}

var handlerTemplates = map[string]string{
	"creator": `package handler

import (
	"gin-starter/common/interfaces"
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-starter/common/errors"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/service"
	"gin-starter/resource"
	"gin-starter/response"
)

// {{.EntityUpper}}CreatorHandler handles HTTP requests for creating {{.EntityLower}}.
type {{.EntityUpper}}CreatorHandler struct {
	{{.EntityLower}}UseCase service.{{.EntityUpper}}CreatorUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// New{{.EntityUpper}}CreatorHandler creates a new {{.EntityUpper}}CreatorHandler.
func New{{.EntityUpper}}CreatorHandler({{.EntityLower}}UseCase service.{{.EntityUpper}}CreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) *{{.EntityUpper}}CreatorHandler {
	return &{{.EntityUpper}}CreatorHandler{
		{{.EntityLower}}UseCase: {{.EntityLower}}UseCase,
		cloudStorage: cloudStorage,
	}
}

// Create{{.EntityUpper}} handles the HTTP request to create a new {{.EntityLower}}.
func (h *{{.EntityUpper}}CreatorHandler) Create{{.EntityUpper}}(c *gin.Context) {
	var req resource.Create{{.EntityUpper}}Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.{{.EntityLower}}UseCase.Create{{.EntityUpper}}(c, req)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
`,

	"finder": `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-starter/common/errors"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/service"
	"gin-starter/response"
)

// {{.EntityUpper}}FinderHandler handles HTTP requests for retrieving {{.EntityLower}}.
type {{.EntityUpper}}FinderHandler struct {
	{{.EntityLower}}UseCase service.{{.EntityUpper}}FinderUseCase
}

// New{{.EntityUpper}}FinderHandler creates a new {{.EntityUpper}}FinderHandler.
func New{{.EntityUpper}}FinderHandler({{.EntityLower}}UseCase service.{{.EntityUpper}}FinderUseCase) *{{.EntityUpper}}FinderHandler {
	return &{{.EntityUpper}}FinderHandler{
		{{.EntityLower}}UseCase: {{.EntityLower}}UseCase,
	}
}

// Get{{.EntityUpper}}ByID handles the HTTP request to retrieve a {{.EntityLower}} by ID.
func (h *{{.EntityUpper}}FinderHandler) Get{{.EntityUpper}}ByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.{{.EntityLower}}UseCase.Get{{.EntityUpper}}ByID(c, id)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
`,

	"updater": `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-starter/common/errors"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/service"
	"gin-starter/resource"
	"gin-starter/response"
)

// {{.EntityUpper}}UpdaterHandler handles HTTP requests for updating {{.EntityLower}}.
type {{.EntityUpper}}UpdaterHandler struct {
	{{.EntityLower}}UseCase service.{{.EntityUpper}}UpdaterUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// New{{.EntityUpper}}UpdaterHandler creates a new {{.EntityUpper}}UpdaterHandler.
func New{{.EntityUpper}}UpdaterHandler({{.EntityLower}}UseCase service.{{.EntityUpper}}UpdaterUseCase, cloudStorage interfaces.CloudStorageUseCase) *{{.EntityUpper}}UpdaterHandler {
	return &{{.EntityUpper}}UpdaterHandler{
		{{.EntityLower}}UseCase: {{.EntityLower}}UseCase,
		cloudStorage: cloudStorage,
	}
}

// Update{{.EntityUpper}} handles the HTTP request to update an existing {{.EntityLower}}.
func (h *{{.EntityUpper}}UpdaterHandler) Update{{.EntityUpper}}(c *gin.Context) {
	var req resource.Update{{.EntityUpper}}Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.{{.EntityLower}}UseCase.Update{{.EntityUpper}}(c, req)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
`,

	"deleter": `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-starter/common/errors"
	"gin-starter/modules/{{.Schema}}/{{.Version}}/service"
	"gin-starter/response"
)

// {{.EntityUpper}}DeleterHandler handles HTTP requests for deleting {{.EntityLower}}.
type {{.EntityUpper}}DeleterHandler struct {
	{{.EntityLower}}UseCase service.{{.EntityUpper}}DeleterUseCase
}

// New{{.EntityUpper}}DeleterHandler creates a new {{.EntityUpper}}DeleterHandler.
func New{{.EntityUpper}}DeleterHandler({{.EntityLower}}UseCase service.{{.EntityUpper}}DeleterUseCase) *{{.EntityUpper}}DeleterHandler {
	return &{{.EntityUpper}}DeleterHandler{
		{{.EntityLower}}UseCase: {{.EntityLower}}UseCase,
	}
}

// Delete{{.EntityUpper}}ByID handles the HTTP request to delete a {{.EntityLower}} by ID.
func (h *{{.EntityUpper}}DeleterHandler) Delete{{.EntityUpper}}ByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.{{.EntityLower}}UseCase.Delete{{.EntityUpper}}ByID(c, id); err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", nil))
}
`,
}

// GenerateHandlers creates handler code files (creator, finder, updater, deleter)
// for the given schema, version, and entity name. The entity name is automatically
// normalized to singular form for cleaner naming conventions.
func GenerateHandlers(schema, version, entity string) error {
	// Normalize entity to singular
	singular := singularize(entity)

	data := HandlerData{
		Schema:      schema,
		Version:     version,
		EntityLower: strings.ToLower(singular),
		EntityUpper: strings.Title(singular),
	}

	dir := filepath.Join("modules", schema, version, "handler")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for action, tmpl := range handlerTemplates {
		t, err := template.New(action).Parse(tmpl)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return err
		}

		filename := filepath.Join(dir, fmt.Sprintf("%s_%s.handler.go", data.EntityLower, action))
		if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
			return err
		}
		fmt.Println("✅ Generated:", filename)
	}
	return nil
}
