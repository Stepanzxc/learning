package main

import "log"

func StringToArray(str string) []string {
	b := []string{str}
	for i := 0; i <= len(str); i++ {
		if b[i] == " " {
			log.Println("adsasds")
		}
	}
	return b
}
func main() {
	log.Println(StringToArray("Robin Singh"))                        //[]string{"Robin", "Singh"}
	log.Println(StringToArray("StarWars"))                           //[]string{"StarWars"}
	log.Println(StringToArray("I love arrays they are my favorite")) //[]string{"I", "love", "arrays", "they", "are", "my", "favorite"}
	log.Println(StringToArray("1 2 3"))                              //[]string{"1", "2", "3"}

}
