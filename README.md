# homeapp

The goal is to create a simple smarthome application, to control e.g. hue lightbulbs.

The application architecture is a microservice architecture, with a frontend based on Vue/Vite in combination with tailwind and tailwindUi. 

The first service (lightning) is based on golang. 

## Get Started

Install Dependencies first:

Within `/frontend` run `npm install` and afterwards `npm run dev`
Within `/services/lightning` run `go mod tidy` and afterwards `go run .`
