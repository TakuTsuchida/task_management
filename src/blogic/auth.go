package blogic

import "github.com/dgrijalva/jwt-go"

type jwtClaim struct {
	UID   int    `json:"uid"`
	email string `json:"email"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")
