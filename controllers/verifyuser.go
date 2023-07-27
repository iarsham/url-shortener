package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type VerifyUserController struct {
	VerifyUserService domain.VerifyUserRepository
}

func (v *VerifyUserController) VerifyUserHandler(ctx *gin.Context) {
	keyParam := ctx.Query("key")
	user, err := v.VerifyUserService.GetUserFromCache(keyParam)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusGone, gin.H{"reponse": "link is invalid or expired"})
		return
	}

	if ok := v.VerifyUserService.CheckUserStatus(&user); ok {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"reponse": "user already verified"})
		return
	}

	user.UserInfo.IsActive = true

	if err := v.VerifyUserService.Save(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reponse": "failed to verify user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "user verified successfully"})
}
