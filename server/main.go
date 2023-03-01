package main

import (
	"os"

	"github.com/Eiliv17/HotelFinder/controllers"
	"github.com/Eiliv17/HotelFinder/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
}

func main() {

	// setting up for production
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// creating the gin engine and setting up configurations
	r := gin.Default()
	r.SetTrustedProxies(nil)

	api := r.Group("/v1")
	{
		// search for nearby hotels endpoint
		api.GET("/search/nearby", controllers.SearchNearby)

		// hotels endpoints
		hotels := api.Group("/hotels")
		{
			// return detailed info about an hotel
			hotels.GET("/:id")

			// add an hotel
			hotels.POST("")

			// update details of an hotel
			hotels.PUT("/:id")

			// delete an hotel
			hotels.DELETE("/:id")
		}
	}

	// running the gin engine
	if os.Getenv("GO_ENV") != "production" {
		r.Run()
	} else {
		port := os.Getenv("PORT")
		r.RunTLS(":"+port, "servercert.pem", "serverkey.key")
	}
}
