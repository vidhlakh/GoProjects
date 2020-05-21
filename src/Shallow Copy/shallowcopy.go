package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	p := Person{name: "Alena"}

	// shallow copy
	p2 := p
	fmt.Printf("%+v\n", p)  // {name:Alena}
	fmt.Printf("%+v\n", p2) // {name:Alena}

	// change the copied name
	p2.name = "Not Alena changed in 2"

	fmt.Printf("p %+v\n", p)   // {name:Alena}
	fmt.Printf("p2 %+v\n", p2) // {name: Not Alena}
	// change the p name. Changes will be affected on both

	p.name = "Not Alena changed in p"
	fmt.Printf("p %+v\n", p)   // {name:A Not lena}
	fmt.Printf("p2 %+v\n", p2) // {name: Not Alena}

	//Array in assignment also memory address is different as arrays are value types 
	a := [...]int{1, 2, 3, 4, 5}
	b := a
	fmt.Println("a:", a, "mem ", &a[0])
	fmt.Println("b: ", b, "mem", &b[0])
}
