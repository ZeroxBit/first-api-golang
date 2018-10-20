package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim Token del user
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
