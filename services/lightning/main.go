package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"services/lightning/handlers"
)

func main() {
	const LightningServiceUrl = "http://localhost:5174"
	const HUE_BRIDGE_IP = "192.168.178.60"
	const HUE_USER_NAME = "xRrAsY4hVE3Avx6VS68Q-eq-jJ8G8S6XWYXzkpXT"

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{LightningServiceUrl},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	router.GET("api/lightning/lightbulb", handlers.GetLightbulbs)
	router.GET("api/lightning/lightbulbs", handlers.GetLightbulbs)

	router.Run("127.0.0.1:3000")
}
