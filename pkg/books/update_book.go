package books

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h bookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body UpdateBookRequestBody

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)
	c.JSON(http.StatusOK, &book)

}
