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
func GenerateJWT(email, pass, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := models.JWTClaim{
		Email:    email,
		Password: pass,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// Validate token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
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
// func CheckPassword(providedPassword string, user models.Login) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
func CheckPassword(providedPassword string, hashedPasswordFromDB string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordFromDB), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func CheckPasswordWitoutHash(providedPassword string, password string) error {
	if password != providedPassword {
		return errors.New("password not valid")
	}
	return nil
}
