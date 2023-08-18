package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func GetUserRouter(db *gorm.DB, rdb *redis.Client, lg *configs.CustomLogger, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db, rdb, lg)

	GetUserController := &controllers.GetUserController{
		UserService: userRepo,
	}

	r.GET("/me", GetUserController.GetUserHandler)
}
