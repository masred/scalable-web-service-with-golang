package session_2

import (
	"fmt"
	"testing"
)

func TestLoopingFirstWay(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println("angka", i)
	}
}

func TestLoopingSecondWay(t *testing.T) {
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}
}

func TestLoopingThirdWay(t *testing.T) {
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 3 {
			break
		}
	}
}

func TestLoopingBreakAndContinue(t *testing.T) {
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}

func TestNestedLooping(t *testing.T) {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

func TestNestedLoopingLabel(t *testing.T) {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Looping ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
