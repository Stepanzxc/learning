package main

import "fmt"

func main() {
	first, second := 23, 12
	var mult func(x, y int) int
	mult = func(x, y int) int { return x * y }
	fmt.Println(mult(first, second))

	div := func(x, y int) int { return x / y }
	fmt.Println(div(first, second))
	SumF := func(x, y int) int {
		return x + y
	}
	fmt.Println(calculate(first, second, SumF))
	divied2 := CreateDivider(2)
	divied10 := CreateDivider(10)
	fmt.Println(divied2(100))
	fmt.Println(divied10(100))
	a := 12
	geta := func() int {
		return a
	}
	fmt.Println(geta())
	a = 1234
	fmt.Println(geta())
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}
func CreateDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
