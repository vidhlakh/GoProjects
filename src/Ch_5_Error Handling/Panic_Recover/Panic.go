package main

import (
	"fmt"
)

//Recover the panic
func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Eeerror:%s\n", err)
		}
	}()
	return vals[index]
}
func main() {
	//Use of panic is discouragered in go
	//Code where panic is created automatically
	// val := []int{1, 2, 3}
	// fmt.Printf("Valaue 10%v", val[10])

	//Creating panic
	// file, err := os.Open("Nofile.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println("v value:", v)
}
