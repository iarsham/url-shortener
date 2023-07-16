package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/entity"
)

type LoginController struct {
	LoginService domain.LoginRepository
}

func (l *LoginController) LoginHandler(ctx *gin.Context) {
	var data entity.Authenticate

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	user, err := l.LoginService.GetUserByEmail(data.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": "user not found"})
		return
	}

	if ok, _ := user.VerifyUserPassword(data.Password); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"response": "password is incorrect"})
		return
	}

	accessToken := l.LoginService.CreateAccessToken(user.ID.String(), data.Email)
	ctx.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}
