package controller

import (
	"client-admin/models"
	"client-admin/pkg/helperdb"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Here we can add new clients
func Addclient(c *gin.Context) {
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

// This function returns clients by last name, city, country, and so on..
func Getclient(c *gin.Context) {

	query := "SELECT * FROM clients WHERE"
	// Retrieve query parameters
	if lastName := c.DefaultQuery("last_name", "jan"); lastName != "" {
		query += fmt.Sprintf(" last_name = '%s'", lastName)
	}
	if postalCode := c.Query("postal_code"); postalCode != "" {
		query += fmt.Sprintf("and postal_code = '%s'", postalCode)
	}
	if city := c.Query("city"); city != "" {
		query += fmt.Sprintf("and city = '%s'", city)
	}
	if country := c.Query("country"); country != "" {
		query += fmt.Sprintf(" and country = '%s'", country)
	}
	fmt.Println(query)
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

// THis function will delete the client by ID
func DeleteClient(c *gin.Context) {
	clientID := c.Param("id")
	id, err := strconv.Atoi(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}
	err = helperdb.DeleteInDb(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "client deleted successfully"})
}

// This function return all the active client
func GetAllclients(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	offset := (pageInt - 1) * pageSizeInt

	clients, err := helperdb.Clients(pageSizeInt, offset)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    clients,
	})
}
func EditClient(c *gin.Context) {
	var client models.Clients
	clientID := c.Param("id")
	// Ensure that the client ID is a valid integer
	id, err := strconv.Atoi(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	// Update the client's details in the database
	err = helperdb.Update(client, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating client"})
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "client updated successfully",
		"status":  true,
		"data":    client,
	})
}

func CountClients(c *gin.Context) {
	total, err := helperdb.CountAllClient()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error number of client"})
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "This is the numbers all client ",
		"status":  true,
		"data":    total,
	})
}
func ClientHistory(c *gin.Context) {
	rs, err := helperdb.ClientHistoryData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting of client"})
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "This is the history all client ",
		"status":  true,
		"data":    rs,
	})
}

func ActivateDeactivat(c *gin.Context) {

	clientID := c.Param("id")
	// Ensure that the client ID is a valid integer
	id, err := strconv.Atoi(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	// Update the client's status in the database
	err = helperdb.ChangeStatus(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating client"})
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "client status updated successfully",
	})
}

func GetClientBasedID(c *gin.Context) {
	clientID := c.Param("id")
	// Ensure that the client ID is a valid integer
	id, err := strconv.Atoi(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	result, err := helperdb.GetClient(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not find the client",
			"status":  "404",
		})
		return
	}

	// Return a success response to user
	c.JSON(http.StatusOK, gin.H{
		"messsage": "success",
		"data":     result,
		"status":   200,
	})
}
