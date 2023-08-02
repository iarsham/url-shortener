package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func ShortLinkRouter(db *gorm.DB, r *gin.RouterGroup) {
	linkRepo := services.ShortLinkRepositoryImpl(db)

	shortLinkController := &controllers.ShortLinkController{
		ShortLinkService: linkRepo,
	}

	r.POST("/create-short", shortLinkController.CreateShortLinkHandler)
}
