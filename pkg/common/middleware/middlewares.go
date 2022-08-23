package middleware

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/auth/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.Request.Header.Get("Authorization")
		_, err := token.VerifyToken(bearerToken)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
