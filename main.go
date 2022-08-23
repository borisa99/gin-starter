package main

import (
	"github.com/borisa99/gin-starter/pkg/auth"
	"github.com/borisa99/gin-starter/pkg/books"

	"github.com/borisa99/gin-starter/pkg/common/db"
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

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	books.RegisterRoutes(r, h)
	auth.RegisterRoutes(r, h)
	r.Run(port)
}
