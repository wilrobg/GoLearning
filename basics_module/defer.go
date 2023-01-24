package main

import "fmt"

func main() {
	defer fmt.Println("Hello")

	fmt.Println("World")

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Is 2")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
