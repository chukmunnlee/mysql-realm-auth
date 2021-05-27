package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func CreateClaim(subject string, issuer, audience string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Id:       (uuid.New()).String()[0:8],
		Issuer:   issuer,
		Subject:  subject,
		Audience: audience,
		IssuedAt: time.Now().Unix(),
	}
}

func GenerateJWT(claims *jwt.StandardClaims, signKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signKey))
}
