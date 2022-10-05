package session_3

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestFunction(t *testing.T) {
	greet("Masred", "Jalan SUdirman")

	var names = []string{"Mas", "Jordan"}
	var printMsg = greetReturn("Hai", names)
	fmt.Println(printMsg)
}

func TestFunctionMultipleReturn(t *testing.T) {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println("Area : ", area)
	fmt.Println("Circumference : ", circumference)
}

func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

func greetReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}

func TestFunctionVariadic(t *testing.T) {
	studentLists := print("Masred", "Nanda", "jajang", "Marco")
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("Result : ", result)
	fmt.Printf("%v", studentLists)
	profile("Masred", "Pizza", "Seblak", "Ayam Geprek", "Ikan Koi")
}
