package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	Year  int
}

func PrintMessage() {
	fmt.Println("Hello")
}

func privatePrintMessage() {
	fmt.Println("Hello")
}
