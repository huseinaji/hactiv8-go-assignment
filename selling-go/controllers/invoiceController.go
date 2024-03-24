package controllers

import (
	"net/http"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (idb Handler) GetAllInvoice(c *gin.Context) {
	Invoice := []models.Invoice{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.Find(&Invoice).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Invoice,
	})
}

func (idb Handler) GetInvoiceById(c *gin.Context) {
	Invoice := models.Invoice{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	err := idb.DB.First(&Invoice, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
	}

	if userData["role"] != "admin" && Invoice.UserID != uint(userData["id"].(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"user":   Invoice,
	})
}
