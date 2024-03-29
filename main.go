package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SecretKey string

func setupRouter() *gin.Engine {
	r := gin.Default()

	SecretKey := os.Getenv("SECRET_KEY")

	r.Use(func(c *gin.Context) {

		fmt.Println("Middleware - ValidateToken")

		header := c.GetHeader("Authorization")

		tokenString := strings.Split(header, " ")[1]
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["userID"].(string)
			fmt.Println("User ID:", userID)
		}

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Email module is running")
	})

	return r
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	port := os.Getenv("PORT")

	r := setupRouter()
	r.Run(port)
}
