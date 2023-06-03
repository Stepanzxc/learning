package main

import (
	"log"
	"strings"
)

func StringToArray(str string) []string {
	v := strings.Split(str, " ")
	return v

}
func main() {
	log.Println(StringToArray("Robin Singh"))                        //[]string{"Robin", "Singh"}
	log.Println(StringToArray("StarWars"))                           //[]string{"StarWars"}
	log.Println(StringToArray("I love arrays they are my favorite")) //[]string{"I", "love", "arrays", "they", "are", "my", "favorite"}
	log.Println(StringToArray("1 2 3"))                              //[]string{"1", "2", "3"}

}
