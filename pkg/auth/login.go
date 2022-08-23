package auth

import (
	"net/http"

	"github.com/borisa99/gin-starter/pkg/auth/token"
	"github.com/borisa99/gin-starter/pkg/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h authHandler) Login(c *gin.Context) {
	var b LoginRequest

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var u models.User

	// Find user in DB
	if result := h.DB.Where(&models.User{Email: b.Email}).First(&u); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	}

	// Verify password hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(b.Password)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Genereate JWT
	jwt, err := token.GenerateToken(u.ID.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, jwt)
}
