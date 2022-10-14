package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
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
	var data Json

	go func() {
		for range time.Tick(time.Second * 15) {
			rand.Seed(time.Now().UnixNano())

			data.Status.Water = rand.Intn(100)
			data.Status.Wind = rand.Intn(100)

			file, _ := os.Create("status.json")
			jsonData, _ := json.Marshal(data)

			file.WriteString(string(jsonData))
			file.Close()
		}
	}()

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		var data Json
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

		dataMap := map[string]interface{}{
			"Water":       data.Status.Water,
			"WaterStatus": waterStatus,
			"Wind":        data.Status.Wind,
			"WindStatus":  windStatus,
		}

		t.Execute(ctx.Writer, dataMap)
	})
	router.Run()
}
