package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
