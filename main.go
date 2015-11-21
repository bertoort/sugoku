package main

import (
	// "encoding/json"
	// "fmt"
	// "github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
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
			"title": "suGOku",
		})
	})
	r.GET("/board", func(c *gin.Context) {
		dif := c.Query("difficulty")
		var b [9][9]int
		if dif == "random" {
			b = puzzle.GenRandom()
		} else if dif == "hard" {
			b = puzzle.Generate(dif)
		} else if dif == "medium" {
			b = puzzle.Generate(dif)
		} else if dif == "easy" {
			b = puzzle.Generate(dif)
		}
		c.JSON(http.StatusOK, gin.H{"board": b})
	})
	r.Run(":" + port)
}
