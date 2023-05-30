package main

import "log"

func monkeyCount(n int) []int {
	a := []int{}
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	return a[:]
}
func main() {
	log.Println(monkeyCount(1))  //[]int{1}
	log.Println(monkeyCount(5))  //[]int{1,2,3,4,5}
	log.Println(monkeyCount(10)) //[]int{1,2,3,4,5,6,7,8,9,10}
}
