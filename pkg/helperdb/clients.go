package helperdb

import (
	"client-admin/models"
	"fmt"
)

func InsertNewCLient(c models.Clients) error {

	fmt.Println("Insert results")
	// Query executes a query that returns rows, typically a SELECT
	result, err := DB.Exec("INSERT INTO clients (first_name, last_name, telephone, email,status,street, postal_code, city, country) VALUES ($1, $2, $3,$4, $5, $6, $7, $8, $9)",
		c.FirstName, c.LastName, c.Telephone, c.Email, c.Status, c.Street, c.PostalCode, c.City, c.Country)
	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}
	fmt.Println(result)

	return nil

}

func ListClient(query string) ([]models.Clients, error) {
	var c []models.Clients
	rows, err := DB.Query(query)
	if err != nil {

		return c, err
	}

	defer rows.Close()

	var clients []models.Clients
	var first_name, last_name, email, telephone, postal_code, country, city, street string
	var status bool
	for rows.Next() {
		var client models.Clients
		if err := rows.Scan(&first_name, &last_name, &telephone, &email, &status, &street, &postal_code, &city, &country); err != nil {
			return c, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
