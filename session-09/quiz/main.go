package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProgrammingLanguage struct {
	Language       string   `json:"language"`
	Appreared      int      `json:"appreared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object_oriented"`
	Relation       Relation `json:"relation"`
}

type Relation struct {
	InfluencedBy []string `json:"influenced_by"`
	Influences   []string `json:"influences"`
}

func main() {
	router := gin.Default()
	c := ProgrammingLanguage{
		Language:  "C",
		Appreared: 1972,
		Created: []string{
			"Dennis Ritchie",
		},
		Functional:     true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{
				"B",
				"ALGOL 68",
				"Assembly",
				"FORTRAN",
			},
			Influences: []string{
				"C++",
				"Objective-C",
				"C#",
				"Java",
				"JavaScript",
				"PHP",
				"Go",
			},
		},
	}

	router.GET("/", func(ctx *gin.Context) {
		if ctx.Request.Method != "GET" {
			fmt.Fprint(ctx.Writer, "Method Not Allowed")
		} else {
			fmt.Fprint(ctx.Writer, "Hello Go Developers")
		}
	})

	router.GET("/language", func(ctx *gin.Context) {
		if ctx.Request.Method != "GET" {
			fmt.Fprint(ctx.Writer, "Method Not Allowed")
		} else {
			ctx.JSON(http.StatusOK, c)
		}
	})

	router.GET("/palindrome", func(ctx *gin.Context) {
		text := ctx.Query("text")
		var reversedText string

		for i := len(text) - 1; i >= 0; i-- {
			reversedText = reversedText + string(text[i])
		}

		if text == reversedText {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Palindrome",
			})

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Not Palindrome",
			})
		}
	})

	router.Run()
}
