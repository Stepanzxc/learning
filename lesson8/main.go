package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 2
	max = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(max-min) + min
	switch value {
	case 2:
		fmt.Println("2")

	case 3, 4:
		fmt.Println("3, 4")

	case Five():
		fmt.Println("5")
	default:
		fmt.Println("6")
	}
}
func Five() int {
	return 5
}
