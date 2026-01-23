package middleware

import (
	"net/http"
	"os"
	"simple_golang_api_with_authentication/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "401 token invalid 0.1",
			})

			ctx.Abort() //stop request
			return
		}

		//BEARER TOKEN
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "401 token invalid 0.2",
			})
			ctx.Abort() //stop request
			return
		}

		tokenStr := parts[1]
		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "401 token invalid 0.3",
			})
			ctx.Abort() //stop request
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
