package routers

import (
	"go_restapi_assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.TesServer)
	router.POST("/order", controllers.CreateOrders)
	router.GET("/order", controllers.GetOrders)
	router.PUT("/order/:orderId", controllers.UpdateOrders)
	router.DELETE("/order/:orderId", controllers.DeleteOrders)

	return router
}
