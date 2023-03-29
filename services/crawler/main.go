package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"services/crawler/lib"
)

const BaseUrl = "localhost"
const Port = "5175"

func main() {
	var ServiceUrl = fmt.Sprintf("%s:%s", BaseUrl, Port)
	const FrontendUrl = "http://localhost:5174"

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FrontendUrl},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	lib.Crawl()

	router.Run(ServiceUrl)
}
