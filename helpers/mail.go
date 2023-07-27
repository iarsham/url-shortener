package helpers

import (
	"math/rand"
	"net/smtp"
	"net/url"
	"os"
)

const (
	MIME    = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateKey() string {
	key := make([]byte, 30)
	for i := range key {
		key[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	return string(key)
}

func generateVerifyLink(key string) string {
	domain := os.Getenv("DOMAIN")
	v := url.Values{}
	v.Set("key", key)
	link := domain + "/verify-email/?" + v.Encode()
	return link
}

func SendVerify(email string) (string,error) {
	var (
		from     = os.Getenv("EMAIL_FROM")
		password = os.Getenv("EMAIL_PASSWORD")
		host     = os.Getenv("EMAIL_HOST")
		port     = os.Getenv("EMAIL_PORT")
		subject  = "Verify Your Account"
	)

	key := generateKey()
	link := generateVerifyLink(key)
	auth := smtp.PlainAuth("", from, password, host)

	msg := "From: " + from + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		"Click to link and active your account:  " + link

	err := smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(msg))
	if err != nil {
		return key,err
	}	
	return key,nil
}