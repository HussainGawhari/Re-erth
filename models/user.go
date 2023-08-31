package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type role string

const (
	RoleAdmin role = "admin"
	RoleUser  role = "user"
	RoleGuest role = "client"
)

type status string

const (
	StatusActive   status = "active"
	StatusInactive status = "inactive"
)

type Users struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email" validate:"required"`
	Password  string     `json:"password" validate:"required"`
	Role      role       `json:"role" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JWTClaim struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}
