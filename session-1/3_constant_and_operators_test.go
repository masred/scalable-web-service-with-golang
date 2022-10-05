package session_1

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	const fullName string = "Masred"
	fmt.Printf("Hi %s", fullName)
}

func TestArithmeticOperators(t *testing.T) {
	var first = 2 + 2*3
	var second = (2 + 2) * 3
	fmt.Println(first, second)
}

func TestRelationalOperators(t *testing.T) {
	var first bool = 2 < 3
	var second bool = "joy" == "Joy"
	var third bool = 10 != 2.3
	var fourth bool = 21 <= 11

	fmt.Println("first condition is ", first)
	fmt.Println("second condition is ", second)
	fmt.Println("third condition is ", third)
	fmt.Println("fourth condition is ", fourth)
}

func TestLogicalOperators(t *testing.T) {
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
