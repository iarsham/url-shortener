package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/services"
)

func ShortLinkRouter(db *gorm.DB, lg *configs.CustomLogger, r *gin.RouterGroup) {
	linkRepo := services.ShortLinkRepositoryImpl(db, lg)

	shortLinkController := &controllers.ShortLinkController{
		ShortLinkService: linkRepo,
	}

	r.POST("/create-short", helpers.JwtAuthMiddelware(), shortLinkController.CreateShortLinkHandler)
}
