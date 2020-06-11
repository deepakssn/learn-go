package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

type shape interface {
	area() float64
}

func main() {
	t := triangle{
		base:   10,
		height: 10,
	}
	s := square{
		side: 10,
	}

	printArea(t)
	printArea(s)

}

func printArea(sh shape) {
	fmt.Println(sh.area())
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (s square) area() float64 {
	return s.side * s.side
}
