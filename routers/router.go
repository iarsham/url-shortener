package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/docs"
	"github.com/iarsham/url-shortener/helpers"
)

func RedirectToDocs(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}

func SetupRouters(db *gorm.DB, rdb *redis.Client, lg *configs.CustomLogger, g *gin.Engine) {
	swaggerRoute := g.Group("")
	docs.SwaggerInfo.BasePath = "/api/"
	swaggerRoute.GET("/", RedirectToDocs)
	swaggerRoute.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authRoute := g.Group("/api/auth")
	SignUpRouter(db, rdb, lg, authRoute)
	LoginRouter(db, rdb, lg, authRoute)
	VerifyUserRouter(db, rdb, lg, authRoute)

	userRoute := g.Group("/api/user")
	userRoute.Use(helpers.JwtAuthMiddelware())
	GetUserRouter(db, rdb, lg, userRoute)
	DeleteUserRouter(db, rdb, lg, userRoute)
	PasswordRouter(db, rdb, lg, userRoute)

	linkRoute := g.Group("/api/link")
	ShortLinkRouter(db, lg, linkRoute)
	LinkRedirectRouter(db, lg, linkRoute)
}
