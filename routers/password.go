package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func PasswordRouter(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db)

	passController := &controllers.PasswordController{
		PasswordService: services.PasswordServiceImpl(userRepo),
	}

	r.POST("/change-password", passController.PasswordChangeHandler)
}
