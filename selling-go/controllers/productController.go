package controllers

import (
	"net/http"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (idb Handler) CreateProduct(c *gin.Context) {
	Product := models.Product{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	//authorization
	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := c.ShouldBindJSON(&Product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request body must be JSON",
			"data":    err.Error(),
		})
		return
	}

	err = idb.DB.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       Product.ID,
		"name":     Product.Name,
		"quantity": Product.Quantity,
		"rate":     Product.Rate,
	})
}

func (idb Handler) UpdateProductById(c *gin.Context) {
	Product := models.Product{}
	newData := models.Product{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.First(&Product, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&newData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request body must be JSON",
		})
		return
	}

	err = idb.DB.Model(&Product).Updates(newData).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Product,
	})
}
func (idb Handler) GetAllProduct(c *gin.Context) {
	Product := []models.Product{}

	err := idb.DB.Find(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   Product,
	})
}
func (idb Handler) DeleteProductById(c *gin.Context) {
	Product := models.Product{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.First(&Product, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	if Product.Orders != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "Product has been Ordered, try to cancel the order first",
		})
		return
	}

	err = idb.DB.Delete(&Product, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Product deleted",
	})
}
