package main

import (
	"github.com/Raihanki/horizont-api/configs"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadConfig()

	app := fiber.New()

	errListen := app.Listen(":" + configs.ENV.APP_PORT)
	if errListen != nil {
		log.Fatal(errListen)
	}
}
