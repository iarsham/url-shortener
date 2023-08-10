package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/controllers"
	"github.com/iarsham/url-shortener/services"
)

func LinkRedirectRouter(db *gorm.DB, r *gin.RouterGroup) {
	linkRepo := services.LinkRedirectRepositoryImpl(db)

	linkRedirectController := &controllers.LinkRedirectController{
		LinkRedirectService: linkRepo,
	}

	r.GET("/:key", linkRedirectController.LinkRedirectHandler)
}
