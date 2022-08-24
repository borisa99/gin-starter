package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(g *gin.RouterGroup, db *gorm.DB) {
	h := &authHandler{
		DB: db,
	}

	routes := g.Group("/auth")

	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)

}
