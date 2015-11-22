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

// CORSMiddleware allows others to access the api
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT",
		)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	g := gin.Default()
	g.LoadHTMLGlob("templates/*.html")
	g.Static("/public", "public")

	g.Use(CORSMiddleware())

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "suGOku",
		})
	})

	g.GET("/board", func(c *gin.Context) {
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
	g.POST("/solve", func(c *gin.Context) {
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

	g.POST("/grade", func(c *gin.Context) {
		j := c.PostForm("board")
		s := parse(j)
		s.Grade()
		_, _, dif := s.Display()
		c.JSON(http.StatusOK, gin.H{
			"difficulty": dif,
		})
	})

	g.POST("/validate", func(c *gin.Context) {
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

	g.Run(":" + port)
}
