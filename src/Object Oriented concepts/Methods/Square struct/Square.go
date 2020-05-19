package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

// Move moves the point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

// Will create a new Square and validate the input
func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("Length cant be zero but given %v", length)
	}
	square := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return square, nil
}
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}
func (s *Square) Area() int {
	return s.Length * s.Length
}
func main() {
	s, err := NewSquare(2, 4, -10)
	if err == nil {
		s.Move(3, 5)
		fmt.Println("X and Y moved to ", s.Center)

		area := s.Area()
		fmt.Printf
		fmt.Printf("Area for length %v is %v", s.Length, area)
		os.Exit(1)
	} else {
		log.Fatalf("Error cant create Square ")
		fmt.Println("Error: ", err)
	}
}
