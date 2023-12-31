package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	// user info fields
	UserID string
	Role   string
	Status string
	// base field(s)
	jwt.RegisteredClaims
}
