package request

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}
type BaseClaims struct {
	UserAccount string `json:"userAccount"`
	UserID      uint64 `json:"userID"`
}
