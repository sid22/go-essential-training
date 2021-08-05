// print a comment here
package main

import "fmt"

type Trade struct {
	Symbol string // uppercase letter means public, else private
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func (t Trade) IncrNp(val float64) { // non ptr method won't change the object of the struct
	t.Price += val
}

func (t *Trade) Incr(val float64) {
	t.Price += val
}

func main() {
	t1 := Trade{"appl", 5, 10.5, true}
	fmt.Println(t1.Value())

	t1.IncrNp(4.1)
	fmt.Printf("%+v \n", t1)

	t1.Incr(4.1)
	fmt.Printf("%+v \n", t1)

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
