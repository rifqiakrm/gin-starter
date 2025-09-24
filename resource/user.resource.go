package resource

import (
	"mime/multipart"
	"os"

	"gin-starter/common/helper"
	"gin-starter/entity"
)

// timeFormat defines the standard layout used for formatting dates and times
// in responses throughout the resource package.
const (
	timeFormat = "2006-01-02 15:04:05"
)

//
// Request Structs
//

// CreateUserRequest represents the payload required to create a new auth.
// It includes basic auth details and an uploaded profile photo.
type CreateUserRequest struct {
	Name        string                `form:"name" json:"name" binding:"required"`
	Email       string                `form:"email" json:"email" binding:"required"`
	Password    string                `form:"password" json:"password" binding:"required"`
	DOB         string                `form:"dob" json:"dob" binding:"required"`
	PhoneNumber string                `form:"phone_number" json:"phone_number" binding:"required"`
	Photo       *multipart.FileHeader `form:"photo" json:"photo" binding:"required"`
}

//
// Request Structs
//

// UpdateUserRequest represents the payload for updating a auth's profile.
// Fields are optional, allowing partial updates of auth information.
type UpdateUserRequest struct {
	ID          string                `form:"id" json:"id"`
	Name        string                `form:"name" json:"name"`
	Email       string                `form:"email" json:"email"`
	DOB         string                `form:"dob" json:"dob"`
	PhoneNumber string                `form:"phone_number" json:"phone_number"`
	Photo       *multipart.FileHeader `form:"photo" json:"photo"`
}

// ChangePasswordRequest represents the payload for changing a auth's password.
// It requires the old password, new password, and confirmation of the new password.
type ChangePasswordRequest struct {
	OldPassword             string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
}

// ForgotPasswordRequest represents the payload used to initiate a
// password reset process via the auth's registered email.
type ForgotPasswordRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

// ForgotPasswordChangeRequest represents the payload used to update a
// password after initiating the reset process with a valid token.
type ForgotPasswordChangeRequest struct {
	Token                   string `form:"token" json:"token" binding:"required"`
	NewPassword             string `form:"new_password" json:"new_password" binding:"required"`
	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" binding:"required"`
}

// GetUserByForgotPasswordTokenRequest represents the request for retrieving
// auth details using a forgot-password token.
type GetUserByForgotPasswordTokenRequest struct {
	Token string `uri:"token" json:"token" binding:"required"`
}

// VerifyOTPRequest represents the payload required to verify an OTP (One-Time Password).
type VerifyOTPRequest struct {
	Code string `form:"code" json:"code" binding:"required"`
}

//
// Response Structs
//

// UserProfile represents the response model for a auth's profile information.
// It includes identity, contact details, status, and timestamps.
type UserProfile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	OTPIsNull   bool   `json:"otp_is_null"`
	PhoneNumber string `json:"phone_number"`
	DOB         string `json:"dob"`
	Status      string `json:"status"`
	Photo       string `json:"photo"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

//
// Factory Functions
//

// NewUserProfile creates a new UserProfile response object from an entity.User.
// It maps fields from the entity model, formats timestamps, and builds
// a fully qualified image path for the auth's photo.
func NewUserProfile(user *entity.User) *UserProfile {
	otpIsNull := false
	if user.OTP.String != "" {
		otpIsNull = true
	}

	dob := "1970-01-01"
	if user.DOB.Valid {
		dob = user.DOB.Time.Format(timeFormat)
	}

	return &UserProfile{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		DOB:         dob,
		Photo:       helper.ImageFullPath(os.Getenv("IMAGE_HOST"), user.Photo.String),
		Status:      user.Status,
		OTPIsNull:   otpIsNull,
		CreatedAt:   user.CreatedAt.Format(timeFormat),
		UpdatedAt:   user.UpdatedAt.Format(timeFormat),
	}
}
