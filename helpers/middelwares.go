package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iarsham/url-shortener/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "authenticate required"})
			return
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "authorization foramt is not correct"})
			return
		}

		token, err := jwt.Parse(authToken[1], func(token *jwt.Token) (interface{}, error) { return []byte(os.Getenv("SECRET_KEY")), nil })
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "token is invalid or expired"})
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		ctx.Set("user_id", claims["user_id"])
		ctx.Set("email", claims["email"])
		ctx.Next()
	}
}

func QueryParamMiddleware(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryParam := ctx.Query(param)

		if queryParam == "" {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"response": fmt.Sprintf("missing ( ?%s= ) query parameter", param)})
			return
		}

		ctx.Next()
	}
}

func RateLimitMiddleware(rdb *redis.Client, db *gorm.DB, limit, burst int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		db.Where("email = ?", ctx.GetString("email")).First(&user)
		clientIP := ctx.ClientIP()
		rateCount, err := rdb.Get(ctx, clientIP).Int()

		if err != nil && err != redis.Nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Internal server error"})
			return
		}

		if !user.IsActive {
			rdb.Incr(ctx, clientIP)
			rdb.Expire(ctx, clientIP, time.Duration(limit)*time.Hour)
		}

		if !user.IsActive && rateCount >= burst {
			seconds := rdb.TTL(ctx, clientIP).Val() / 1000000000
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"response": "To many requests", "try after": fmt.Sprintf("%d seconds", seconds)})
			return
		}

		ctx.Next()
	}
}
