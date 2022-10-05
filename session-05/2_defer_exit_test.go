package session_5

import (
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("defer function start to execute")
	fmt.Println("Hai everyone")
	fmt.Println("welcome back to GO learning center")
}

func TestDeferFunc(t *testing.T) {
	callDeferFunc()
	fmt.Println("Hai everyone")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func TestExit(t *testing.T) {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
