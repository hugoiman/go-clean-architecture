package dto

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
