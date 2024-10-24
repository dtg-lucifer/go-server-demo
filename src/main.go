package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
    c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	})

	admin := r.Group("/admin")

	admin.POST("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
