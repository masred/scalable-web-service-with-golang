package session_1

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	var first uint8 = 89
	var second int8 = -12
	var third float32 = 3.65

	fmt.Printf("%d is %T\n", first, first)
	fmt.Printf("%d is %T\n", second, second)
	fmt.Printf("%f is %T\n", third, third)
}

func TestBool(t *testing.T) {
	var condition bool = true
	fmt.Printf("is if permitted? %t \n", condition)
}

func TestString(t *testing.T) {
	var message string = "Hi"
	var anotherMessage = `Welcome to "Hacktiv8".
	Nice to meet you:).
	Let's GO!`

	fmt.Print(message)
	fmt.Print(anotherMessage)
}
