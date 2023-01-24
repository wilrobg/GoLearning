package main

import (
	"fmt"
	p "pcpackage"
)

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := p.Pc{Brand: "Toyota", Ram: 16}
	fmt.Println(myPc)

	myPc.DuplicateRam()
	fmt.Println(myPc)

	myPc.DuplicateRamWithOutPointer()
	fmt.Println(myPc)
}
