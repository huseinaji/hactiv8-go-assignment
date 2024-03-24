package controllers

import (
	"net/http"
	"selling-go/helpers"
	"selling-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (idb Handler) UserRegister(c *gin.Context) {
	var User = models.User{}

	err := c.ShouldBindJSON(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request body must be JSON",
		})
		return
	}

	err = idb.DB.Debug().Create(&User).Error

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

func (idb Handler) UserLogin(c *gin.Context) {
	var User = models.User{}

	err := c.ShouldBindJSON(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request body must be JSON",
		})
		return
	}

	password := User.Password
	err = idb.DB.Where("email = ?", User.Email).Take(&User).Error

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

	token := helpers.GenToken(User.ID, User.Email, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"id":    User.ID,
		"email": User.Email,
		"role":  User.Role,
		"token": token,
	})
}

func (idb Handler) GetAllUser(c *gin.Context) {
	User := []models.User{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	//authorization
	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	// err := idb.DB.Select("id", "user_name", "email", "role").Find(&User).Error
	err := idb.DB.Preload("Orders").Find(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   User,
	})
}

func (idb Handler) DeleteUserById(c *gin.Context) {
	User := models.User{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.First(&User, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	err = idb.DB.Delete(&User, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "User deleted",
		"user":   User.Email,
	})
}

func (idb Handler) UpdateUserById(c *gin.Context) {
	User := models.User{}
	newData := models.User{}
	id := c.Param("id")
	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User doesn't have permission to access this service",
		})
		return
	}

	err := idb.DB.First(&User, id).Error

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

	err = idb.DB.Model(&User).Updates(newData).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   User,
	})
}
