package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shine-bright-team/LAAS/v2/initialize"
	"github.com/shine-bright-team/LAAS/v2/routes"
)

func main() {
	app := fiber.New()
	initialize.LookForEnv()
	log.Print("Connecting to DB")
	if err := initialize.DbSetUp(); err != nil {
		log.Fatalf("There is an error setting up Database: %s", err)
	}

	app.Get("/", routes.DefaultPage)

	log.Print("Initialize routes")

	initialize.Router(app)

	log.Print("Listening to port 8000")

	app.Listen(":8000")
}
