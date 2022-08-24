package main

import (
	"github.com/borisa99/gin-starter/pkg/auth"
	"github.com/borisa99/gin-starter/pkg/books"
	"github.com/borisa99/gin-starter/pkg/user"

	"github.com/borisa99/gin-starter/pkg/common/db"
	"github.com/borisa99/gin-starter/pkg/common/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	// Public router
	p := r.Group("/")

	// Auth router with attached middleware
	a := r.Group("/")
	a.Use(middleware.JwtAuthMiddleware())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	books.RegisterRoutes(p, h)
	auth.RegisterRoutes(p, h)
	user.RegisterRoutes(a, h)

	r.Run(port)
}
