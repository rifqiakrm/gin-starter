package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/errors"
	"gin-starter/common/response"
	"gin-starter/modules/auth/v1/service"
	"gin-starter/resource"
)

// AuthHandler is a handler for auth
type AuthHandler struct {
	authUseCase service.AuthUseCase
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(
	authUseCase service.AuthUseCase,
) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Login is a handler for login
func (ah *AuthHandler) Login(c *gin.Context) {
	var request resource.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res, err := ah.authUseCase.AuthValidate(c, request.Email, request.Password)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	token, err := ah.authUseCase.GenerateAccessToken(c, res)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	otpIsNull := false

	if res.OTP.String != "" {
		otpIsNull = true
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", resource.NewLoginResponse(token.Token, otpIsNull)))
}
