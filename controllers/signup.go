package controllers

import (
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
// @Success			201		{object}	entity.LoginSignUpOkResponse
// @Failure			400		{object}	entity.DataBodyResponse
// @Failure			409		{object}	entity.UserExistResponse
// @Failure			500		{object}	entity.DBErrorResponse
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

	newUser := models.User{Email: data.Email, Password: s.SignUpService.EncryptPassword(data.Password)}
	if err := s.SignUpService.Create(&newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response": "craete user failed"})
		return
	}

	go s.SignUpService.SendVerifyEmail(data.Email)

	accessToken := s.SignUpService.CreateAccessToken(newUser.ID.String(), data.Email)
	ctx.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}
