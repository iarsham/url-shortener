package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iarsham/url-shortener/domain"
)

type LinkRedirectController struct {
	LinkRedirectService domain.LinkRedirectRepository
}

func (l *LinkRedirectController) LinkRedirectHandler(ctx *gin.Context) {
	key := ctx.Param("url")
	shortURL := l.LinkRedirectService.GetCurrentHost(ctx) + key

	dbLink, ok := l.LinkRedirectService.CheckLinkExists(shortURL)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "url not found"})
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, dbLink.LongUrl)
}
