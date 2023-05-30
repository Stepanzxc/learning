package main

import "log"

func RepeatStr(repetitions int, value string) string {
	str := ""
	for i := 0; i < repetitions; i++ {
		str += value
	}

	return str
}

func main() {
	log.Println(RepeatStr(4, "a") == "aaaa")
	log.Println(RepeatStr(2, "test ") == "test test ")
}
