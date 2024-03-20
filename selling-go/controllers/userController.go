package controllers

import (
	"net/http"
	"selling-go/databases"
	"selling-go/helpers"
	"selling-go/models"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	db := databases.GetDB()
	var User = models.User{}

	c.ShouldBindJSON(&User)

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      User.Role,
	})
}

func UserLogin(c *gin.Context) {
	db := databases.GetDB()
	var User = models.User{}

	c.ShouldBindJSON(&User)
	password := User.Password
	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorize",
			"message": "invalid email/password",
		})
		return
	}

	comparedPass := helpers.ComparePass(User.Password, password)

	if !comparedPass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "invalid email/password",
		})
		return
	}

	token = 
}
