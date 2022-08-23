package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &authHandler{
		DB: db,
	}

	routes := r.Group("/auth")

	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)

}
