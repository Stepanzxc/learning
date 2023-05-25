package main

import "fmt"

func main() {
	var a int64 = 8
	var PointerA *int64 = &a
	fmt.Println(*PointerA)
	var newPointer = new(int32)
	fmt.Println(newPointer, *newPointer)
	*newPointer = 12
	fmt.Println(newPointer, *newPointer)
	num := 12
	sqrt(num)
	fmt.Println(num)
	sqrty(&num)
	fmt.Println(num)
	var wallet *int
	fmt.Println(haswallet(wallet))
}
func sqrt(num int) {
	num *= num
}
func sqrty(num *int) {
	*num *= *num
}
func haswallet(wallet *int) bool {
	return wallet != nil
}
