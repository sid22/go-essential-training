// print a comment here

// go has a garbade collector, we don't need to manually release memory
// for other resource sockets etc we use defer
package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("cleaning up %s \n", name)
}

func worker() {
	defer cleanup("a") // will be called even if we have a panic

	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}
