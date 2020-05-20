//using the custom package
package main

import (
	"fmt"
	"log"
	"rectangle" // put the custom package in the GOPATH eg. /home/****/go
)

/*
 * 1. package variables
 */
var rectLen, rectWidth float64 = 6, 7

func init() {
	println("main package initialized now")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}
func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	 */
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used
	 */
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
