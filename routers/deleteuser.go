package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func DeleteUserRouter(db *gorm.DB, rdb *redis.Client, lg *configs.CustomLogger, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db, rdb, lg)

	deleteUserController := &controllers.DeleteUserController{
		UserSerivce: userRepo,
	}

	r.DELETE("/delete-user", deleteUserController.DeleteUserHandler)
}
