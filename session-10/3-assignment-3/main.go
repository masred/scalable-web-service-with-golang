package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(ctx *gin.Context) {
		var data Json
		rand.Seed(time.Now().UnixNano())

		data.Status.Water, data.Status.Wind = rand.Intn(100), rand.Intn(100)

		file, _ := os.Create("status.json")
		jsonData, _ := json.Marshal(data)

		file.WriteString(string(jsonData))
		file.Close()

		dataFile, _ := os.ReadFile("status.json")
		json.Unmarshal(dataFile, &data)

		var waterStatus string
		if data.Status.Water > 8 {
			waterStatus = "bahaya"
		} else if data.Status.Water >= 6 {
			waterStatus = "siaga"
		} else {
			waterStatus = "aman"
		}

		var windStatus string
		if data.Status.Wind > 15 {
			windStatus = "bahaya"
		} else if data.Status.Water >= 7 {
			windStatus = "siaga"
		} else {
			windStatus = "aman"
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Water":       data.Status.Water,
			"WaterStatus": waterStatus,
			"Wind":        data.Status.Wind,
			"WindStatus":  windStatus,
		})
	})

	router.Run()
}
