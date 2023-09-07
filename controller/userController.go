package controller

import (
	"client-admin/models"
	"client-admin/pkg/helperdb"
	"client-admin/pkg/helperjwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var doc models.Users
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	// fmt.Println("user details", user.Name, user.Email)
	user := models.Users{

		Name:      doc.Name,
		Email:     doc.Email,
		Password:  doc.Password,
		Role:      doc.Role,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}
	// to hash the user password
	password, err := helperjwt.HashPassword(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	user.Password = password

	errs := helperdb.AddUser(user)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

func LoginUser(c *gin.Context) {
	var user models.Login
	fmt.Println(" checking user")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	token, role, err := helperdb.CheckUser(user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if role == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success ",
		"Token":   token,
		"Roles":   role,
	})
}

func GetToken(c *gin.Context) {

	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func GetUsers(c *gin.Context) {
	users, err := helperdb.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    users,
	})
}
