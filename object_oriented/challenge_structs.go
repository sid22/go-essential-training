// print a comment here
package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

type Square struct {
	center Point
	length int
}

func (sq Square) Area() int {
	return sq.length * sq.length
}

func (sq *Square) Move(dx int, dy int) {
	sq.center.Move(dx, dy)
}

func NewSquare(x int, y int, len int) (*Square, error) {
	if len < 0 {
		return nil, fmt.Errorf("len cannot be less than 0")
	}
	sq := &Square{Point{x, y}, len}
	return sq, nil
}

func main() {
	sq1, err := NewSquare(2, 2, 3)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("%+v \n", sq1)
	}
	sq1.Move(1, 2)
	fmt.Printf("%+v \n", sq1)
	ar := sq1.Area()
	fmt.Println("area is ", ar)
}
