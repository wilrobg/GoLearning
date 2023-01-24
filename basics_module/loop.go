package main

import "fmt"

func main() {
	//Contidional for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	couterForever := 0
	for {
		fmt.Println(couterForever)
		couterForever++
	}
}
