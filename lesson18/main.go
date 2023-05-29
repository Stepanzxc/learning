package main

import "fmt"

func main() {
	var defaultMap map[int64]string
	fmt.Println(defaultMap)
	mapLiter := map[string]int{"Vasya": 13, "Dima": 32}
	fmt.Println(mapLiter)
}
