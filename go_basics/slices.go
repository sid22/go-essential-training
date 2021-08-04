// print a comment here

package main

import "fmt"

func main() {

	loons := []string{"bugs", "daffy", "test"}
	fmt.Printf(" loons is %v, type %T ", loons, loons)

	fmt.Println("-------")
	fmt.Println(len(loons))

	fmt.Println("-------")
	fmt.Println(loons[0])

	fmt.Println("-------")
	fmt.Println(loons[1:])

	fmt.Println("-------")
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i, loon := range loons {
		fmt.Println(i, loon)
	}

	fmt.Println("-------")
	for _, loon := range loons {
		fmt.Println(loon)
	}

	fmt.Println("-------")
	loons = append(loons, "test2")
	fmt.Println(loons)
}
