// print a comment here

package main

import (
	"fmt"
)

func main() {
	x := 10.0 // auto assign
	fmt.Println("-------")

	if x > 5 {
		fmt.Println("x is big")
	}
	fmt.Println("-------")

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x > 8 || x < 4 {
		fmt.Println("x is not that right")
	}
	fmt.Println("-------")

	switch x {
	case 1:
		fmt.Println("x is 1")
	default:
		fmt.Println("x not known")
	}
	fmt.Println("-------")

	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 5:
		fmt.Println("x is big")
	default:
		fmt.Println("x not known here also")
	}

}
