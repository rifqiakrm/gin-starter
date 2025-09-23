package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/errors"
	"gin-starter/common/response"
	"gin-starter/modules/auth/v1/service"
)

// UserDeleterHandler handles HTTP requests for deleting user.
type UserDeleterHandler struct {
	userUseCase service.UserDeleterUseCase
}

// NewUserDeleterHandler creates a new UserDeleterHandler.
func NewUserDeleterHandler(userUseCase service.UserDeleterUseCase) *UserDeleterHandler {
	return &UserDeleterHandler{
		userUseCase: userUseCase,
	}
}

// DeleteUserByID handles the HTTP request to delete a user by ID.
func (h *UserDeleterHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.userUseCase.DeleteUserByID(c, id); err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", nil))
}
