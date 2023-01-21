func main() {
	//Consts declaration
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Variables declaration
	base := 12          //Without variable declaration
	var altura int = 14 //Assigning variable during declaration
	var area int        //Go wont compile if variables is not in use.

	fmt.Println(base, altura, area)

	//Zero values
	//Go assign default values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Calculating squareArea
	const baseSquare = 10
	areaSquare := baseSquare * baseSquare

	fmt.Println("Area Square:", areaCuadrado)
}```