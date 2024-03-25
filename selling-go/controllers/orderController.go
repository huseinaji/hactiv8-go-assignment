package controllers

import (
	"net/http"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (idb Handler) CreateOrder(c *gin.Context) {
	Order := models.Order{}
	Product := models.Product{}

	userData := c.MustGet("userData").(jwt.MapClaims)

	err := c.ShouldBindJSON(&Order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Requests",
			"message": "Request body must be JSON",
		})
	}

	err = idb.DB.First(&Product, Order.ProductID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Product Id Not found",
		})
		return
	}

	//validasi product qty
	if Product.Quantity < Order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not enough quantity on this product",
		})
		return
	}

	Product.Quantity = Product.Quantity - Order.Quantity

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

	err = idb.DB.Preload("Payments").Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = idb.DB.Save(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Created",
		"data":    Order,
	})
}

func (idb Handler) GetAllOrder(c *gin.Context) {
	Order := []models.Order{}

	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.Preload("Invoices").Preload("Payments").Find(&Order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"order":  Order,
	})
}

func (idb Handler) GetOrderById(c *gin.Context) {
	Order := models.Order{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)
	err := idb.DB.Preload("Invoices").First(&Order, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if userData["role"] != "admin" && Order.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"user":   Order,
	})
}

func (idb Handler) DeleteOrder(c *gin.Context) {
	Order := models.Order{}
	Invoice := models.Invoice{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	//validasi punya payment atau tidak
	err := idb.DB.Preload("Payments").First(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	if userData["role"] != "admin" && Order.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	if Order.Payments == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "This order has not paid yet",
		})
		return
	}

	err = idb.DB.Delete(&Invoice, Order.Invoices.ID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = idb.DB.Delete(&Order, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"user":   Order,
	})
}
