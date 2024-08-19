package infrastructure

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create a method for generating a token for authentication and authorization
func GenerateToken(user_id primitive.ObjectID, role string) (string, error) {
	//load the key from the .env var
	err := godotenv.Load("../.env")
	if err != nil {
		return "", errors.New("error while loading")
	}
	key := os.Getenv("KEY")
	//generate the token based on the key
	claims := jwt.MapClaims{
		"user_id": user_id,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString, nil
}

// create a method for verifying the token from the contex and ...........
func VerifyToken(tokenString string) (*jwt.Token, error) {
	//decode the token by using the key from the env....
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		err := godotenv.Load("../.env")
		if err != nil {
			return nil, errors.New("error while loading")
		}
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("error on the validty")
	}
	return token, nil

}
