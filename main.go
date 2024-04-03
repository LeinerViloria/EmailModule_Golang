package main

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
	// "strings"
	//"crypto/tls"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/go-mail/mail"
)

var UserEmail string
var SecretKey string

func setupRouter() *gin.Engine {
	r := gin.Default()

	// SecretKey := os.Getenv("SECRET_KEY")

	r.Use(func(c *gin.Context) {
		// header := c.GetHeader("Authorization")

		// tokenString := strings.Split(header, " ")[1]

		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(SecretKey), nil
		// })

		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		// 	c.Abort()
		// 	return
		// }

		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	value := claims["user"].(string)

		// 	var data map[string]string
		// 	json.Unmarshal([]byte(value), &data)

		// 	UserEmail = data["Email"]
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Email module is running")
	})

	r.POST("/SendMessage", func(c *gin.Context) {

		from := "t84098030@gmail.com"

		m := mail.NewMessage()

		m.SetHeader("From", from)
		m.SetHeader("To", "pokop58224@evimzo.com")
		m.SetHeader("Subject", "Hello!")

		m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

		d := mail.NewDialer("smtp.gmail.com", 465, from, "OtraPruebaCecar123")

		if err := d.DialAndSend(m); err != nil {

		panic(err)

		}

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
