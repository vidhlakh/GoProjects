package main
import (
	"fmt"
)
func main(){
book:="This is my book"
fmt.Println("Book:",book)
fmt.Println("Lenght of the Book:",len(book))
fmt.Println("Book 0:",book[1])
fmt.Println("Slcies:",book[1:6])
//String are unicode
unic:="String are unicode so u can many symbols1/2 :"
fmt.Println(unic)
//Multiline
multi:=`This is 
Multiline
String`
fmt.Println(multi)
}
