package token

import "github.com/dgrijalva/jwt-go"

// Contract of the JWTService can do

type IJwtService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}
