package session_4

import (
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Masred"}
}

func TestEmptyInterfaceTypeAssertion(t *testing.T) {
	var v interface{} = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}

func TestEmptyInterfaceMapSlice(t *testing.T) {
	rs := []interface{}{1, "Masred", true, 2, "Clara", 1.2}

	rm := map[string]interface{}{
		"Name":   "Masred",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
