package session_2

import (
	"fmt"
	"testing"
)

func TestLoopingOverString(t *testing.T) {
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf(" index: %d, byte: %d", i, city[i])
	}
	var jakarta []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println()
	fmt.Println(string(jakarta))
}
