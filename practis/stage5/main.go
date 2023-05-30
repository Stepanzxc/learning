package main

import (
	"log"
	"strconv"
)

func MultiTable(number int) string {
	var a string
	b := ""
	for i := 10; i >= 1; i-- {
		if i <= 9 {
			a = strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(number*i) + "\n"
			b = a + b
		} else {
			a = strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(number*i)
			b = a + b
		}
	}
	return b
}

func main() {
	log.Println(MultiTable(5) == "1 * 5 = 5\n2 * 5 = 10\n3 * 5 = 15\n4 * 5 = 20\n5 * 5 = 25\n6 * 5 = 30\n7 * 5 = 35\n8 * 5 = 40\n9 * 5 = 45\n10 * 5 = 50")
	log.Println(MultiTable(1) == "1 * 1 = 1\n2 * 1 = 2\n3 * 1 = 3\n4 * 1 = 4\n5 * 1 = 5\n6 * 1 = 6\n7 * 1 = 7\n8 * 1 = 8\n9 * 1 = 9\n10 * 1 = 10")
}
