package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	. "services/lightning/models"
)

func GetLightbulb(id int) (LightbulbDTO, error) {
	client := resty.New()
	url := fmt.Sprintf("http://%s/api/%s/lights/%d", os.Getenv("HUE_BRIDGE_IP"), os.Getenv("HUE_USER"), id)

	var lightbulb LightbulbDTO

	res, err := client.R().
		SetResult(&lightbulb).
		Get(url)

	fmt.Println(res)

	if err != nil {
		fmt.Println("ERROR: %v", err)
	}

	return lightbulb, nil
}

func GetLightbulbs() (LightbulbsDTO, error) {
	client := resty.New()
	url := fmt.Sprintf("http://%s/api/%s/lights", os.Getenv("HUE_BRIDGE_IP"), os.Getenv("HUE_USER"))
	var lightbulbs LightbulbsDTO

	res, err := client.R().
		SetResult(&lightbulbs).
		Get(url)

	fmt.Println(res)

	if err != nil {
		fmt.Println("ERROR: %v", err)
		return nil, err
	}

	return lightbulbs, nil
}
