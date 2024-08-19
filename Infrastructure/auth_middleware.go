package infrastructure

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// create middleware for authorization...............for the user
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "error while getting the token"})
			c.Abort()
		}

		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if err := godotenv.Load("../.env"); err != nil {
				return nil, errors.New("error while loading environment variables")
			}
			return []byte(os.Getenv("KEY")), nil
		})
		// Get the payload from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "invalid token while comparing"})
			c.Abort()
			return
		}
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "error of authority"})
			c.Abort()
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}

// create the admin role middleware for the ........
func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")
		if !exist || role != "admin" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "error of authority.."})
			c.Abort()
			return
		}
		c.Next()
	}
}

// create middleware for user role middleware for the ....
func UserMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok || role != "user" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "error of authority"})
			c.Abort()
			return
		}
		c.Next()
	}
}
