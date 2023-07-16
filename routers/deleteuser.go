package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func DeleteUserRouter(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db)

	deleteUserController := &controllers.DeleteUserController{
		UserSerivce: userRepo,
	}

	r.DELETE("/delete-user", deleteUserController.DeleteUserHandler)
}
