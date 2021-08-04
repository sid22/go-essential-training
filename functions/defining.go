// print a comment here

package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	a, b := 4, 2
	c := add(a, b)
	d, rem := divide(a, b)
	fmt.Println(c, d, rem)
}
