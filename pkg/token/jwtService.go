package token

import (
	"os"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type JwtService struct {
	secretKey string
	issuer    string
}

func getSecretKey() string {
	secretKey := os.Getenv("JWTSECRETKEY")
	if secretKey != "" {
		secretKey = "sfsdf"
	}
	return secretKey
}

func (j *JwtService) GenerateToken(userId string) string {
	claims := &jwtCustomClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}


func (j *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
