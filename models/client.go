package models

import (
	"time"
)

type Client struct {
	// ID primitive.ObjectID `bson:"_id" json:"ID" `
	// UserID primitive.ObjectID `bson:"userID" json:"userID"`
	ID        int    `json:"firstName"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   Address
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `bson:"createdAt" json:"-"`
	UpdatedAt *time.Time `bson:"updatedAt" json:"-"`
}

type Address struct {
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Password string `json:"-"`
// 	Role     string `json:"role"`
// }

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
