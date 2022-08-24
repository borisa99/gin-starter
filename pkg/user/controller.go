package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(g *gin.RouterGroup, db *gorm.DB) {
	h := &userHandler{
		DB: db,
	}

	routes := g.Group("/user")

	routes.GET("/me", h.Me)
}
