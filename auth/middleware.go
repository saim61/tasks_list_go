package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/utils"
)

var jwtSecret = []byte("your_jwt_secret")

func AuthMiddleware() gin.HandlerFunc {

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
			g.JSON(http.StatusUnauthorized, utils.NewErrorResponse("000x76", "Invalid JWT", "JWT not valid anymore"))
			g.Abort()
			return
		}

		g.Next()
	}
}
