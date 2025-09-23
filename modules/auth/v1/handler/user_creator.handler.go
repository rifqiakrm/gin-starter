package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/common/response"
	"gin-starter/entity"
	"gin-starter/modules/auth/v1/service"
	"gin-starter/resource"
)

// UserCreatorHandler handles HTTP requests for creating user.
type UserCreatorHandler struct {
	userUseCase  service.UserCreatorUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserCreatorHandler creates a new UserCreatorHandler.
func NewUserCreatorHandler(userUseCase service.UserCreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) *UserCreatorHandler {
	return &UserCreatorHandler{
		userUseCase:  userUseCase,
		cloudStorage: cloudStorage,
	}
}

// CreateUser handles the HTTP request to create a new user.
func (h *UserCreatorHandler) CreateUser(c *gin.Context) {
	var req resource.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user := &entity.User{
		Name: req.Name,
	}

	res, err := h.userUseCase.CreateUser(c, user)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
