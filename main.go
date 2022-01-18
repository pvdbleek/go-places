package main

import (
	"github.com/pvdbleek/go-places/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/place", controllers.CreatePlace)
		v1.GET("/place/:id", controllers.GetPlace)
		v1.GET("/places", controllers.GetAllPlaces)
		v1.DELETE("/place", controllers.DeletePlace)
		v1.GET("/health", controllers.HealthCheck)
		v1.GET("/url/:id", controllers.GetPlaceUrl)
	}
	return router
}
