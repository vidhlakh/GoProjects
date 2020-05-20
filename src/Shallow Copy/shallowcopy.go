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
	p2.name = "Not Alena"

	fmt.Printf("%+v\n", p)  // {name:Alena}
	fmt.Printf("%+v\n", p2) // {name: Not Alena}
	// change the p name. Changes will be affected on both

	p.name = "Not Alena"
	fmt.Printf("%+v\n", p)  // {name:A Not lena}
	fmt.Printf("%+v\n", p2) // {name: Not Alena}
}
