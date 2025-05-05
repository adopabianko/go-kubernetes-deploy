package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin ke mode release
	gin.SetMode(gin.ReleaseMode)

	// Buat router Gin default
	router := gin.Default()

	// Route untuk halaman root
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World dari Kubernetes dengan Go dan Gin!",
		})
	})

	// Route untuk health check
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Ambil port dari environment variable atau gunakan 8080 sebagai default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Jalankan server
	log.Println("Starting server on port " + port)
	router.Run(":" + port)
}
