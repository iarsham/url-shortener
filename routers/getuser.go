package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func GetUserRouter(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db)

	GetUserController := &controllers.GetUserController{
		UserService: userRepo,
	}

	r.GET("/me", GetUserController.GetUserHandler)
}
