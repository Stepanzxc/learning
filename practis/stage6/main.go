package main

import "log"

func CountBy(x, n int) []int {
	a := []int{}
	for i := 1; i <= n; i++ {
		a = append(a, (i * x))
	}
	return a
}

func main() {
	log.Println(CountBy(1, 5))   //[]int{1, 2, 3, 4, 5}
	log.Println(CountBy(2, 5))   //[]int{2, 4, 6, 8, 10}
	log.Println(CountBy(3, 5))   //[]int{3, 6, 9, 12, 15}
	log.Println(CountBy(50, 5))  //[]int{50, 100, 150, 200, 250}
	log.Println(CountBy(100, 5)) //[]int{100, 200, 300, 400, 500}

}
