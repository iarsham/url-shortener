package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouters(db *gorm.DB, gin *gin.Engine) {
	authRoute := gin.Group("/api/auth")
	SignUpRouter(db, authRoute)
	LoginRouter(db, authRoute)
}
