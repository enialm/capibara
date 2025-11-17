package main

import (
	"log"
	"os"

	"capibara/internal"
	"github.com/gin-gonic/gin"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

func main() {
	dsn := getEnv("DSN")
	apiKey := getEnv("API_KEY")

	db := internal.ConnectToDB(dsn)
	defer db.Close()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		if c.GetHeader("X-API-Key") != apiKey {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	})

	internal.Route(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
