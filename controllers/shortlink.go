package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
	"github.com/iarsham/url-shortener/entity"
	"github.com/iarsham/url-shortener/models"
)

type ShortLinkController struct {
	ShortLinkService domain.ShortLinkRepository
}

func (s *ShortLinkController) CreateShortLinkHandler(ctx *gin.Context) {
	var data entity.LinkRequest
	userID := ctx.GetString("user_id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if ok := s.ShortLinkService.ValidateLink(data.URL); !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "url is invalid"})
		return
	}

	if ok := s.ShortLinkService.CheckLinkExists(data.URL, userID); ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "url already exists"})
		return
	}

	newLink := models.Link{LongUrl: data.URL, UserID: userID}

	if err := s.ShortLinkService.CreateShortLink(&newLink); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "cant short long url"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": newLink.ShortUrl})
}
