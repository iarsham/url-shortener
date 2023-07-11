package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/routers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("environment variables file error!")
	}
	_, err := configs.GetDB()
	if err == nil {
		configs.Logger.Println("Mysql is Connected Successfuly!")
	}
}

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())

	routers.UrlRouter(server)
	server.Run(":8000")
}
