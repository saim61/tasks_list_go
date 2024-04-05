package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/utils"
)

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

		token, err := utils.GetJWKTokenFromJWT(authParts)
		if err != nil || !token.Valid {
			g.JSON(http.StatusUnauthorized, utils.NewErrorResponse("000x76", err.Error(), "JWT not valid anymore"))
			g.Abort()
			return
		}

		g.Next()
	}
}
