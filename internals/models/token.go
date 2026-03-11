package models

import "github.com/golang-jwt/jwt/v5"

type TokenBody struct {
	Id     uint   `json:"id"`
	Mobile string `json:"mobile"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}