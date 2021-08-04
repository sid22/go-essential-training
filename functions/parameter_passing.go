// print a comment here

package main

import (
	"fmt"
)

func doubleA(values []int, i int) {
	values[i] *= 2
}

func double(i int) {
	i *= 2
}

func doublePtr(i *int) {
	*i *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleA(values, 2)
	fmt.Println(values) // go passes slice as reference here

	n := 10
	double(n) // go passes number copy here not reference
	fmt.Println(n)

	doublePtr(&n) // go passes number pointer here
	fmt.Println(n)

}
