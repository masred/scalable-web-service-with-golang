package session_2

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var person map[string]string = map[string]string{}

	person["name"] = "Masred"
	person["age"] = "23"
	person["address"] = "Jalan Sudirman"

	fmt.Println("name : ", person["name"])
	fmt.Println("age : ", person["age"])
	fmt.Println("address : ", person["address"])
}

func TestMapLooping(t *testing.T) {
	var person = map[string]string{
		"name":    "Masred",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}

func TestMapDeletingValue(t *testing.T) {
	var person = map[string]string{
		"name":    "Masred",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("Before deleting : ", person)

	delete(person, "age")

	fmt.Println("After deleting : ", person)
}

func TestMapDetectingValue(t *testing.T) {
	var person = map[string]string{
		"name":    "Masred",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}
}

func TestCombiningSliceWithMap(t *testing.T) {
	var people = []map[string]string{
		{"name": "Masred", "age": "23"},
		{"name": "Nanda", "age": "29"},
		{"name": "Mailo", "age": "13"},
	}

	for i, person := range people {
		fmt.Printf("Index : %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
