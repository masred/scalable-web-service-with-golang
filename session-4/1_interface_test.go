package session_4

import (
	"fmt"
	"math"
	"testing"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) area() float64 {
	return r.Height * r.Width
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func printResult(t string, s Shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func TestInterface(t *testing.T) {
	var c1 Shape = Circle{Radius: 5}
	var r1 Shape = Rectangle{Width: 3, Height: 2}

	printResult("Rectangle", r1)
	printResult("Circle", c1)
}

func (c Circle) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.Radius, 3)
}

func TestInterfaceTypeAssertion(t *testing.T) {
	var c1 Shape = Circle{Radius: 5}

	value, ok := c1.(Circle)

	if ok {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle Volume:%f\n", value.Volume())
	}
}
