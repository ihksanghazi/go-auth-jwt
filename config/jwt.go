package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("skjsahfdlllfaslfuasndflkladsfh")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}