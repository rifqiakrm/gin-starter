package entity

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gin-starter/common/helper"
)

const (
	usersTableName = "auth.users"
)

// User represents the auth user entity.
type User struct {
	ID                  uuid.UUID      `json:"id"`
	EmployeeID          string         `json:"employee_id"`
	Name                string         `json:"name"`
	Title               sql.NullString `json:"title"`
	Email               string         `json:"email"`
	Username            sql.NullString `json:"username"`
	Password            string         `json:"password"`
	PhoneNumber         sql.NullString `json:"phone_number"`
	Address             sql.NullString `json:"address"`
	DOB                 sql.NullTime   `json:"dob"`
	Photo               sql.NullString `json:"photo"`
	ForgotPasswordToken sql.NullString `json:"forgot_password_token"`
	OTP                 sql.NullString `json:"otp"`
	Status              string         `json:"status"`
	OrgUnitID           sql.NullInt64  `json:"org_unit_id"`
	Auditable
}

// TableName returns the database table name for User.
func (m *User) TableName() string {
	return usersTableName
}

// NewUser creates a new User instance with hashed password and auditable info.
func NewUser(
	employeeID string,
	name string,
	email string,
	password string,
	address string,
	dob string,
	photo string,
	phoneNumber string,
	createdBy string,
) *User {
	// Generate hashed password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:          uuid.New(),
		EmployeeID:  employeeID,
		Name:        name,
		Email:       email,
		Password:    string(passwordHash),
		Address:     helper.StringToNullString(address),
		PhoneNumber: helper.StringToNullString(phoneNumber),
		Photo:       helper.StringToNullString(photo),
		DOB:         helper.StringToNullTime(dob),
		OTP:         sql.NullString{},
		Status:      "ACTIVATED",
		Auditable:   NewAuditable(createdBy),
	}
}
