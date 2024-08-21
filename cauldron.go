package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func connectToDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./cauldron")
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}

func main() {
	db := connectToDb()
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
	defer db.Close()
}
