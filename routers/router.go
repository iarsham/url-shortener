package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/helpers"
)

func SetupRouters(db *gorm.DB, gin *gin.Engine) {
	authRoute := gin.Group("/api/auth")
	SignUpRouter(db, authRoute)
	LoginRouter(db, authRoute)

	userRoute := gin.Group("/api/user")
	userRoute.Use(helpers.JwtAuthMiddelware())
	GetUserRouter(db, userRoute)
	DeleteUserRouter(db, userRoute)
	PasswordRouter(db, userRoute)
}
