package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claim Token de usuario
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
