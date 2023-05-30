package main

import "log"

func ReverseSeq(n int) []int {
	return make([]int, n)
}

func main() {
	log.Println(ReverseSeq(5)) //[]int{5, 4, 3, 2, 1}

}
