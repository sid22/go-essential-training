// print a comment here

package main

import "fmt"

func main() {

	mps := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Printf(" nos is %v, type %T ", mps, mps)

	fmt.Println(mps["c"])

	fmt.Println(mps["d"])

	fmt.Println("-------")
	value, ok := mps["d"]
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(value)
	}

	mps["d"] = 4

}
