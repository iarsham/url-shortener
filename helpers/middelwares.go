package helpers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddelware() gin.HandlerFunc {
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

		cliams, _ := token.Claims.(jwt.MapClaims)

		ctx.Set("user_id", cliams["user_id"])
		ctx.Set("email", cliams["email"])
		ctx.Next()
	}
}
