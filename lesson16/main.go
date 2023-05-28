package main

import "fmt"

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println(val)
	}
}
func main() {
	arra := []int{321, 54, 43, 1, 2, 3}
	showAllElements(123, 123, 332, 4224, 21)
	showAllElements(arra...)
	intArray := (*[6]int)(arra)
	fmt.Println(intArray)
}
