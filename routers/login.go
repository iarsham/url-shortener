package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func LoginRouter(db *gorm.DB, rdb *redis.Client, lg *configs.CustomLogger, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db, rdb, lg)

	loginController := &controllers.LoginController{
		LoginService: services.LoginRepositoryImpl(userRepo),
	}

	r.POST("/login", loginController.LoginHandler)
}
