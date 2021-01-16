package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	jwt.StandardClaims
}
