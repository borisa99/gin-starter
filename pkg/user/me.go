package user

import "github.com/gin-gonic/gin"

func (h userHandler) Me(c *gin.Context) {

	c.JSON(200, "OK")
}
