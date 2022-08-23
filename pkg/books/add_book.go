package books

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h bookHandler) AddBook(c *gin.Context) {

	var body AddBookRequestBody

	// getting request body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: body.Description,
	}

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
