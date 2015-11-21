package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	// "strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/public", "public")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "sudoku",
		})
	})
	r.Run(":" + port)
}
