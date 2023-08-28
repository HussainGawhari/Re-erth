package models

type Client struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   Address
	IsActive  bool `json:"isActive"`
}

type Address struct {
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Country    string `json:"country"`
}
