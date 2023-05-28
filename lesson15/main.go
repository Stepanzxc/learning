package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func change(arr [3]int) [3]int {
	arr[1] = 666
	return arr
}
func main() {
	var intArr [3]int
	fmt.Println(intArr)
	person := [2]Person{{
		Name: "asdda",
		Age:  23,
	},
		{
			Name: "qweqwe",
			Age:  12,
		},
	}
	newIntArr := change(intArr)
	fmt.Println(newIntArr)
	fmt.Println(person)
	stringArr := [...]string{"1asd", "2df"}
	fmt.Println(stringArr)
	for i := 0; i < len(stringArr); i++ {
		fmt.Println(stringArr[i])
	}
	for _, value := range stringArr {
		fmt.Println(value)
	}
	slices()
}
func slices() {
	slice := []string{"First", "Second"}
	fmt.Println(slice)
	sliceByMake := make([]int, 0, 5)
	fmt.Println(sliceByMake)
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5, 3, 32, 421)
	fmt.Println(sliceByMake)
	for i, value := range sliceByMake {
		fmt.Println(i, value)
	}
}
