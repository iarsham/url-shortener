package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/docs"
	"github.com/iarsham/url-shortener/helpers"
)

func RedirectToDocs(ctx *gin.Context) {
	fmt.Println(ctx.Request.Host)
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}

func SetupRouters(db *gorm.DB, rdb *redis.Client, gin *gin.Engine) {
	swaggerRoute := gin.Group("")
	docs.SwaggerInfo.BasePath = "/api/"
	swaggerRoute.GET("/", RedirectToDocs)
	swaggerRoute.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authRoute := gin.Group("/api/auth")
	SignUpRouter(db, rdb, authRoute)
	LoginRouter(db, rdb, authRoute)
	VerifyUserRouter(db, rdb, authRoute)

	userRoute := gin.Group("/api/user")
	userRoute.Use(helpers.JwtAuthMiddelware())
	GetUserRouter(db, rdb, userRoute)
	DeleteUserRouter(db, rdb, userRoute)
	PasswordRouter(db, rdb, userRoute)

	linkRoute := gin.Group("/api/link")
	ShortLinkRouter(db, linkRoute)
	LinkRedirectRouter(db, linkRoute)
}
