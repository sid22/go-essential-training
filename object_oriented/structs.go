// print a comment here
package main

import "fmt"

type Trade struct {
	Symbol string // uppercase letter means public, else private
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"appl", 1, 10.5, true}
	fmt.Println(t1)

	fmt.Printf("%+v \n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "msft",
		Volume: 10,
		Price:  14.0,
		Buy:    true,
	}

	fmt.Printf("%+v \n", t2)

	t3 := Trade{}
	fmt.Printf("%+v \n", t3)

}
