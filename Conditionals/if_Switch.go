package main
import (
    "fmt"
)
func main(){
   
    x:=10.0
    y:=20.0
    if (x>5){
		fmt.Printf("X is greater than 5\n")
	}else {
		fmt.Printf("X is greater than 5\n")
	}
	if frac:=x/y; frac>1{
		fmt.Printf("Y is greater than X\n")

	}else{
		fmt.Printf(" X is greater than Y\n")
	}
	switch(x){
	case 1:
		fmt.Printf("One")
	case 10:
		fmt.Printf("five")
	default:
		fmt.Printf("default switch")
	}
	for i:=0;i<5;i++{
		fmt.Printf("i=%v\n",i)
	}
	//Like While loop
	j:=0
	for j<4{
		fmt.Println("J=",j)
		j++
	}
    }