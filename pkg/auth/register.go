package auth

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

func (h authHandler) Register(c *gin.Context) {

	var b RegisterRequest

	if err := c.ShouldBind(&b); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	u := models.User{
		FirstName:    b.FirstName,
		LastName:     b.LastName,
		Email:        b.Email,
		PasswordHash: string(bytes),
	}

	if result := h.DB.Create(&u); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, &u)
}
