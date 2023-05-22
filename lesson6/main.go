package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	sum := 12
	for sum <= 100 { //while
		sum++
		fmt.Println(sum)
	}
}
