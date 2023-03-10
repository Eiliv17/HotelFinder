package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Eiliv17/HotelFinder/models"
	"github.com/gin-gonic/gin"
)

func SearchNearby(c *gin.Context) {
	if os.Getenv("GO_ENV") != "production" {
		c.Header("Access-Control-Allow-Origin", "*")
	}

	// takes the query from the context
	latitudeRaw := c.Query("latitude")
	longitudeRaw := c.Query("longitude")
	radiusRaw := c.Query("radius")
	offsetRaw := c.Query("offset")
	limitRaw := c.Query("limit")

	// required values check
	if latitudeRaw == "" || longitudeRaw == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "coordinates not present",
		})
		return
	}

	if radiusRaw == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "radius not present",
		})
		return
	}

	// converts the raw queries
	latitude, err := strconv.ParseFloat(latitudeRaw, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	longitude, err := strconv.ParseFloat(longitudeRaw, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	radius, err := strconv.ParseFloat(radiusRaw, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	offset, err := strconv.ParseInt(offsetRaw, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	limit, err := strconv.ParseInt(limitRaw, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// create point object with provided coordinates
	point, err := models.CreatePoint(latitude, longitude)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// find hotel based on point coordinates and parameters
	hotels, err := models.SearchHotel(
		c.Request.Context(),
		point,
		radius,
		int(offset),
		int(limit),
	)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, hotels)
}
