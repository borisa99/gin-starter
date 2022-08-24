package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(g *gin.RouterGroup, db *gorm.DB) {
	h := &bookHandler{
		DB: db,
	}

	routes := g.Group("/books")

	routes.GET("/:id", h.GetBook)
	routes.GET("/", h.GetBooks)
	routes.POST("/", h.AddBook)
	routes.PUT("/", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)

}
