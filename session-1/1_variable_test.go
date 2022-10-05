package session_1

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	var name string
	age := 22
	name = "Masred"

	fmt.Printf("%T, %T", name, age)
}

func TestMultipleVariable(t *testing.T) {
	var name, age, country = "Masred", 22, "Indonesia"
	first, second, third := 1, "2", 3

	fmt.Println(name, age, country)
	fmt.Println(first, second, third)
}

func TestUnderscoreVariable(t *testing.T) {
	var firstVariable string
	var name, age, country = "Masred", 22, "Indonesia"
	_, _, _, _ = firstVariable, name, age, country
}

func TestPrintf(t *testing.T) {
	var name = "Masred"
	age := 22
	country := "Indonesia"
	fmt.Printf("Hello my name is %s, i am %d years old, and i am from %s", name, age, country)
}
