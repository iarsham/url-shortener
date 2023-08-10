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

// @Summary 		Create Short URL
// @Description		Create a short url that redirect to your main url
// @Tags			URL
// @Accept			json
// @Router			/link/create-short	[post]
// @Param 			request body entity.LinkRequest true "create short url required body"
// @Success			200		{object}	entity.ShortLinkOkResponse
// @Failure			400		{object}	entity.DataBodyResponse
// @Failure			400		{object}	entity.ShortLinkValidateResponse
// @Failure			409		{object}	entity.ShortLinkExistsResponse
// @Failure			500		{object}	entity.ShortLinkDBErrorResponse
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
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"response": "url already exists"})
		return
	}

	shortLink, key := s.ShortLinkService.RandomShortURL()
	newLink := models.Link{LongUrl: data.URL, ShortUrl: shortLink, Keyword: key, UserID: userID}

	if err := s.ShortLinkService.CreateShortLink(&newLink); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "cant short long url"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": newLink.ShortUrl})
}
