package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type JSON struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	var data JSON

	for range time.Tick(time.Second * 15) {
		rand.Seed(time.Now().UnixNano())

		data.Status.Water = rand.Intn(100)
		data.Status.Wind = rand.Intn(100)

		file, _ := os.Create("status.json")
		jsonData, _ := json.Marshal(data)

		file.WriteString(string(jsonData))
		file.Close()
	}

	app := gin.Default()
	app.Run()
}
