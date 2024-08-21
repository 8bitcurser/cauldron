package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func connectToDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./cauldron"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}

func main() {
	db := connectToDb()
	err := db.AutoMigrate(&Event{})
	if err != nil {
		log.Fatal(err)
	}
	//router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"status": "ok",
	//	})
	//})
	//
	//router.GET("/message", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "powered_by.html", gin.H{})
	//})
	//router.Run()
}
