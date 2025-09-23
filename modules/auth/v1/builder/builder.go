package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	cronjobHandler "gin-starter/modules/auth/v1/cronjob/handler"
	userRepo "gin-starter/modules/auth/v1/repository"
	"gin-starter/modules/auth/v1/service"
)

// BuildUserHandler builds auth handler
// starting from handler down to repository or tool.
func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, cache interfaces.Cacheable, cloudStorage interfaces.CloudStorageUseCase) {
	// Repository
	ar := userRepo.NewAuthRepository(db)
	urf := userRepo.NewUserFinderRepository(db, cache)
	urc := userRepo.NewUserCreatorRepository(db, cache)
	uru := userRepo.NewUserUpdaterRepository(db, cache)
	urd := userRepo.NewUserDeleterRepository(db, cache)

	// Service
	uc := service.NewUserCreator(cfg, urc)
	uf := service.NewUserFinder(cfg, urf)
	uu := service.NewUserUpdater(cfg, uru)
	ud := service.NewUserDeleter(cfg, urd)
	ac := service.NewAuthService(cfg, ar)

	// Handler
	app.UserAuthHTTPHandler(cfg, router, ac)
	app.UserFinderHTTPHandler(cfg, router, uf)
	app.UserCreatorHTTPHandler(cfg, router, uc, uf, cloudStorage)
	app.UserUpdaterHTTPHandler(cfg, router, uu, cloudStorage)
	app.UserDeleterHTTPHandler(cfg, router, ud)
}

// BuildExampleCronjobHandler is used to build the cron handler.
func BuildExampleCronjobHandler(cfg config.Config, db *gorm.DB, cache interfaces.Cacheable) *cronjobHandler.ExampleCronjob {
	// Repository
	urf := userRepo.NewUserFinderRepository(db, cache)

	// Service
	svc := service.NewUserFinder(cfg, urf)

	return cronjobHandler.NewExampleCronjobHandler(svc)
}
