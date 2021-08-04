// print a comment here

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-------")
	y := 0 // like a while loop
	for y < 3 {
		fmt.Println(y)
		y++
	}

	fmt.Println("-------")
	y = 0 // like a while loop
	for {
		fmt.Println(y)
		y++
		if y > 2 {
			break
		}
	}

}
