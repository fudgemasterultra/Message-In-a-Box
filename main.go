package main

import (
	"fudgemasterultra/messageinabox/backend/databse"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func homePage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	databse.Connect()
	databse.GetMessageBox(1);
	router := gin.Default()
	routes(router)
	router.LoadHTMLFiles("./node/build/index.html")
	router.Static("/_app", "./node/build/_app")
	router.Static("/static", "./static")
	router.GET("/", homePage)
	router.Run("localhost:8090")
}
