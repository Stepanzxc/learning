package main

import "log"

func Between(a, b int) []int {
	c := []int{a}
	for i := a; i < b; i++ {
		c = append(c, (a + 1))
		a++
	}
	return c
}

func main() {
	log.Println(Between(1, 4))  //[]int{1, 2, 3, 4}
	log.Println(Between(-2, 2)) //[]int{-2, -1, 0, 1, 2}

}
