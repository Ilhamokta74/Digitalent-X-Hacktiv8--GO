package routers

import (
	"Gin-Framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Create Data
	router.POST("/cars", controllers.CreateCar)

	// Update Data
	router.PUT("/cars/:carID", controllers.UpdateCar)

	// Get Data Spesific
	router.GET("/cars/:carID", controllers.GetCar)

	// Get All Data
	router.GET("/cars", controllers.GetAllCars)

	// Delete Data
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
