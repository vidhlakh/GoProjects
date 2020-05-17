package main
import(
	"fmt"
	
)

func main(){
	a:=100
	b:=20
	c:=add(a,b)
	fmt.Println("c=",c)
	div,rem := divmod(a,b)
	fmt.Println("Result=",div,"Reminder=",rem)
	//Using DoubleValue Function
	values:=[]int {1,2,3,4}
	doubleValue(values,2)
	fmt.Println(values)
}