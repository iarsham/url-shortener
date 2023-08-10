package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type DeleteUserController struct {
	UserSerivce domain.UserRepository
}

// @Summary 		Delete User
// @Description		Delete user record from db
// @Tags			User
// @Accept			json
// @Router			/user/delete-user/ [delete]
// @Success			204		{string}	string
// @Failure			404		{object}	entity.User404Responsse
// @Failure			500		{object}	entity.DBErrorResponse
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
