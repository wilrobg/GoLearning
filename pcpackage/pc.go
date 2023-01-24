package pcpackage

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) ping() {
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) DuplicateRamWithOutPointer() {
	myPc.Ram = myPc.Ram * 2
}
