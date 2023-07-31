package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/entity"
	"github.com/iarsham/url-shortener/models"
)

type SignUpController struct {
	SignUpService domain.SignUpRepository
}
// @Summary 		Register User
// @Description		register user with email and password and sending verfication email
// @Tags			Auth
// @Accept			json
// @Router			/auth/signup/ [post]
// @Param 			request body entity.Authenticate true "Register Data"
func (s *SignUpController) SignUpHandler(ctx *gin.Context) {
	var data entity.Authenticate

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if _, err := s.SignUpService.GetUserByEmail(data.Email); err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"response": "user with this email already exists"})
		return
	}

	encryptedPass, err := s.SignUpService.EncryptPassword(data.Password)

	if err != nil {
		log.Panic("error in hash user password : ", err.Error())
	}

	newUser := models.User{
		UserInfo: models.BaseUser{Email: data.Email, Password: encryptedPass},
	}

	if err := s.SignUpService.Create(&newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	go s.SignUpService.SendVerifyEmail(data.Email)
	
	accessToken := s.SignUpService.CreateAccessToken(newUser.ID.String(), data.Email)
	ctx.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}