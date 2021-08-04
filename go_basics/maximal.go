// print a comment here

package main

import "fmt"

func main() {

	nos := []int{4, 12, 34, 56, 87, 1, 789, 997}
	fmt.Printf(" nos is %v, type %T ", nos, nos)

	fmt.Println("-------")
	max_no := -1
	for _, no := range nos {
		if no > max_no {
			max_no = no
		}
	}
	fmt.Println("max no is ", max_no)

}
