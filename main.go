package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shine-bright-team/LAAS/v2/initialize"
	"github.com/shine-bright-team/LAAS/v2/routes"
)

func main() {
	app := fiber.New()
	log.Print("Looking up env")
	if mode, isfound := os.LookupEnv("MODE"); !isfound || mode == "DEV" {
		log.Print("Coundn't find env, looking up .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	log.Print("Connecting to DB")
	initialize.DbSetUp()

	app.Get("/", routes.DefaultPage)

	log.Print("Initialize routes")

	initialize.Router(app)

	log.Print("Listening to port 8000")

	app.Listen(":8000")

	log.Print("App is ready")
}
