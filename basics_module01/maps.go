package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jhon"] = 14
	m["Doe"] = 20

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	//Finding a value
	value, ok := m["Jhon1"]
	fmt.Println(value, ok)
}
