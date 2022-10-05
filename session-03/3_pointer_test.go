package session_3

import (
	"fmt"
	"strings"
	"testing"
)

/*
* Kita dapat mendapatkan alamat memori dari variable biasa dengan menggunakan tanda ampersand &
* Kita juga dapat mendapatkan nilai asli dari variable pointer dengan cara menggunakan tanda asterisk *
 */

func TestPointer(t *testing.T) {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memory address", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memory address) :", secondNumber)
}

func TestPointerChanginValueThroughPointer(t *testing.T) {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memory address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memory address) :", secondPerson)

	fmt.Print("\n", strings.Repeat("#", 30), "\n\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memory address) :", &firstPerson)
	fmt.Println("secondPerson (value) :", *secondPerson)
	fmt.Println("secondPerson (memory address) :", secondPerson)
}

func TestPointerAsAParameter(t *testing.T) {
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
