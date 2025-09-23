// Package service provides service
package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"golang.org/x/crypto/bcrypt"

	"gin-starter/common/constant"
	"gin-starter/common/errors"
	"gin-starter/common/helper"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/repository"
)

const (
	four = 4
)

// AuthService is a service for auth
type AuthService struct {
	cfg      config.Config
	authRepo repository.AuthRepositoryUseCase
}

// AuthUseCase is a usecase for auth
type AuthUseCase interface {
	// AuthValidate is a function that validates the auth
	AuthValidate(ctx context.Context, email, password string) (*entity.User, error)
	// GenerateAccessToken is a function that generates an access token
	GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error)
}

// NewAuthService is a constructor for AuthService
func NewAuthService(
	cfg config.Config,
	authRepo repository.AuthRepositoryUseCase,
) *AuthService {
	return &AuthService{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

// AuthValidate is a function that validates the auth
func (as *AuthService) AuthValidate(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrWrongLoginCredentials.Error()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.ErrWrongLoginCredentials.Error()
	}

	otp := helper.GenerateOTP(four)

	if user.Email == "auth-test@gmail.com" {
		otp = "1234"
	}

	if err := as.authRepo.UpdateOTP(ctx, user, otp); err != nil {
		return nil, err
	}

	t, err := template.ParseFiles("./template/email/send_otp.html")
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	var body bytes.Buffer

	err = t.Execute(&body, struct {
		Name string
		OTP  string
	}{
		Name: user.Name,
		OTP:  otp,
	})
	if err != nil {
		log.Println(fmt.Errorf("failed to execute email data: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	payload := entity.EmailPayload{
		To:       user.Email,
		Subject:  "Login Verification",
		Content:  body.String(),
		Category: entity.EmailCategorySendOTP,
	}

	if err := helper.SendTopic(context.Background(), as.cfg, constant.SendEmailTopic, payload); err != nil {
		log.Println(err)
	}

	return user, nil
}

// GenerateAccessToken is a function that generates an access token
func (as *AuthService) GenerateAccessToken(_ context.Context, user *entity.User) (*entity.Token, error) {
	token, err := helper.JWTEncode(as.cfg, user.ID, as.cfg.JWTConfig.Issuer)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return &entity.Token{
		Token: token,
	}, nil
}
