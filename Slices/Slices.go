package main
import(
	"fmt"
)
func main(){
	//Slices built on top of arrays 
	slices:= []string{"vidhya","lakshmi","Sankaranarayanan"}
	fmt.Printf("Slices=%v type of (%T)\n",slices,slices)
	fmt.Println(len(slices))
	a:=[]int{1,2,3,4,5,6}
	b:=a[3:]
	a[4]=20
	fmt.Println(a)
	fmt.Println(b)
	// ITerating through For loop
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
	for i:= range b{   // u can use _ instead of i if u not gonna use i
		fmt.Println("B value",i, b[i])
	}
	for i,name:= range slices{
		fmt.Println("B vlue ",i, name)
	}
	a = append(b,9)
	fmt.Println(a)

	// Challenge -> Find the maximum in the slice  a
	max:=a[0]
	for i:=0;i<len(a);i++{
		if a[i]> max {
			max=a[i]
		}
	
	}
	fmt.Println(max)
}