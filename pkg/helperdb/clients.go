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

	var clients []models.Clients
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	for rows.Next() {
		var client models.Clients

		err := rows.Scan(
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Telephone,
			&client.Email,
			&client.Status,
			&client.Street,
			&client.PostalCode,
			&client.City,
			&client.Country,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		clients = append(clients, client)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return clients, nil
}

func DeleteInDb(id int) error {
	row := DB.QueryRow("SELECT id FROM clients WHERE id = ?", id)
	var idFromDb int
	err := row.Scan(&idFromDb)
	if err != nil {
		fmt.Print("ID does not exist in DB")
		return err
	}
	fmt.Println("this is i ", idFromDb)
	stmt, err := DB.Prepare("DELETE FROM clients WHERE id = $1")
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func Clients(pageSize, offset int) ([]models.Clients, error) {
	var clients []models.Clients
	query := fmt.Sprintf("SELECT * FROM clients WHERE status = TRUE ORDER BY first_name ASC LIMIT %d OFFSET %d", pageSize, offset)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	for rows.Next() {
		var client models.Clients
		err := rows.Scan(
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Telephone,
			&client.Email,
			&client.Status,
			&client.Street,
			&client.PostalCode,
			&client.City,
			&client.Country,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		clients = append(clients, client)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return clients, nil
}

func Update(newClient models.Clients, id int) error {

	query := "UPDATE clients SET first_name = $1, last_name = $2, telephone = $3, email = $4, status = $5, street = $6, postal_code = $7, city = $8, country = $9 WHERE id = $10"

	_, err := DB.Exec(query,
		newClient.FirstName,
		newClient.LastName,
		newClient.Telephone,
		newClient.Email,
		newClient.Status,
		newClient.Street,
		newClient.PostalCode,
		newClient.City,
		newClient.Country,
		id,
	)
	if err != nil {
		return err
	}

	return nil

}
