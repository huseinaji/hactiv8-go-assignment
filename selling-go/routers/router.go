package routers

import (
	"selling-go/controllers"
	"selling-go/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	orderRouter := r.Group("/order")
	{
		orderRouter.Use(middlewares.Authentication())
		orderRouter.POST("/", controllers.CreateOrder)
	}

	return r
}
