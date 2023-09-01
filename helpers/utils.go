package helpers

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

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

func GenerateJWT(userId, email string) (string, error) {
	ttl, _ := strconv.Atoi(os.Getenv("EXPIRE_JWT_TOKEN_Minutes"))
	cliams := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Minute * time.Duration(ttl)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cliams)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsValidURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	if !strings.Contains(parsedURL.String(), "://") {
		return false
	}
	if domain := os.Getenv("FRONTEND_DOMAIN"); !strings.HasSuffix(domain, "/") {
		os.Setenv("FRONTEND_DOMAIN", domain+"/")
	}
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

func MakeShortURL() (string, string) {
	domain := os.Getenv("FRONTEND_DOMAIN")
	key := generateKey(7)
	if !IsValidURL(domain) {
		panic("domain in env is not valid")
	}
	return domain + key, key
}

func CustomShortURL(key string) string {
	domain := os.Getenv("FRONTEND_DOMAIN")
	if !IsValidURL(domain) {
		panic("domain in env is not valid")
	}
	return domain + key
}
