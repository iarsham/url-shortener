package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type LinkRedirectController struct {
	LinkRedirectService domain.LinkRedirectRepository
}

// @Summary 		Redirect URL
// @Description		Redirect short url to your main url
// @Tags			URL
// @Accept			json
// @Router			/link/{key}	[get]
// @Param 			key		path	string 		true 	"url key"
// @Success			200		{object}	entity.ShortLinkOkResponse
// @Failure			404		{object}	entity.ShortLinkNotExistsResponse
func (l *LinkRedirectController) LinkRedirectHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	dbLink, ok := l.LinkRedirectService.CheckLinkExists(key)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "url not found"})
		return
	}
	dbLink.ViewCount += 1
	l.LinkRedirectService.Save(&dbLink)

	ctx.JSON(http.StatusOK, gin.H{"response": dbLink.LongUrl, "views": dbLink.ViewCount})
}
