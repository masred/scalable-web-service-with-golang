package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name": "Masred",
		"email": "Masred@gmail.com",
		"age": 22
	}
	`

	var resultStruct Employee

	var err = json.Unmarshal([]byte(jsonString), &resultStruct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", resultStruct.FullName)
	fmt.Println("email:", resultStruct.Email)
	fmt.Println("age:", resultStruct.Age)

	var resultMap map[string]interface{}

	err = json.Unmarshal([]byte(jsonString), &resultMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", resultMap["full_name"])
	fmt.Println("email:", resultMap["email"])
	fmt.Println("age:", resultMap["age"])

	var jsonStringSlice = `[
	{
		"full_name": "Masred",
		"email": "Masred@gmail.com",
		"age": 22
	},
	{
		"full_name": "Masred",
		"email": "Masred@gmail.com",
		"age": 22
	}
	]
	`
	var resultSliceOfStruct []Employee

	err = json.Unmarshal([]byte(jsonStringSlice), &resultSliceOfStruct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range resultSliceOfStruct {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
