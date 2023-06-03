package main

import "log"

func ReverseSeq(n int) []int {
	a := []int{}
	for i := n; i > 0; i-- {
		a = append(a, i)
	}
	return a
}

func main() {
	log.Println(ReverseSeq(5)) //[]int{5, 4, 3, 2, 1}

}
