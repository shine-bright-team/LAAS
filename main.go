package main

import (
	"github.com/shine-bright-team/LAAS/v2/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	if !fiber.IsChild() {
		for {
			time.Sleep(time.Hour * 24)
			log.Print("Adding Interest Rate to Load")
			utils.AddInterestRateToLoad()
		}
	}

	app.Use(cors.New())

	app.Get("/", routes.DefaultPage)

	log.Print("Initialize routes")

	initialize.Router(app)

	log.Print("Listening to port 8000")

	app.Listen(":8000")
}
