package infrastructure

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// create a method for hashing the password from the original
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error while hashing the password")
	}
	hashedPassword := string(hashed)
	return string(hashedPassword), nil

}

// create a method comparing hashed password with normal password
func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
