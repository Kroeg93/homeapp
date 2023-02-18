package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"services/lightning/api"
	. "services/lightning/models"
)

func GetLightbulb(c *gin.Context) {
	var lightbulb LightbulbDTO

	lightbulb, err := api.GetLightbulb(1)

	if err != nil {
		log.Println("Error while fetching lightbulbs")
		log.Println(err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Println("SUCCESS: %v", lightbulb)
	}

	c.JSON(http.StatusOK, lightbulb)
}

func GetLightbulbs(c *gin.Context) {
	var lightbulbsDTO LightbulbsDTO

	lightbulbsDTO, err := api.GetLightbulbs()

	if err != nil {
		log.Println("Error while fetching lightbulbs")
		log.Println(err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Println("SUCCESS: %v", lightbulbsDTO)
	}

	var lightbulbs []LightbulbDTO

	for _, value := range lightbulbsDTO {
		lightbulbs = append(lightbulbs, value)
	}

	c.JSON(http.StatusOK, lightbulbs)
}
