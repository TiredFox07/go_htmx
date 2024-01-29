package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/htmx.min.js", func(c *gin.Context) {
		c.File("htmx.min.js")
	})
	router.GET("/update", func(c *gin.Context) {
		c.HTML(http.StatusOK, "update.tmpl", gin.H{})
	})
	router.Run(":8080")
}
