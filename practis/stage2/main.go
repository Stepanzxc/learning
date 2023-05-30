package main

import "log"

func PositiveSum(numbers []int) int {
	x := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			x += numbers[i]
		}

	}
	return x
}

func main() {
	log.Println(PositiveSum([]int{1, 2, 3, 4, 5}) == 15)
	log.Println(PositiveSum([]int{1, -2, 3, 4, 5}) == 13)
	log.Println(PositiveSum([]int{}) == 0)
	log.Println(PositiveSum([]int{-1, -2, -3, -4, -5}) == 0)
	log.Println(PositiveSum([]int{-1, 2, 3, 4, -5}) == 9)
}
