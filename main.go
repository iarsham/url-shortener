package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iarsham/url-shortener/configs"
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
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "Hello World"})
	})
	r.Run(":8000")
}
