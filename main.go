package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/routers"
)

func init() {
	if gin.Mode() != gin.ReleaseMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error in loading .env file")
		}
	}
	lg := configs.NewLogger()
	configs.GetDB(lg)
	configs.GetRedis(lg)
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
	lg := configs.NewLogger()
	routers.SetupRouters(configs.DB, configs.RDB, lg, server)
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				helpers.RemoveExpiredUrls(configs.DB)
			}
		}
	}()
	if err := server.Run(":8000"); err != nil {
		lg.Logger.Fatal(err.Error())
	}
}
