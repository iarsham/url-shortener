package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func LoginRouter(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db)

	loginController := &controllers.LoginController{
		LoginService: services.LoginServiceImpl(userRepo),
	}

	r.POST("/login", loginController.LoginHandler)
}
