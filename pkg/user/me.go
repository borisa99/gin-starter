package user

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Me struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

func (h userHandler) Me(c *gin.Context) {
	uId := c.GetString("uId")

	var m Me

	if result := h.DB.Model(&models.User{}).Find(&m, "id = ?", uId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(200, m)
}
