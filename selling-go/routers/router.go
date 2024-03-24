package routers

import (
	"selling-go/controllers"
	"selling-go/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartApp(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	inDB := &controllers.Handler{
		DB: db,
	}

	userRouter := r.Group("/user")
	{
		userRouter.GET("/", middlewares.Authentication(), inDB.GetAllUser)
		userRouter.DELETE("/:id", middlewares.Authentication(), inDB.DeleteUserById)
		userRouter.POST("/register", inDB.UserRegister)
		userRouter.POST("/login", inDB.UserLogin)
		userRouter.PATCH("/:id", middlewares.Authentication(), inDB.UpdateUserById)
	}
	productRouter := r.Group("/product")
	{
		productRouter.POST("/", middlewares.Authentication(), inDB.CreateProduct)
		productRouter.PATCH("/:id", middlewares.Authentication(), inDB.UpdateProductById)
		productRouter.GET("/", inDB.GetAllProduct)
		productRouter.DELETE("/:id", middlewares.Authentication(), inDB.DeleteProductById)
	}
	orderRouter := r.Group("/order")
	{
		orderRouter.POST("/", middlewares.Authentication(), inDB.CreateOrder)
		orderRouter.GET("/", middlewares.Authentication(), inDB.GetAllOrder)
		orderRouter.GET("/:id", middlewares.Authentication(), inDB.GetOrderById)
		orderRouter.DELETE("/:id", middlewares.Authentication(), inDB.DeleteOrder)
	}
	paymentRouter := r.Group("/payment")
	{
		paymentRouter.POST("/", middlewares.Authentication(), inDB.CreatePayment)
		paymentRouter.GET("/", middlewares.Authentication(), inDB.GetAllPayment)
		paymentRouter.DELETE("/:id", middlewares.Authentication(), inDB.DeletePaymentById)
	}
	invoiceRouter := r.Group("/invoice")
	{
		invoiceRouter.GET("/", middlewares.Authentication(), inDB.GetAllInvoice)
		invoiceRouter.GET("/:id", middlewares.Authentication(), inDB.GetInvoiceById)
	}

	return r
}
