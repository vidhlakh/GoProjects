package main
import(
	"fmt"
	"strings"
)
func main(){
	stocks:= map[string]float64{
		"Amazon": 120.22,
		"Microsoft": 2424.24,
	}
	for key, value:= range stocks{
		fmt.Println(key,":",value)
	}
	fmt.Println(len(stocks))
	//use 2 value to find
	val,ok:=stocks["Microsoft"]
	if !ok{
		fmt.Println("Key not found")
	}else{
		fmt.Println(val)
	}
	//set value
	stocks["Ford"] = 130.44

	//Delete Map 
	delete (stocks,"Amazon")
	fmt.Println(stocks)

	// Challenge on Maps -> Word count
	sent:="This is a string where word count many to be many . How many words are present more present once in this sentence"
	words:=strings.Fields(sent)
	wordMap:=map[string]int{

	}
	for _,word:= range words{
		wordInLow:=strings.ToLower(word)
		_,ok:=wordMap[wordInLow]
		if ok{
			wordMap[wordInLow]+=1
		}else{
			wordMap[wordInLow]=1
		}
		
	}
	fmt.Println(wordMap)

	//Another method
	counts:=map[string]int{}
	for _,word :=range(words){
		counts[strings.ToLower(word)]++

	}
	fmt.Println("Counts:",counts)	
}