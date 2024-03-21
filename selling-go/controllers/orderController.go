package controllers

import (
	"fmt"
	"net/http"
	"selling-go/databases"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateOrder(c *gin.Context) {
	db := databases.GetDB()
	Order := models.Order{}
	Product := models.Product{}

	userData := c.MustGet("userData").(jwt.MapClaims)

	c.ShouldBindJSON(&Order)
	err := db.Where("id = ?", Order.ProductID).Take(&Product).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err,
			"message": "Product Id Not found",
		})
		return
	}
	fmt.Println("tes", userData["email"], userData["id"])

	Order.UserID = uint(userData["id"].(float64))

	Invoice := models.Invoice{
		OrderID:      Order.ID,
		UserID:       uint(userData["id"].(float64)),
		Quantity:     Order.Quantity,
		Rate:         Product.Rate,
		TotalPayment: Order.Quantity * Product.Rate,
		Status:       "Unpaid",
	}

	Order.Invoices = Invoice

	fmt.Printf("%+v", Order)
	err = db.Debug().Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Order Created",
		"data":          Order,
		"verifiedToken": userData,
	})
}
