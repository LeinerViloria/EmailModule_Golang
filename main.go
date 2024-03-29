package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		parts := strings.Split(header, " ")
		fmt.Println("Middleware - ValidateToken")
		fmt.Println(parts[1])

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
