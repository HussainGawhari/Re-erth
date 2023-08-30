package models

import (
	"time"
)

type Clients struct {
	ID         int        `json:"id" validate:"required"`
	FirstName  string     `json:"first_name" validate:"required"`
	LastName   string     `json:"last_name" validate:"required"`
	Telephone  string     `json:"telephone" validate:"required"`
	Email      string     `json:"email" validate:"required,email"`
	Status     bool       `json:"status"`
	Street     string     `json:"street" validate:"required"`
	PostalCode string     `json:"postal_code" validate:"required"`
	City       string     `json:"city" validate:"required"`
	Country    string     `json:"country" validate:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type GetClientBy struct {
	ID         int        `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Telephone  string     `json:"telephone"`
	Email      string     `json:"email"`
	Status     bool       `json:"status"`
	Street     string     `json:"street"`
	PostalCode string     `json:"postal_code"`
	City       string     `json:"city"`
	Country    string     `json:"country"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// type Address struct {
// 	Street     string `json:"street" validate:"required"`
// 	PostalCode string `json:"postal_code" validate:"required"`
// 	City       string `json:"city" validate:"required"`
// 	Country    string `json:"country" validate:"required"`
// }
