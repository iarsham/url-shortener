package helpers

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var CurrentTime = func() time.Time {
	return time.Now()
}

func Hash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(bytes), err
}

func VerifyHash(hash string, plain string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateJWT(userId, email string) (tokenString string) {
	ttl, _ := strconv.Atoi(os.Getenv("EXPIRE_JWT_TOKEN_Minutes"))
	cliams := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Minute * time.Duration(ttl)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cliams)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatal("generate jwt failed :", err)
	}
	return
}

func IsValidURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

func CurrentHost(ctx *gin.Context) string {
	schema := "http://"
	if ctx.Request.TLS != nil {
		schema = "https://"
	}
	host := ctx.Request.Host
	return schema + host + "/"
}

func MakeShortURL(ctx *gin.Context) string {
	return CurrentHost(ctx) + generateKey(7)
}
