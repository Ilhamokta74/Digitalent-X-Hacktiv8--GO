package router

import (
	"Go-JWT/controllers"
	"Go-JWT/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/get", controllers.GetAllUsers)

		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

	}

	productRouter := r.Group("/product")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreatePoduct)

		productRouter.PUT("/:productId", middlewares.Authentication(), controllers.UpdatePoduct)
	}

	return r
}
