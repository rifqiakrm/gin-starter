// Package middleware provides authentication middleware
package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	commonCache "gin-starter/common/cache"
	"gin-starter/common/helper"
	"gin-starter/common/interfaces"
	"gin-starter/common/response"
	"gin-starter/config"
)

const (
	// ContextUserIDKey is a constant to store user ID
	ContextUserIDKey = "userID"
	// ContextPermissionsKey is a constant to store user permissions
	ContextPermissionsKey = "permissions"
)

// Auth returns a middleware that validates JWT, checks Redis for token validity,
// and fetches user permissions.
func Auth(cfg config.Config, cache interfaces.Cacheable) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "missing or invalid Authorization header"))
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Decode JWT
		claims, err := helper.JWTDecode(cfg, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "invalid token: "+err.Error()))
			c.Abort()
			return
		}

		// Check if token is valid in Redis (session store)
		tokenKey := fmt.Sprintf(commonCache.AuthToken, tokenString)
		bytes, _ := cache.Get(tokenKey)
		if bytes == nil {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "unauthorized: token not found in redis"))
			c.Abort()
			return
		}

		// Fetch user permissions from Redis
		permKey := fmt.Sprintf(commonCache.UserPermissionByUserID, claims.Subject) // e.g. "user:perm:%s"
		permBytes, _ := cache.Get(permKey)

		var permissions []string
		if permBytes != nil {
			_ = json.Unmarshal(bytes, &permissions)
		}

		// Store userID (UUID) and permissions in context
		c.Set(ContextUserIDKey, claims.Subject)
		c.Set(ContextPermissionsKey, permissions)

		c.Next()
	}
}

// GetUserID extracts userID (UUID string) from gin.Context
func GetUserID(c *gin.Context) (string, bool) {
	val, exists := c.Get(ContextUserIDKey)
	if !exists {
		return "", false
	}
	userID, ok := val.(string)
	return userID, ok
}

// GetPermissions extracts permissions slice from gin.Context
func GetPermissions(c *gin.Context) ([]string, bool) {
	val, exists := c.Get(ContextPermissionsKey)
	if !exists {
		return nil, false
	}
	perms, ok := val.([]string)
	return perms, ok
}
