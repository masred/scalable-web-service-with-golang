package session_4

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroupImplementation(t *testing.T) {
	fruits := []string{"apple", "mango", "durian", "rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index = %d, fruit => %s\n", index, fruit)
	wg.Done()
}
