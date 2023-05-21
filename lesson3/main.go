package main

import "fmt"

func main() {
	Greet()
	a, b := 23, 1
	PersonGreet("Stepan")
	FIO("Maks", "QWE")
	summa := Sum(a, b)
	sum, mult := Named(a, b)
	fmt.Println(sum, mult)
	fmt.Println(summa)
}
func Greet() {
	fmt.Println("Hello")
}
func PersonGreet(name string) {
	fmt.Printf("Privet %s\n", name)
}
func FIO(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}
func Sum(a, b int) int {
	sum := a + b
	return sum
}
func SumAndMulti(a, b int) (int, int) {
	return a + b, a * b
}
func Named(a, b int) (sum int32, mult float32) {
	sum = int32(a + b)
	mult = float32(a * b)
	return
}
