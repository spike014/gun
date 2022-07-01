package main

import (
	"net/http"
	"test/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		app.Logs()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("....panic....")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
