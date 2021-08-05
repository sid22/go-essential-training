// print a comment here
package main

import (
	// "errors"
	"fmt"
	// "os"
)

func safeVal(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error %s \n ", err)
		}
	}()
	return vals[index]
}

func main() {
	// vals := []int{1, 2, 3}
	// v := vals[10]
	// fmt.Println(v)

	// file, err := os.Open("anotexsitingfile")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// fmt.Println("opened")
	v := safeVal([]int{1, 2, 3}, 10)
	fmt.Println("v is ", v)
}
