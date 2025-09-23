// Package entity provides entity for database tables
package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"gin-starter/common/helper"
)

// Auditable is an interface that can be embedded in structs that need to be auditable
type Auditable struct {
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedBy sql.NullString `json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// NewAuditable creates a new Auditable struct
func NewAuditable(createdBy string) Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: helper.StringToNullString(createdBy),
		UpdatedBy: helper.StringToNullString(createdBy),
	}
}
