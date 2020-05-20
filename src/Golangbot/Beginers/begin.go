package main

import (
	"fmt"
)

//Returning multiple values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

//NAmed Return Values
func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
func main() {
	// var a int = 89
	// b := 95
	// fmt.Println("value of a is", a, "and b is", b)
	// fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	// fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
	//constants internlly
	//var a int = 50
	// const b = "I love Go"
	// //b = "language"
	// fmt.Printf("Value of b %v Type of b%T", b, b)
	// var name = "Sam"
	// fmt.Printf("type %T value %v", name, name)

	// Assigning defaultname to customNamre
	// var defaultName = "Sam" //allowed
	// type myString string
	// var customName myString = "Sam" //allowed
	// customName = defaultName        //not allowed
	// fmt.Println("custom name:", customName, "default name:", defaultName)

	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a

	fmt.Printf("a: %T", a)
	fmt.Printf("\nintVar %T", intVar)
	fmt.Printf("\nint32Var %T", int32Var)
	fmt.Printf("\nfloat64Var %T", float64Var)

	area, perimeter := rectProps(10, 20)
	fmt.Println("area:", area)
	fmt.Println("Perimeter:", perimeter)

	//Call Named Return valuef function
	area1, perimeter1 := rectPropsNamed(10, 20)
	fmt.Println("area:", area1)
	fmt.Println("Perimeter:", perimeter1)
}
