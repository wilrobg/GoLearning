package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v1 := 2

	if v1%2 == 0 {
		fmt.Printf("Number %d is pair", v1)
	} else {
		fmt.Printf("Number %d is unpair", v1)
	}

	// With and
	if v1 == 2 && true {
		fmt.Println("Yes, it is a two")
	}

	// With or
	if v1 == 2 && true {
		fmt.Println("Yes, it is a two")
	}

	// Convert text to numeric and handled errors
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)

	num := 0

	switch {
	case num > 0:
		fmt.Println("Number is positive")
	case num == 0:
		fmt.Println("Number Zero")
	default:
		fmt.Println("Number is Negative")
	}
}
