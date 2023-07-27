package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/helpers"
)

func SetupRouters(db *gorm.DB, rdb *redis.Client, gin *gin.Engine) {
	authRoute := gin.Group("/api/auth")
	SignUpRouter(db, rdb, authRoute)
	LoginRouter(db, rdb, authRoute)
	VerifyUserRouter(db, rdb, authRoute)

	userRoute := gin.Group("/api/user")
	userRoute.Use(helpers.JwtAuthMiddelware())
	GetUserRouter(db, rdb, userRoute)
	DeleteUserRouter(db, rdb, userRoute)
	PasswordRouter(db, rdb, userRoute)
}
