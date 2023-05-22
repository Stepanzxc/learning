package main

import "fmt"

func main() {
	age := 23
	if age < 30 {
		fmt.Println("You are too young")
	}
	age = 8
	if children := Age(age); !children == false {
		fmt.Println("Young")
	}
	if age <= 30 && age < 45 {
		fmt.Println("=30")
	} else {
		fmt.Println("Ok")
	}

}
func Age(age int) bool {
	return age < 10
}
