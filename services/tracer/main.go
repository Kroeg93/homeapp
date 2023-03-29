package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"services/lightning/handlers"
)

func main() {
	const FrontendUrl = "http://localhost:5174"

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FrontendUrl},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	router.GET("api/lightning/lightbulb", handlers.GetLightbulbs)
	router.GET("api/lightning/lightbulbs", handlers.GetLightbulbs)
	//router.POST('api/lightning/lightbulb/:?/state', handlers.ChangeLightbulbState)

	router.Run("127.0.0.1:3000")
}
