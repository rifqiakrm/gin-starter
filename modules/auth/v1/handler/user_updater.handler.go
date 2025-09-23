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

// UserUpdaterHandler handles HTTP requests for updating user.
type UserUpdaterHandler struct {
	userUseCase  service.UserUpdaterUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserUpdaterHandler creates a new UserUpdaterHandler.
func NewUserUpdaterHandler(userUseCase service.UserUpdaterUseCase, cloudStorage interfaces.CloudStorageUseCase) *UserUpdaterHandler {
	return &UserUpdaterHandler{
		userUseCase:  userUseCase,
		cloudStorage: cloudStorage,
	}
}

// UpdateUser handles the HTTP request to update an existing user.
func (h *UserUpdaterHandler) UpdateUser(c *gin.Context) {
	var req resource.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user := &entity.User{
		Name: req.Name,
	}

	res, err := h.userUseCase.UpdateUser(c, user)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
