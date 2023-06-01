package utils

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost := GetEnv("HASH_COST")
	costInt, err := strconv.Atoi(cost)
	if err != nil {
		log.Fatal("Not convertable")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costInt)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
