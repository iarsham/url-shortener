package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type DeleteUserController struct {
	UserSerivce domain.UserRepository
}

func (d *DeleteUserController) DeleteUserHandler(ctx *gin.Context) {
	userID := ctx.GetString("user_id")

	user, err := d.UserSerivce.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": "user not found"})
		return
	}

	if err := d.UserSerivce.Delete(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"response": "user deleted"})
}
