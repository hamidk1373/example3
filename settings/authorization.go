package settings

import "github.com/dgrijalva/jwt-go"

// JWTCustomClaims declatres what to be saved after a user logged in.
type JWTCustomClaims struct {
	Email string `json:"email"`
	// Role  string `json:"role"`
	jwt.StandardClaims
}
