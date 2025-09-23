package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/response"
)

// RequirePermission enforces a specific permission
func RequirePermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		perms, ok := GetPermissions(c)
		if !ok || !hasPermission(perms, required) {
			c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "forbidden: insufficient permission"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// hasPermission breaks and match permission
func hasPermission(perms []string, required string) bool {
	for _, p := range perms {
		if p == required {
			return true
		}
	}
	return false
}
