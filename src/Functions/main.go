package main
import(
	"fmt"
	"functions"
)

func main(){
	a:=100
	b:=20
	c:=functions.add(a,b)
	fmt.Println("c=",c)
	div,rem := functions.divmod(a,b)
	fmt.Println("Result=",div,"Reminder=",rem)
}