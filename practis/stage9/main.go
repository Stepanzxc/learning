package main

import (
	"log"
)

func Solution(a, b string) string {
	c := ""
	if len(a) < len(b) {
		c = a + b + a
	} else {
		c = b + a + b
	}
	return c
}

func main() {
	log.Println(Solution("45", "1") == "1451")
	log.Println(Solution("Soon", "Me") == "MeSoonMe")
	log.Println(Solution("U", "False") == "UFalseU")
	log.Println(Solution("13", "200") == "1320013")

}
