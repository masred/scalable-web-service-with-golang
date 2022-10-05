package session_3

import (
	"fmt"
	"strings"
	"testing"
)

type Employee struct {
	Name     string
	Age      int
	Division string
}

func TestStructGivenValue(t *testing.T) {
	var employee Employee

	employee.Name = "Masred"
	employee.Age = 22
	employee.Division = "Backend Developer"

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.Division)
}

func TestStructInitialized(t *testing.T) {
	var employee1 = Employee{}
	employee1.Name = "Masred"
	employee1.Age = 23
	employee1.Division = "Backend Developer"

	var employee2 = Employee{
		Name:     "Clara",
		Age:      17,
		Division: "Frondend Developer",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}

func TestStructPointerToAStruct(t *testing.T) {
	var employee1 = Employee{
		Name:     "Clara",
		Age:      17,
		Division: "Frondend Developer",
	}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.Name)
	fmt.Println("Employee2 name:", employee2.Name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.Name = "Clarissa"

	fmt.Println("Employee1 name:", employee1.Name)
	fmt.Println("Employee2 name:", employee2.Name)
}

type Person struct {
	Name string
	Age  int
}

type Student struct {
	School string
	Person Person
}

func TestStructEmbed(t *testing.T) {
	var student1 = Student{}

	student1.Person.Name = "Clara"
	student1.Person.Age = 17
	student1.School = "SMA"

	fmt.Printf("%+v", student1)
}

func TestStructAnonymousStruct(t *testing.T) {
	var student1 = struct {
		Person Person
		School string
	}{}

	student1.Person = Person{
		Name: "Clara",
		Age:  17,
	}
	student1.School = "SMA"

	var student2 = struct {
		Person Person
		School string
	}{
		Person: Person{
			Name: "Clarissa",
			Age:  17,
		},
		School: "SMA",
	}

	fmt.Printf("Student1: %+v\n", student1)
	fmt.Printf("Student2: %+v\n", student2)
}

func TestStructSliceOfStruct(t *testing.T) {
	var people = []Person{
		{
			Name: "Clara",
			Age:  17,
		},
		{
			Name: "Masred",
			Age:  22,
		},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}

func TestStructSliceOfAnonymousStruct(t *testing.T) {
	var student = []struct {
		Person Person
		School string
	}{
		{
			Person: Person{
				Name: "Masred",
				Age:  22,
			},
			School: "STM",
		},
		{
			Person: Person{
				Name: "Clara",
				Age:  17,
			},
			School: "SMA",
		},
	}

	for _, v := range student {
		fmt.Printf("%+v\n", v)
	}
}
