package helperdb

import (
	"client-admin/models"
	"client-admin/pkg/helperjwt"
	"fmt"
)

func AddUser(user models.Users) error {
	// Query executes a query that returns rows, typically a SELECT
	_, err := DB.Exec("INSERT INTO users (name, email_id, password,role) VALUES ($1, $2, $3,$4)", user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}

	return nil
}

func CheckUser(eval models.Login) (string, error) {
	// Execute a query that returns data
	var email string
	var password string
	rows, err := DB.Query("SELECT email_id , password FROM users")
	if err != nil {
		panic(err)

	}
	defer rows.Close()
	// Scan the data into variables
	for rows.Next() {

		err := rows.Scan(&email, &password)
		if err != nil {
			panic(err)
		}
		// Print the data
		err = helperjwt.CheckPasswordWitoutHash(password, eval.Password)
		if err != nil {
			return "", err
		}

		// return tokenString, nil
		// fmt.Println("verified", email, tokenString)
	}
	tokenString, err := helperjwt.GenerateJWT(email)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
