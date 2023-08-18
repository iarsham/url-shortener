package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func LinkRedirectRouter(db *gorm.DB, lg *configs.CustomLogger, r *gin.RouterGroup) {
	linkRepo := services.LinkRedirectRepositoryImpl(db, lg)

	linkRedirectController := &controllers.LinkRedirectController{
		LinkRedirectService: linkRepo,
	}

	r.GET("/:key", linkRedirectController.LinkRedirectHandler)
}
