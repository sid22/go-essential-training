// print a comment here

package main

import "fmt"

func main() {
	num1 := 11
	num2 := 11

	num_mul := num1 * num2

	s := fmt.Sprintf("%d", num_mul)
	if s[0] == s[len(s)-1] {
		fmt.Println("even ended")
	} else {
		fmt.Println("not even ended")
	}

	count := 0
	for a := 1000; a < 9999; a++ {
		for b := 1000; b < 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// fmt.Println("even ended")
				count++
			}
		}
	}
	fmt.Printf("count is %d \n", count)
}
