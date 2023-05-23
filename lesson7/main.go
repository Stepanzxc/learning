package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
fvm:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if i > 3 {
				continue fvm
			}
			if i == 2 {
				break
			}
		}
	}
}
