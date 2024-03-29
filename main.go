package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type User struct {
	Rowid int
	Email string
}

var SecretKey string

func setupRouter() *gin.Engine {
	r := gin.Default()

	SecretKey := os.Getenv("SECRET_KEY")

	r.Use(func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		tokenString := strings.Split(header, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			value := claims["user"].(string)

			var data map[string]string
			json.Unmarshal([]byte(value), &data)

			fmt.Println(data["Email"])
		} else {
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Email module is running")
	})

	r.POST("/SendMessage", func(c *gin.Context) {
		var Items = c.Request.Body
		fmt.Println(Items)
		c.String(http.StatusOK, "Email sended")
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
