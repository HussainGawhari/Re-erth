package controller

import (
	"client-admin/models"
	"client-admin/pkg/helperdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	fmt.Println(" creating user")
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println("user details", user.Name, user.Email)

	err := helperdb.AddUser(user)
	if err != nil {
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

	var validate models.Login

	fmt.Println(" checking user")
	if err := c.ShouldBindJSON(&validate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println("user details", validate.Email, validate.Password)

	token, err := helperdb.CheckUser(validate)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"Token":   token,
	})

}

func GetToken(c *gin.Context) {

	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
