package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 2
	bytePass := []byte(pass)

	result, err := bcrypt.GenerateFromPassword(bytePass, salt)

	if err != nil {
		panic(err.Error)
	}

	return string(result)
}

func ComparePass(hp, p string) bool {
	hashedPass, pass := []byte(hp), []byte(p)
	err := bcrypt.CompareHashAndPassword(hashedPass, pass)

	return err == nil
}
