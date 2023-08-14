package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
	// "github.com/iarsham/url-shortener/entity"
)

type GetUserController struct {
	UserService domain.UserRepository
}

// @Summary 		Get User
// @Description		Get user data information
// @Tags			User
// @Accept			json
// @Router			/user/me/ [get]
// @Success			200		{object}	entity.UserResponse
// @Failure			404		{object}	entity.User404Responsse
func (g *GetUserController) GetUserHandler(ctx *gin.Context) {
	userID := ctx.GetString("user_id")

	user, err := g.UserService.GetUserWithLinks(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
