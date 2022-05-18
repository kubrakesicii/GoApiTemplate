package middleware

import (
	"goapitemplate/pkg/helper"
	"goapitemplate/pkg/token"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorize JWT validate the token user giveni return 401 if not valid
func Authorize(jwtService token.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.Split(ctx.GetHeader("Authorization"), "Bearer")[1]
		if authHeader == "" {
			helper.Unauthorized()
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			helper.InternalServerError()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			// log claims
			log.Println("Claim userId : ", claims["user_id"])
		} else {
			helper.Unauthorized()
			return
		}
		ctx.Next()
	}
}
