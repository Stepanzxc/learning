package main

import "fmt"

type Square struct {
	Side int
}

func (s Square) Perimetr() {
	fmt.Println(s.Side * 4)
}
func (s *Square) Scale(mult int) {
	fmt.Println(s.Side * mult)
}
func main() {
	square := Square{Side: 4}
	pSquare := &square
	square.Perimetr()
	pSquare.Scale(2)
}
