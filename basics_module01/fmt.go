package main

import "fmt"

func main() {
	var num int32 = 10

	fmt.Printf("num type: %T\n", num)

	fmt.Printf("num: %d %v\n", num, num)

	message := fmt.Sprintf("Test %d", num)

	fmt.Println(message)
}
