package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/errors"
	"gin-starter/common/response"
	"gin-starter/modules/auth/v1/service"
)

// UserFinderHandler handles HTTP requests for retrieving user.
type UserFinderHandler struct {
	userUseCase service.UserFinderUseCase
}

// NewUserFinderHandler creates a new UserFinderHandler.
func NewUserFinderHandler(userUseCase service.UserFinderUseCase) *UserFinderHandler {
	return &UserFinderHandler{
		userUseCase: userUseCase,
	}
}

// GetUserByID handles the HTTP request to retrieve a user by ID.
func (h *UserFinderHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.userUseCase.GetUserByID(c, id)
	if err != nil {
		parseErr := errors.ParseError(err)
		c.JSON(parseErr.Code, response.ErrorAPIResponse(parseErr.Code, parseErr.Message))
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponse(http.StatusOK, "success", res))
}
