package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/controllers"
)

func UrlRouter(r *gin.Engine) {
	userRoutes := r.Group("/api/user/")

	userRoutes.POST("register/", controllers.RegisterHandler)
	userRoutes.POST("login/", controllers.LoginHandler)

}
