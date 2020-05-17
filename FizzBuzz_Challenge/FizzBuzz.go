package main
import (
	"fmt"
)
func main(){
	number:=15
	if number%3==0 && number%5==0{
		fmt.Println("Fizz Buzz")
	}else if number%3==0{
		fmt.Println("fizz")
	}else if number%5==0{
		fmt.Println("Buzz")
	}else{
		fmt.Println("i")
	}
	
}