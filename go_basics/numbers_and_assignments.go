// print a comment here

package main

import (
	"fmt"
)

func main() {
	// var x float64
	// var y float64

	// x = 1
	// y = 2
	x := 1.0 // auto assign
	y := 2.0

	fmt.Printf("x=%v, type of %T \n", x, x)
	fmt.Printf("y=%v, type of %T \n", y, y)

	var mean2 float64
	mean2 = (x + y) / 2

	fmt.Printf("mean2=%v, type of %T \n", mean2, mean2)
}
