package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type VerifyUserController struct {
	VerifyUserService domain.VerifyUserRepository
}

// @Summary 		Verify User Account
// @Description		send verification link key in query param to active user
// @Tags			Auth
// @Accept			json
// @Router			/auth/verify-user/ [post]
// @Param key query string true "verification query param"
// @Success			200		{object}	entity.VerifyOKResponse
// @Failure			410		{object}	entity.LinkExpireResponse
// @Failure			409		{object}	entity.AlreadyVerifiedResponse
// @Failure			400		{object}	entity.DBErrorResponse
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

	user.IsActive = true
	if err := v.VerifyUserService.Save(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reponse": "failed to verify user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "user verified successfully"})
}
