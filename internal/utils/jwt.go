// internal/utils/jwt.go
package utils

import "github.com/golang-jwt/jwt/v5"

// JWTClaims — структура полезной нагрузки токена
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
