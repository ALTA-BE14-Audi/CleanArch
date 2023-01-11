package helper

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return "", errors.New("password process error")
	}

	return string(hashed), nil
}
