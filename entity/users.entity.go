package entity

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gin-starter/common/helper"
)

const (
	usersTableName = "auth.users"
)

// User entity
type User struct {
	ID                  int64          `json:"id"`
	Name                string         `json:"name"`
	Title               sql.NullString `json:"title"`
	Email               string         `json:"email"`
	PhoneNumber         sql.NullString `json:"phone_number"`
	DOB                 sql.NullTime   `json:"dob"`
	Photo               sql.NullString `json:"photo"`
	Password            string         `json:"password"`
	ForgotPasswordToken sql.NullString `json:"forgot_password_token"`
	OTP                 sql.NullString `json:"otp"`
	Status              string         `json:"status"`
	OrgUnitID           sql.NullInt64  `json:"org_unit_id"`
	Auditable
}

// TableName define table name of the struct
func (m *User) TableName() string {
	return usersTableName
}

// NewUser is a function to create new user struct
func NewUser(
	id int64,
	name string,
	email string,
	password string,
	dob sql.NullTime,
	photo string,
	phoneNumber string,
	createdBy string,
) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		Password:    string(passwordHash),
		PhoneNumber: helper.StringToNullString(phoneNumber),
		Photo:       helper.StringToNullString(photo),
		DOB:         dob,
		OTP:         sql.NullString{},
		Status:      "ACTIVATED",
		Auditable:   NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (m *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":         m.Name,
			"email":        m.Email,
			"phone_number": m.PhoneNumber,
			"photo":        m.Photo,
			"otp":          m.OTP,
			"status":       m.Status,
			"updated_at":   m.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if m.Name != from.Name {
		mapped["name"] = from.Name
	}

	if m.Email != from.Email {
		mapped["email"] = from.Email
	}

	if m.PhoneNumber != from.PhoneNumber {
		mapped["phone_number"] = from.PhoneNumber
	}

	if m.DOB != from.DOB {
		mapped["dob"] = from.DOB
	}

	if (m.Photo != from.Photo) && from.Photo.String != "" {
		mapped["photo"] = from.Photo
	}

	if m.OTP != from.OTP {
		mapped["otp"] = helper.StringToNullString(from.OTP.String)
	}

	if m.Status != from.Status {
		mapped["status"] = from.Status
	}

	if m.ForgotPasswordToken != from.ForgotPasswordToken {
		mapped["forgot_password_token"] = from.ForgotPasswordToken
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
