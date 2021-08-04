// print a comment here

package main

import (
	"fmt"
	"strings"
)

func main() {
	txt := `
	Apple Bana
	Apple Orange
	orange bana
	new
	`
	txt = strings.ToLower(txt)
	words := strings.Fields(txt)

	var wc map[string]int
	wc = make(map[string]int)

	// wc2 := map[string]int{} // another way
	for _, word := range words {
		// _, ok := wc[word]
		// if !ok {
		// 	wc[word] = 1
		// } else {
		// 	wc[word]++
		// } // commented part not required
		wc[word]++
	}
	fmt.Println(wc)
}
