package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"status": "ok",
		})
	})

	router.GET("/message", func(c *gin.Context) {
		c.HTML(http.StatusOK, "powered_by.html", gin.H{})
	})
	router.Run()
}
