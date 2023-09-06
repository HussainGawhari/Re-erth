package helperdb

import (
	"client-admin/models"
	"client-admin/pkg/helperjwt"
	"fmt"
)

func AddUser(user models.Users) error {
	// Query executes a query that returns rows, typically a SELECT
	_, err := DB.Exec("INSERT INTO users (name, email, password,role) VALUES ($1, $2, $3,$4)", user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}
	return nil
}

func CheckUser(user models.Login) (string, error) {
	var email string
	var password string
	var role string
	rows, err := DB.Query("SELECT email, password,role FROM users where email = $1", user.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Scan the data into variables
	for rows.Next() {
		err := rows.Scan(&email, &password, &role)
		if err != nil {
			panic(err)
		}
		err = helperjwt.CheckPassword(password, user)
		if err != nil {
			fmt.Print(err)
			return "", err
		}
	}

	fmt.Printf("testing \t %s, with role  %s  \t       ", email, role)
	tokenString, err := helperjwt.GenerateJWT(email, role)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// To get all users from database using this funtion
func GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	rows, err := DB.Query("SELECT id,name,email,password,role FROM users")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	for rows.Next() {
		var user models.Users
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return users, nil
}
