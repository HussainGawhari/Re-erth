package controller

import (
	"client-admin/models"
	"client-admin/pkg/helperdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addclient(c *gin.Context) {
	fmt.Println(" creating client")
	var client models.Clients
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := helperdb.InsertNewCLient(client)
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
	query := "SELECT * FROM clients WHERE TRUE "

	// Retrieve query parameters
	if lastName := c.Query("last_name"); lastName != "" {
		query += fmt.Sprintf(" AND last_name = '%s'", lastName)
	}
	if postalCode := c.Query("postal_code"); postalCode != "" {
		query += fmt.Sprintf(" AND postal_code = '%s'", postalCode)
	}
	if city := c.Query("city"); city != "" {
		query += fmt.Sprintf(" AND city = '%s'", city)
	}
	if country := c.Query("country"); country != "" {
		query += fmt.Sprintf(" AND country = '%s'", country)
	}

	clients, err := helperdb.ListClient(query)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    clients,
	})

}
func GetClient(c *gin.Context) {

}
func EditClient(c *gin.Context) {

}
