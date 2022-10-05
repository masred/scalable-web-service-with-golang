package session_4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectIdentifyingDataTypeValue(t *testing.T) {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Variable type :", reflectValue.Type())
	fmt.Println("Variable value :", reflectValue.Interface())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Variable value :", reflectValue.Int())
	}
}

type Student struct {
	Name  string
	Grade int
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func TestReflectionIdentifyingMethod(t *testing.T) {
	var s1 = &Student{Name: "john wick", Grade: 100}
	fmt.Println("Name :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("Name :", s1.Name)
}
