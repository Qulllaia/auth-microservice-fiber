package dto

import (
	"github.com/golang-jwt/jwt/v5"
)
type AuthDto struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims 
}
