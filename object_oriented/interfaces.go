// print a comment here
package main

import (
	"fmt"
	"math"
)

// type Point struct {
// 	x int
// 	y int
// }

// func (p *Point) Move(dx int, dy int) {
// 	p.x += dx
// 	p.y += dy
// }

type Square struct {
	// center Point
	Length float64
}

func (sq Square) Area() float64 {
	return sq.Length * sq.Length
}

type Circle struct {
	// center Point
	Radius float64
}

func (cl Circle) Area() float64 {
	return math.Pi * cl.Radius * cl.Radius
}

func SumAread(shapes []Shape) float64 { // here we need Shape to be an interface i.e. it should be able to keep struct of type Circle and Square both
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// Here we define the interface based on the struct method
type Shape interface {
	Area() float64
}

func main() {
	sq := &Square{4}
	cl := &Circle{2}

	shapes := []Shape{sq, cl}
	sa := SumAread(shapes)
	fmt.Println("total area is ", sa)

}
