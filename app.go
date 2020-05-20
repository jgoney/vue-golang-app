package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getLuckyNumber(c *gin.Context) {
	n := rand.Int31n(100)
	c.JSON(http.StatusOK, gin.H{"number": n})
}

func main() {
	router := gin.New()

	router.LoadHTMLGlob("dist/index.html")

	router.GET("/api/lucky", getLuckyNumber)
	router.GET("/", indexPage)

	// Set a PORT environment variable for more flexibility
	router.Run(":8869")
}
