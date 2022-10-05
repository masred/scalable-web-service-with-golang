package session_2

import (
	"fmt"
	"testing"
)

func TestSliceMakeFunction(t *testing.T) {
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)
}

func TestSliceAppendFunction(t *testing.T) {

	var fruits = make([]string, 3)
	fruits = append(fruits, "apple", "banana", "mango")
	fmt.Printf("%#v", fruits)
}

func TestSliceAppendFunctionWithEllipsis(t *testing.T) {
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v", fruits1)
}

func TestSliceCopyFunction(t *testing.T) {
	var fruits1 = []string{"apple", "banaan", "mango"}
	var fruits2 = []string{"durian", "pineaple", "starfruit"}

	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elemets =>", nn)
}

func TestSliceSlicing(t *testing.T) {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineaple"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%3v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%v\n", fruits3)

	var fruits5 = fruits1[:]
	fmt.Printf("%#v\n", fruits5)
}

func TestSliceCombiningSlicingAndAppend(t *testing.T) {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineaple"}
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v\n", fruits1)
}

func TestSliceBackingArray(t *testing.T) {
	var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}
	var fruits2 = fruits1[2:4]

	fruits2[0] = "rambutan"

	fmt.Println("fruit1 =>", fruits1)
	fmt.Println("fruit =>", fruits2)
}
