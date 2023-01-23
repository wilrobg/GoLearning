package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

//It is a good practice avoiding multiple types declaration for the same type (a,b) example
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func calculateTriangleArea(height, base int) int {
	return base * height / 2
}

func main() {
	normalFunction("Hello world")
	tripleArgument(1, 2, "Hello")
	fmt.Println(returnValue(2))

	value1, value2 := doubleReturn(3)
	fmt.Println(value1, value2)
	fmt.Println(doubleReturn(3))

	area := calculateTriangleArea(2, 3)
	fmt.Printf("Height: %d Base: %d Area: %d", 2, 3, area)
}
