package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/configs"
	"github.com/iarsham/url-shortener/entity"
	"github.com/iarsham/url-shortener/helpers"
	"github.com/iarsham/url-shortener/models"
)

func RegisterHandler(ctx *gin.Context) {
	var data entity.Authenticate
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	passWord, _ := helpers.Hash(data.Password)
	newUser := models.User{
		UserInfo: models.BaseUser{Email: data.Email, Password: passWord},
	}
	configs.DB.Create(&newUser)
	ctx.JSON(http.StatusCreated, gin.H{"response": "User created"})
}

func LoginHandler(ctx *gin.Context) {
	var data entity.Authenticate
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	var dbUser models.User
	result := configs.DB.Where("email = ?", data.Email).First(&dbUser)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "user not found with this email"})
		return
	}

	if _, err := helpers.VerifyHash(dbUser.UserInfo.Password, data.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "password is incorrect"})
		return
	}

	access_token := helpers.GenerateJWT(dbUser.ID.String(), data.Email)
	ctx.JSON(http.StatusOK, gin.H{"access_token": access_token})
}
