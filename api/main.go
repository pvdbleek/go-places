package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type place struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Ctry string  `json:"country"`
	Desc string  `json:"description"`
	Lat  float64 `json:"latitude"`
	Lon  float64 `json:"longitude"`
}

var places = []place{
	{ID: "1", Name: "Darvaza Gas Crater", Ctry: "Turkmenistan", Desc: "Also known as the Gate to Hell.", Lat: 40.252605604792635, Lon: 58.439763430286064},
	{ID: "2", Name: "The Great Wall", Ctry: "China", Desc: "One of the ancient wonders of the world.", Lat: 40.4324742310965, Lon: 116.56400733368996},
}

func main() {
	router := gin.Default()
	router.GET("/places", getPlaces)
	router.GET("/places/:id", getPlaceByID)
	router.GET("/url/:id", getMapsUrlByID)
	router.POST("/places", postPlaces)
	router.GET("/health", healthCheck)

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	status := "OK"
	c.IndentedJSON(http.StatusOK, status)
}

func getPlaces(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, places)
}

func postPlaces(c *gin.Context) {
	var newPlace place

	if err := c.BindJSON(&newPlace); err != nil {
		return
	}

	places = append(places, newPlace)
	c.IndentedJSON(http.StatusCreated, newPlace)
}

func getPlaceByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range places {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "place not found"})
}

func getMapsUrlByID(c *gin.Context) {
	id := c.Param("id")

	for _, MapCoordinates := range places {
		if MapCoordinates.ID == id {
			UrlLat := strconv.FormatFloat(MapCoordinates.Lat, 'f', -1, 64)
			UrlLon := strconv.FormatFloat(MapCoordinates.Lon, 'f', -1, 64)
			Url := "https://maps.google.com/maps?q=" + UrlLat + "," + UrlLon

			c.PureJSON(http.StatusOK, gin.H{"mapsurl": Url + "&t=k"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "place not found"})
}
