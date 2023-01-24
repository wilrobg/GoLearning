package main

import (
	"fmt"
	pk "mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Toyta"
	myCar.Year = 2020
	myCar.Print
	fmt.Println(myCar)

	pk.PrintMessage()
}
