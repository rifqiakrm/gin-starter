package gen

import (
	"bytes"
	"strings"
	"text/template"
)

var entityTemplate = `package entity

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
)

const (
	{{.NameLower}}TableName = "{{.Schema}}.{{.Name}}"
)

// {{.NameUpper}} entity
type {{.NameUpper}} struct {
{{- range .Columns}}
{{- if not (IsAuditable .Name) }}
	{{.Name | ToCamel}} {{ GoType . }} ` + "`json:\"{{.Name}}\"`" + `
{{- end}}
{{- end}}
	Auditable
}

// TableName define table name of the struct
func (m *{{.NameUpper}}) TableName() string {
	return {{.NameLower}}TableName
}

// New{{.NameUpper}} constructor
func New{{.NameUpper}}(/* TODO: args */) *{{.NameUpper}} {
	return &{{.NameUpper}}{}
}

// MapUpdateFrom generates map[string]interface{} for updates
func (m *{{.NameUpper}}) MapUpdateFrom(from *{{.NameUpper}}) *map[string]interface{} {
	mapped := make(map[string]interface{})
	// TODO: generate diff logic
	return &mapped
}
`

// GenerateEntity renders Go code for a table
func GenerateEntity(table *Table) (string, error) {
	funcMap := template.FuncMap{
		"ToCamel": func(s string) string {
			parts := strings.Split(s, "_")
			for i := range parts {
				parts[i] = strings.Title(parts[i])
			}
			return strings.Join(parts, "")
		},
		"GoType": func(col Column) string {
			sqlType := strings.ToUpper(col.Type)
			nullable := col.Nullable
			switch {
			case strings.HasPrefix(sqlType, "UUID"):
				return "uuid.UUID"
			case strings.HasPrefix(sqlType, "VARCHAR"), strings.HasPrefix(sqlType, "TEXT"):
				if nullable {
					return "sql.NullString"
				}
				return "string"
			case strings.HasPrefix(sqlType, "DATE"), strings.HasPrefix(sqlType, "TIMESTAMPTZ"):
				if nullable {
					return "sql.NullTime"
				}
				return "time.Time"
			case strings.HasPrefix(sqlType, "INT"):
				if nullable {
					return "sql.NullInt64"
				}
				return "int64"
			case strings.HasPrefix(sqlType, "SERIAL"):
				if nullable {
					return "sql.NullInt64"
				}
				return "int64"
			default:
				return "string"
			}
		},
		"IsAuditable": func(name string) bool {
			switch strings.ToLower(name) {
			case "created_at", "updated_at", "deleted_at",
				"created_by", "updated_by", "deleted_by":
				return true
			}
			return false
		},
	}

	tmpl, err := template.New("entity").Funcs(funcMap).Parse(entityTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, table); err != nil {
		return "", err
	}

	return buf.String(), nil
}
