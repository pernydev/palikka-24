package models

import "github.com/golang-jwt/jwt/v5"

type AuthTokenData struct {
	ID    string `json:"id"`
	Staff bool   `json:"staff"`
	jwt.RegisteredClaims
}
