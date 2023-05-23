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
}
