package main

import (
	"fmt"
	"log"
)

type Runner interface {
	Run() string
}
type Book struct {
	Title  string
	Author string
}

func Write(s fmt.Stringer) {
	log.Println(s.String())
}
func (b Book) String() string {
	return fmt.Sprintf("Book: %s-%s", b.Title, b.Author)
}
func main() {
	var runner Runner
	fmt.Println(runner)

	book := Book{"dfff", "qwer"}
	Write(book)
}
