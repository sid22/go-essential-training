package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val2 := <-ch2:
		fmt.Printf("got %d from ch2\n", val2)
	}

	out := make(chan float64)
	fmt.Println("------")
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(210 * time.Millisecond):
		fmt.Println("timeout")
	}
}
