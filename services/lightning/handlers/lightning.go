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
	var light LightDTO

	light, err := api.GetLightbulb(1)

	if err != nil {
		log.Println("Error while fetching lightbulbs")
		log.Println(err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Println("SUCCESS: %v", light)
	}

	c.JSON(http.StatusOK, light)
}

func GetLightbulbs(c *gin.Context) {
	var lightbulbsDTO LightsDTO

	lightbulbsDTO, err := api.GetLightbulbs()

	if err != nil {
		log.Println("Error while fetching lights")
		log.Println(err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Println("SUCCESS: %v", lightbulbsDTO)
	}

	var lights []LightDTO

	for _, value := range lightbulbsDTO {
		lights = append(lights, value)
	}

	c.JSON(http.StatusOK, lights)
}

func ChangeLightbulbState(c *gin.Context) {

}
