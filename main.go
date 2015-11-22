package main

import (
	"encoding/json"
	"github.com/bertoort/sugoku/puzzle"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func parse(j string) puzzle.Puzzle {
	board := []byte(j)
	var b [9][9]int
	json.Unmarshal(board, &b)
	return puzzle.New(b)
}

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
	r.POST("/solve", func(c *gin.Context) {
		j := c.PostForm("board")
		s := parse(j)
		s.Grade()
		err := s.SlowSolve()
		v := s.Validate()
		sol, stat, dif := s.Display()
		if err && v {
			stat = "solved"
		} else if err {
			stat = "broken"
		}
		c.JSON(http.StatusOK, gin.H{
			"solution":   sol,
			"status":     stat,
			"difficulty": dif,
		})
	})

	r.POST("/grade", func(c *gin.Context) {
		j := c.PostForm("board")
		s := parse(j)
		s.Grade()
		_, _, dif := s.Display()
		c.JSON(http.StatusOK, gin.H{
			"difficulty": dif,
		})
	})

	r.POST("/validate", func(c *gin.Context) {
		j := c.PostForm("board")
		s := parse(j)
		v := s.Validate()
		_, stat, _ := s.Display()
		err := s.SlowSolve()
		_, newStat, _ := s.Display()
		if err && v {
			stat = "solved"
		} else if err || newStat == "unsolvable" {
			stat = "broken"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": stat,
		})
	})

	r.Run(":" + port)
}
