package session_5

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelImplementation(t *testing.T) {
	c := make(chan string)

	go introduce("Masred", c)
	go introduce("Clara", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

func TestChannelAnonymousFunction(t *testing.T) {

	c := make(chan string)

	students := []string{"Masred", "Clara"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 2; i++ {
		print(c)
	}
	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}

func TestChannelDirection(t *testing.T) {
	c := make(chan string)

	students := []string{"Masred", "Clara"}

	for _, v := range students {
		go intro(v, c)
	}

	for i := 1; i <= 2; i++ {
		println(c)
	}

	close(c)
}

func intro(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

func println(c <-chan string) {
	fmt.Println(<-c)
}

func TestChannelUnbuffered(t *testing.T) {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("goroutine anonymous func start sending data into channel")
		c <- 10
		fmt.Println("goroutine func after sending data into channel")
	}(c1)

	fmt.Println("sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	fmt.Println("start receiving data")
	d := <-c1
	fmt.Println("received data:", d)

	close(c1)
	time.Sleep(time.Second)
}

func TestChannelBuffered(t *testing.T) {
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("goroutine anonymous func start sending data into channel")
			c <- i
			fmt.Println("goroutine func after sending data into channel")
		}

		close(c1)

	}(c1)

	fmt.Println("sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	for v := range c1 {
		fmt.Println("received value from channel:", v)
	}
}

func TestChannelSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "Salut"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
