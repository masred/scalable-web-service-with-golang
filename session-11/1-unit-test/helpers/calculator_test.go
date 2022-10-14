package helpers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("Result should be 20")
	}
}

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Fail()
	}

	fmt.Println("TestFailSum Execution will be stopped")
}

func TestFailNowSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestFailNowSum Execution will be stopped")
}
