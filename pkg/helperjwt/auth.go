package helperjwt

import (
	"client-admin/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// secret key
var jwtKey = []byte("supersecretkey")

// Generate JWT Token
func GenerateJWT(email, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := models.JWTClaim{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// Validate token for admin
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return errors.New("the token is not valid")

	}
	claims, ok := token.Claims.(*models.JWTClaim)

	if !ok {
		err := errors.New("couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err := errors.New("token expired")
		return err
	}
	role := claims.Role
	if role == "user" {
		// Create a permission error
		err := errors.New("you don't have permission to access this resource")
		return err
	}
	return nil
}

// Validate token for the users
func ValidateTokenUserRole(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return errors.New("provided token is not valid")
	}
	claims, ok := token.Claims.(*models.JWTClaim)

	if !ok {
		err := errors.New("couldn't parse the claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err := errors.New("token expired")
		return err
	}
	return nil
}

// Hash user password
func HashPassword(user models.Users) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return "", err
	}
	user.Password = string(bytes)
	return user.Password, nil
}

// check password
func CheckPassword(dbPass string, user models.Login) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}

// check password without hashed
func CheckPasswordWitoutHash(providedPassword string, password string) error {
	if password != providedPassword {
		return errors.New("password not valid")
	}
	return nil
}
