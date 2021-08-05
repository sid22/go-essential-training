// print a comment here
package main

import "fmt"

type Trade struct {
	Symbol string // uppercase letter means public, else private
	Volume int
	Price  float64
	Buy    bool
}

// to create a constructor in Go, we create a method starting with New{Struct_name}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be empty")
	}

	if volume < 0 {
		return nil, fmt.Errorf("volume cannot be less than 0")
	}
	trade := &Trade{symbol, volume, price, buy}
	return trade, nil
}

func main() {
	t1, err := NewTrade("appl", 5, 10.5, true)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("%+v \n", t1)

	}

	t2, err2 := NewTrade("appl2", -1, 10.5, true)
	if err2 != nil {
		fmt.Println("error")
	} else {

		fmt.Printf("%+v \n", t2)
	}
	// fmt.Println(t1.Symbol)

	// t2 := Trade{
	// 	Symbol: "msft",
	// 	Volume: 10,
	// 	Price:  14.0,
	// 	Buy:    true,
	// }

	// fmt.Printf("%+v \n", t2)

	// t3 := Trade{}
	// fmt.Printf("%+v \n", t3)

}
