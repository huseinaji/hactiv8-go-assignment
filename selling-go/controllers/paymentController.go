package controllers

import (
	"net/http"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (idb Handler) CreatePayment(c *gin.Context) {
	Payment := models.Payment{}
	Invoice := models.Invoice{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	err := c.ShouldBindJSON(&Payment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
	}

	Payment.UserID = uint(userData["id"].(float64))

	if userData["role"] != "admin" && Payment.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err = idb.DB.Find(&Invoice, Payment.InvoiceID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
	}

	Payment.Amount = Invoice.TotalPayment

	err = idb.DB.Create(&Payment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	//Update Invoice status
	Invoice.Status = "Paid"

	err = idb.DB.Save(&Invoice).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Payment,
	})
}

func (idb Handler) GetAllPayment(c *gin.Context) {
	Payment := []models.Payment{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.Find(&Payment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
			"data":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Payment,
	})
}

func (idb Handler) DeletePaymentById(c *gin.Context) {
	Payment := models.Payment{}
	id := c.Param("id")
	Invoice := models.Invoice{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	err := idb.DB.First(&Payment, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	if userData["role"] != "admin" && Payment.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err = idb.DB.First(&Invoice, Payment.InvoiceID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	Invoice.Status = "Unpaid"
	err = idb.DB.Save(&Invoice).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = idb.DB.Delete(&Payment, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Payment,
	})
}
