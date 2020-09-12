package model

import "github.com/dgrijalva/jwt-go"

//JwtToken struct
type JwtToken struct {
	Token   string `json:"token"`
	Success string `json:"success"`
}

//Exception struct
type Exception struct {
	Message string `json:"message"`
}

//Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
