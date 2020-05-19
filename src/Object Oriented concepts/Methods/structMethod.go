package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}
func main() {
	p := Point{1, 2}
	p.move(2, 3)
	fmt.Printf("%+v\n", p)
}
