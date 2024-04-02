package auth

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/utils"
)

var jwtSecret = []byte("your_jwt_secret")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, utils.NewErrorResponse("000x74", "Authorization header is required", "Authorization header is required"))
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, utils.NewErrorResponse("000x75", "Invalid authorization header", "Invalid authorization header"))
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, utils.NewErrorResponse("000x76", "Invalid JWT", "Invalid JWT"))
			c.Abort()
			return
		}

		c.Next()
	}
}
