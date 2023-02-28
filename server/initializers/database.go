package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client
var DBName string

func LoadDatabase() {
	var err error

	// retrieve env variable
	URI := os.Getenv("MONGODB_URI")

	// Create a new client and connect to the server
	DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged MongoDB.")

	DBName = os.Getenv("DB_NAME")
}
