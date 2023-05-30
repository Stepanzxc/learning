package main

import (
	"log"
)

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func main() {
	log.Println(MakeNegative(42) == -42)
	log.Println(MakeNegative(-5) == -5)
	log.Println(MakeNegative(0) == 0)
}
