package main

import (
	"fudgemasterultra/messageinabox/backend/databse"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func homePage(c *gin.Context) {
	c.HTML(200, "pageCount.html", gin.H{
		"title": "Home Page",
		"id":    c.Param("id"),
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	databse.Connect()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", homePage)
	router.Run("localhost:8090")
}
