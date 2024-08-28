package data

import golangjwt "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	golangjwt.RegisteredClaims
}
