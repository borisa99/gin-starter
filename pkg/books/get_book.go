package books

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h bookHandler) GetBook(c *gin.Context) {

	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}
