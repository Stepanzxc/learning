package main

import "fmt"

type Bulder interface {
	Build()
}
type Person struct {
	Name string
	Age  int
}
type WoodBuilder struct {
	Person
}

func (wd WoodBuilder) Build() {
	fmt.Println("WOOD")
}

type StoneBuilder struct {
	Person
}

func (sb StoneBuilder) Build() {
	fmt.Println("Stone")
}

type Building struct {
	Bulder
	Name string
}

func usecase() {
	WoodenBuilder := Building{
		Bulder: WoodBuilder{
			Person{
				Name: "VAsz",
				Age:  12,
			},
		},
		Name: "Wood",
	}
	fmt.Println(WoodenBuilder)
	StonenBuilder := Building{
		Bulder: StoneBuilder{
			Person{
				Name: "Ara",
				Age:  123,
			},
		},
		Name: "Stone",
	}
	fmt.Println(StonenBuilder)
}
func main() {
	builder := WoodBuilder{Person{"sdada", 123}}
	fmt.Println(builder)
	usecase()
}
