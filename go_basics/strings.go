// print a comment here

package main

import "fmt"

func main() {
	book := "sample str"
	fmt.Println(book)
	fmt.Println("-------")
	fmt.Println(len(book))
	fmt.Println("-------")
	fmt.Printf("book[0] = %v type of book[0] %T and type of book %T \n ", book[0], book[0], book)

	// strings are immutable
	// book[0] = 116

	fmt.Println(book[2:6])
	fmt.Println(book[2:])
	fmt.Println("-------")

	fmt.Println("test " + book[2:])

	fmt.Println("-------")

	fmt.Println("strings are unicode so 😛 add those emojis")

	fmt.Println("-------")

	mls := `
	test 1
	test 2
	`

	fmt.Println(mls)
}
