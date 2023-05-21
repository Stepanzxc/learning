package main

import "fmt"

func main() {
	var str string = "Bld bdld"
	var a int64 = 5
	b := 6
	fmt.Println(str)
	fmt.Printf("Type: %T Value: %v\n", str, str)
	fmt.Println(a + int64(b))
}
