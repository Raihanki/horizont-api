package main

import (
	"github.com/Raihanki/horizont-api/configs"
	v1 "github.com/Raihanki/horizont-api/routes/v1"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadConfig()
	db, errDatabase := configs.ConnectDatabase()
	if errDatabase != nil {
		log.Fatal(errDatabase)
	}

	app := fiber.New()
	v1.Routes(app, db)

	errListen := app.Listen(":" + configs.ENV.APP_PORT)
	if errListen != nil {
		log.Fatal(errListen)
	}
}
