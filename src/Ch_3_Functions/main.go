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
	//Using DoubleValue Function . Go pass Slice, Map by reference . 
	// Without return the value , values are changed
	values:=[]int {1,2,3,4}
	doubleValue(values,2)
	fmt.Println("Value Changed as it is passed by Reference",values)
	//Integer pass by value 
	val:=20
	double(val)
	fmt.Println("Value not changes as it is passed by Value",val)
	doublePtr(&val)
	fmt.Println("Value passing the pointer ",val)

	//Return value and error 
	sq , err := sqroot(-2)
	if err==nil{
		fmt.Println("Sqroot value",sq)
	}else{
		fmt.Println("error:",err)
	}
	//Using Defer
	worker()

	//Challenge 
	contentHeader,error:= contentType("https://golang.org")
	if error==nil{
		fmt.Println("Contt",contentHeader)
	}else{
		fmt.Println("ErRorr: ",error)
	}
}