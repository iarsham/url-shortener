package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func SignUpRouter(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := services.UserRepositoryImpl(db)

	signUpContoller := &controllers.SignUpController{
		SignUpService: services.SignUpServiceImpl(userRepo),
	}

	r.POST("/signup", signUpContoller.SignUpHandler)
}
