package main

import "fmt"

type Shape interface {
	GetArea() float64
}

type Triangle struct {
	Height float64
	Base float64
}
type Square struct {
	SideLength float64
}


func main() {
	t := Triangle{
		Height: 0.3,
		Base: 0.2,
	}
	s := Square{
		SideLength: 4,
	}
	printArea(s)
	printArea(t)
}

func (t Triangle) GetArea() float64{
	return 0.5 * t.Base * t.Height
}

func (s Square) GetArea() float64 {
	return s.SideLength * s.SideLength
}

func printArea(s Shape) {
	fmt.Println("Area: ", s.GetArea())
}