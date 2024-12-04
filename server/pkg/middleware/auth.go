package middleware

import (
	"instagram/pkg/constants"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenFromHeader := ctx.GetHeader("Authorization")
		if len(tokenFromHeader) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		if !strings.EqualFold(tokenFromHeader[:7], "bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		username, err := verifyToken(tokenFromHeader[len("bearer "):])
		if err != nil {
			log.Println("error verifying token: ", err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("username", username)
	}
}

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
	if err != nil {
		log.Println("error parsing token: ", err.Error())
		return "", err
	}

	if !token.Valid {
		log.Println("invalid token")
		return "", constants.ErrInvalidToken
	}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		log.Println("error getting subject from token: ", err.Error())
		return "", err
	}

	return subject, nil
}
