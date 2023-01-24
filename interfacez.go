package main

import "fmt"

type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangule struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (s rectangule) area() float64 {
	return s.base * s.height
}

func calculateArea(f figure2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	mySquare := square{base: 2}
	myRectangule := rectangule{base: 2, height: 4}

	calculateArea(mySquare)
	calculateArea(myRectangule)
}
