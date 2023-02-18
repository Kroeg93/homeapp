# homeapp

The goal is to create a simple smarthome application, to control e.g. hue lightbulbs.

The application architecture is a microservice architecture, with a frontend based on Vue/Vite in combination with tailwind and tailwindUi. 

The first service (lightning) is based on golang. 

## Get Started
### Frontend

- Inside `/frontend` run `npm install` and afterwards `npm run dev`

### Services

 To control hue lightbulbs over a Hue Bridge there are two informations necessary:
 - The IP of the Bridge within the local network
 - A user has to be created via the hue API (https://developers.meethue.com/develop/get-started-2/)

There are two ENV-Variables which need to be exported before the application can started:
- export HUE_BRIDGE_IP=ip_of_hue_bridge_within_network
- export HUE_USER=generated_hue_user

- Inside `/services/lightning` run `go mod tidy` and afterwards `go run .`
