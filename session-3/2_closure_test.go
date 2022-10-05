package session_3

import (
	"fmt"
	"strings"
	"testing"
)

func TestClosure(t *testing.T) {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))
}

func TestClosureIIFE(t *testing.T) {
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)
}

func TestClosureAsAReturnValue(t *testing.T) {
	var studentLists = []string{"Masred", "Dinda", "Sandi", "Mario", "Santi", "Makro"}

	var find = findStudent(studentLists)

	fmt.Println(find("Dinda"))
	fmt.Println(findStudent(studentLists)("sandi"))
}

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.EqualFold(v, s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

type isOddNum func(int) bool

func TestClosureCallback(t *testing.T) {
	var numbers = []int{2, 5, 8, 10, 3, 99, 23}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}

func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
