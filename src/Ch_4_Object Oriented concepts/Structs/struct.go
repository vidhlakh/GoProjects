package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value -= value
	}
	return value
}
func main() {
	t1 := Trade{"Ama", 1, 12.2, false}
	t2 := Trade{
		Symbol: "Micro",
		Volume: 2,
		Price:  23.33,
		Buy:    true,
	}
	t3 := Trade{}
	fmt.Printf("Trade t1: %+v", t1)
	fmt.Println("Trade t2:", t2)
	fmt.Println("Trade t3:", t3)
	// Value method having Trad struct as reference
	val := t2.Value()
	fmt.Println("Value after calling Value method", val)
}
