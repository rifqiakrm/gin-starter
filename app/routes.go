package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/common/interfaces"
	"gin-starter/common/response"
	"gin-starter/config"
	userhandlerv1 "gin-starter/modules/auth/v1/handler"
	userservicev1 "gin-starter/modules/auth/v1/service"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(_ config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

// UserAuthHTTPHandler is a handler for auth APIs
func UserAuthHTTPHandler(_ config.Config, router *gin.Engine, auc userservicev1.AuthUseCase) {
	hnd := userhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
	}
}

// UserFinderHTTPHandler is a handler for auth APIs
func UserFinderHTTPHandler(_ config.Config, router *gin.Engine, cf userservicev1.UserFinderUseCase) {
	_ = userhandlerv1.NewUserFinderHandler(cf)
	_ = router.Group("/v1")
	{
		_ = "Todo 1"
		// TODO: Add route
	}
}

// UserCreatorHTTPHandler is a handler for auth APIs
func UserCreatorHTTPHandler(_ config.Config, router *gin.Engine, uc userservicev1.UserCreatorUseCase, _ userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	_ = userhandlerv1.NewUserCreatorHandler(uc, cloudStorage)
	_ = router.Group("/v1")
	{
		_ = "Todo 2"
		// TODO: Add route
	}
}

// UserUpdaterHTTPHandler is a handler for auth APIs
func UserUpdaterHTTPHandler(_ config.Config, router *gin.Engine, uu userservicev1.UserUpdaterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	_ = userhandlerv1.NewUserUpdaterHandler(uu, cloudStorage)
	_ = router.Group("/v1")
	{
		_ = "Todo 3"
		// TODO: Add route
	}
}

// UserDeleterHTTPHandler is a handler for auth APIs
func UserDeleterHTTPHandler(_ config.Config, router *gin.Engine, ud userservicev1.UserDeleterUseCase) {
	_ = userhandlerv1.NewUserDeleterHandler(ud)
	_ = router.Group("/v1")
}
