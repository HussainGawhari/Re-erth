package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type status string
type role string

// Enum for Status
const (
	StatusActive   status = "ACTIVE"
	StatusInActive status = "INACTIVE"
)

// Enum for Role
const (
	RoleAdmin role = "ADMIN"
	RoleSEO   role = "SEO"
)

type Users struct {
	UserID primitive.ObjectID `bson:"_id" json:"userID"`

	Name     string `bson:"name" json:"name"`
	EmailID  string `bson:"emailID" json:"emailID" validate:"required"`
	Status   status `bson:"status" json:"status"`
	Role     role   `bson:"role" json:"role" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`

	CreatedAt time.Time  `bson:"createdAt" json:"-"`
	UpdatedAt *time.Time `bson:"updatedAt" json:"-"`
}

type JWTClaim struct {
	Name    string `json:"name"`
	EmailID string `json:"emailID"`
	Id      string `bson:"_id" json:"userID"`
	Role    string `bson:"role" json:"role"`
	jwt.StandardClaims
}
