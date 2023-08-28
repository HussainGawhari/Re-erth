package api

import (
	"client-admin/models"
	db "client-admin/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addclient(c *gin.Context) {
	fmt.Println(" creating client")
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := db.InsertNewCLient(client)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    client,
	})
}

func Getclients(c *gin.Context) {

}
func GetClient(c *gin.Context) {

}
func EditClient(c *gin.Context) {

}
func Adduser(user models.User) {

}
