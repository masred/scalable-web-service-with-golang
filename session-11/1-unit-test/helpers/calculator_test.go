package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTestifyFailSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestTestifyFailSum Execution stopped")
}

func TestTestifySum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestTestifySum execution stopped")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		Request  int
		Expected int
		ErrMsg   string
	}{
		{
			Request:  Sum(10, 10),
			Expected: 20,
			ErrMsg:   "Result has to be 20",
		},
		{
			Request:  Sum(20, 20),
			Expected: 40,
			ErrMsg:   "Result has to be 40",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.Expected, test.Request, test.ErrMsg)
		})
	}
}
