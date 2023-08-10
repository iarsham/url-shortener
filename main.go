package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/routers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading .env file")
	}
	configs.GetDB()
	configs.GetRedis()
}

// @title			UrlShortener Swagger Document
// @version			1.0
// @termsOfService  http://swagger.io/terms/
// @contact.email   arshamdev2001@gmail.com
// @BasePath	 	/api/
// @host			localhost:8000
// @licence.Name	MIT
// @licence.url     https://www.mit.edu/~amini/LICENSE.md
// @schemes			http https
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @accept json
func main() {
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	server.Use(cors.Default())
	routers.SetupRouters(configs.DB, configs.GetRedis(), server)
	server.Run(":8000")
}
