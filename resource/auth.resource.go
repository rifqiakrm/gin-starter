// Package resource defines request and response models
package resource

// LoginRequest represents the payload for a login request.
// It requires the auth's email and password.
type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginResponse represents the response returned after a successful login.
// It contains the authentication token and whether the OTP is null.
type LoginResponse struct {
	Token     string `json:"token"`
	OTPIsNull bool   `json:"otp_is_null"`
}

// NewLoginResponse creates a new LoginResponse with the provided
// token and OTP state.
func NewLoginResponse(token string, otpIsNull bool) *LoginResponse {
	return &LoginResponse{Token: token, OTPIsNull: otpIsNull}
}

// RegisterRequest represents the payload required to register a new auth.
// It includes personal details such as name, email, password, gender,
// birthday, birthplace, and phone number.
type RegisterRequest struct {
	Name       string `form:"name,omitempty" json:"name,omitempty" binding:"required"`
	Email      string `form:"email,omitempty" json:"email,omitempty" binding:"required"`
	Password   string `form:"password,omitempty" json:"password,omitempty" binding:"required"`
	Gender     int    `form:"gender,omitempty" json:"gender,omitempty" binding:"required"`
	Birthday   string `form:"birthday,omitempty" json:"birthday,omitempty" binding:"required"`
	Birthplace string `form:"birthplace,omitempty" json:"birthplace,omitempty" binding:"required"`
	Phone      string `form:"phone,omitempty" json:"phone,omitempty" binding:"required"`
}
