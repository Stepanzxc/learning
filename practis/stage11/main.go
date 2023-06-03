package main

import (
	"log"
	"strconv"
)

// SumMix my favorite functions
func SumMix(arr []any) int {
	var result []int

	for i := range arr {
		anyToStr, ok := arr[i].(string)
		if ok {
			strToInt, err := strconv.Atoi(anyToStr)
			if err != nil {
				continue
			}
			result = append(result, strToInt)
		}
		anyToInt, ok := arr[i].(int)
		if ok {
			result = append(result, anyToInt)
		}
	}

	var count int
	for i := range result {
		count += result[i]
	}
	return count
}

func main() {
	log.Println(SumMix([]any{9, 3, "7", "3"}) == 22)
	log.Println(SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}) == 42)
	log.Println(SumMix([]any{"1", "5", "8", 8, 9, 9, 2, "3"}) == 45)
	log.Println(SumMix([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}) == 61)
	log.Println(SumMix([]any{1, 3}) == 61)

}
