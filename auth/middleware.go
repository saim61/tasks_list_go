package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/saim61/tasks_list_go/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	return func(g *gin.Context) {
		// JWT validation logic
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			g.JSON(http.StatusUnauthorized, utils.NewErrorResponse("000x74", "Authorization header is required", "Authorization header is required"))
			g.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			g.JSON(http.StatusUnauthorized, utils.NewErrorResponse("000x75", "Invalid authorization header", "Invalid authorization header"))
			g.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			g.JSON(http.StatusUnauthorized, utils.NewErrorResponse("000x76", err.Error(), "JWT not valid anymore"))
			g.Abort()
			return
		}

		g.Next()
	}
}
