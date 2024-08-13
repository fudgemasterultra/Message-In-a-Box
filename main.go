package main

import (
	"fudgemasterultra/messageinabox/backend/databse"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func homePage(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	databse.Connect()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/index", homePage)
	router.Run("localhost:8090")
}
