package main

import "fmt"

type OurString string
type Person struct {
	Name string
	Age  int
}

func main() {
	var a OurString
	fmt.Println(a)
	var Dan Person
	fmt.Println(Dan)
	Dan.Name = "Dan"
	Dan.Age = 32
	fmt.Println(Dan)
	Bread := Person{
		Name: "Bread",
		Age:  31,
	}
	fmt.Println(Bread)
	Vlad := Person{"Vlad", 321}
	fmt.Println(Vlad)
	pIvan := &Person{"Ivan", 23}
	fmt.Println(pIvan)
	non := struct {
		Name, name, Lastn string
	}{
		Name:  "asd",
		name:  "dasd",
		Lastn: "q	w",
	}
	fmt.Println(non)

}
