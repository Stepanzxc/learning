package main

import "fmt"

func main() {
	destineshion := make([]string, 0, 4)
	intArr := []string{"1", "2", "3", "4", "5"}
	fmt.Println(intArr)
	intSlice := intArr[1:3]
	fmt.Println(intSlice)
	fullSlice := intArr[:]
	fmt.Println(fullSlice)
	destineshion = make([]string, len(intArr))
	fmt.Println("Copy:", copy(destineshion, intArr))
	slice := []int{1, 2, 3, 4, 5}
	i := 2
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
	withCopy := slice[:i+copy(slice[i:], slice[i+1:])]
	fmt.Println(withCopy)
}
