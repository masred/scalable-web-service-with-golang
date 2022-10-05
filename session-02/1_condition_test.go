package session_2

import (
	"fmt"
	"testing"
)

func TestIfElseCondition(t *testing.T) {
	var currentYear = 2022

	if age := currentYear - 2000; age < 17 {
		fmt.Println("You don't have permission to create driver license")
	} else {
		fmt.Println("You can create driver license")
	}
}

func TestSwitchCondition(t *testing.T) {
	var score = 8
	switch score {
	case 8:
		fmt.Println("perfect")
	case 9:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

func TestSwitcConditionhWithRelationalOperator(t *testing.T) {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
}

func TestSwitchFallthrough(t *testing.T) {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("it's ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

}

func TestNestedCondition(t *testing.T) {
	var score = 10
	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
