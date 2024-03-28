package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

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
