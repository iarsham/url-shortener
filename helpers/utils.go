package helpers

import (
	"time"

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