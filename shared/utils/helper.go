package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	id := uuid.Must(uuid.NewRandom())
	return id.String()
}

func HashPassword(password string) (result string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func CheckPasswordHash(password, hash string) error {
	fmt.Println(password)
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckImageType(url string) bool {
	url = url[strings.LastIndex(url, ".")+1:]
	if url == "png" || url == "jpg" || url == "jpeg" {
		return true
	}
	return false
}

func IsValidPhoneNumber(phoneNumber string) error {
	phoneNumberRegex := `^\+[0-9]{1,4}-?[0-9]{1,15}$`
	match, _ := regexp.MatchString(phoneNumberRegex, phoneNumber)
	if match {
		return nil
	}

	return errors.New("format Phone tidak valid")
}
