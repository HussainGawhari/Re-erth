package helperdb

import (
	"client-admin/models"
	"errors"
	"fmt"
)

func InsertNewCLient(c models.Clients) error {
	_, err := DB.Exec("INSERT INTO clients (first_name, last_name, telephone, email,status,street, postal_code, city, country) VALUES ($1, $2, $3,$4, $5, $6, $7, $8, $9)",
		c.FirstName, c.LastName, c.Telephone, c.Email, c.Status, c.Street, c.PostalCode, c.City, c.Country)
	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}
	ClientHistoryData()
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
	row := DB.QueryRow("SELECT id FROM clients WHERE id = $1", id)
	var idFromDb int
	err := row.Scan(&idFromDb)
	if err != nil {
		fmt.Print("ID does not exist in DB")
		return err
	}

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
	ClientHistoryData()
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
	ClientHistoryData()
	return nil

}

func CountAllClient() (int, error) {
	var totalRecords int
	err := DB.QueryRow("SELECT COUNT(id) FROM clients").Scan(&totalRecords)
	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func ClientHistoryData() ([]models.ClientsHistory, error) {
	var clients_history []models.ClientsHistory
	rows, err := DB.Query("SELECT * FROM clients_history")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	for rows.Next() {
		var client models.ClientsHistory
		err := rows.Scan(
			&client.ID,
			&client.ClientID,
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

		clients_history = append(clients_history, client)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return clients_history, nil
}

func ChangeStatus(id int) error {
	row := DB.QueryRow("SELECT status FROM clients WHERE id = $1", id)
	var statusFromDB bool
	err := row.Scan(&statusFromDB)
	if err != nil {
		return errors.New("error in querying client")
	}

	// Update the status value
	statusFromDB = !statusFromDB

	query := "UPDATE clients SET status = $1 WHERE id = $2"
	_, err = DB.Exec(query, statusFromDB, id)

	if err != nil {
		fmt.Println("Error updating the status:", err)
		return err
	}
	ClientHistoryData()
	return nil
}
