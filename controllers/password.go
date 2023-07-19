package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/entity"
)

type PasswordController struct {
	PasswordService domain.PasswordRepository
}

func (p *PasswordController) PasswordChangeHandler(ctx *gin.Context) {
	var data entity.PasswordRequest
	user_id := ctx.GetString("user_id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	user, err := p.PasswordService.GetUserByID(user_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": "user not found"})
		return
	}

	if ok, _ := user.VerifyUserPassword(data.CurrentPassword); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"response": "current password is incorrect"})
		return
	}

	if data.Password != data.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": "new passwords must be equal"})
		return
	}

	newPass, _ := p.PasswordService.EncryptPassword(data.Password)
	user.UserInfo.Password = newPass

	if err := p.PasswordService.Save(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response": "password change failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "password changed successfully"})
}
