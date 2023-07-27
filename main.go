package main

import (
	"log"

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

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	routers.SetupRouters(configs.DB, configs.GetRedis(), server)
	server.Run(":8000")
}
