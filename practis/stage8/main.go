package main

import "log"

func Between(a, b int) []int {
	// your code here
	return nil
}

func main() {
	log.Println(Between(1, 4))  //[]int{1, 2, 3, 4}
	log.Println(Between(-2, 2)) //[]int{-2, -1, 0, 1, 2}

}
