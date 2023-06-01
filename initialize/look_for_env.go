package initialize

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LookForEnv() {
	log.Print("Looking up env")
	if mode, isfound := os.LookupEnv("MODE"); !isfound || mode == "DEV" {
		log.Print("Coundn't find env, looking up .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
