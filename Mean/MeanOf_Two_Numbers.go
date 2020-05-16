package main
import (
    "fmt"
)
func main(){
var x int
var y int
x=10
y=20
var mean int
mean = (x+y)/2

fmt.Printf("x=%v type of %T\n",x,x)
fmt.Printf("y=%v type of %T\n",y,y)
fmt.Printf("mean=%v type of %T",mean,mean)
}