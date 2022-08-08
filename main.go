package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"openapi-eds-verification/helpers"
	"os"
)

func main() {
	_ = godotenv.Load()
	var err error

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := helpers.GetEnvDefault("PORT", "3000")

	if err = r.Run(":" + port); err != nil {
		os.Exit(1)
	}
}
