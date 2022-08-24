package middleware

import (
	"net/http"
	"strings"

	"github.com/borisa99/gin-starter/pkg/auth/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.Request.Header.Get("Authorization")
		formatedToken := strings.ReplaceAll(bearerToken, "Bearer ", "")
		ok := token.VerifyToken(formatedToken)

		if !ok {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
