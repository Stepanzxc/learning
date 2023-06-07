package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LogItem struct {
	date         string
	ip           string
	statusCode   string
	url          string
	responseTime string
}

// getMaxTimeURL найти максимальное URL по времени выполнения
func getMaxTimeURL(arr []LogItem) string {
	mapTimeUrl := map[string]float64{}
	addArr := []string{}
	for i := range arr {
		addArr = append(addArr, arr[i].responseTime)
	}
	floatArr := getFloatResponseTime(addArr)
	for i := 0; i < len(arr); i++ {

		mapTimeUrl[arr[i].url] = floatArr[i]
	}
	url := ""
	var maxFloat float64
	for key, value := range mapTimeUrl {
		if value >= maxFloat {
			maxFloat = value
			url = key
		}
	}
	return url
}
func getFloatResponseTime(arr []string) []float64 {
	time := arr
	timefl := []string{}
	var q []float64
	for i := range time {
		if len(time[i]) == 0 {
			time[i] = time[i+1]

		}
		del := time[i]
		del = del[1:]
		del = del[:len(del)-1]
		timefl = append(timefl, del)
		floatTime, err := strconv.ParseFloat(timefl[i], 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		q = append(q, floatTime)

	}
	return q
}

// getAvgTimeURL почитать среднее время выполнениея всех запросов
func getAvgTimeURL(arr []LogItem) float64 {
	time := []string{}
	for i := range arr {
		time = append(time, arr[i].responseTime)
	}

	resTime := getFloatResponseTime(time)
	var avgTime float64
	for i := 0; i < len(resTime); i++ {
		avgTime = avgTime + resTime[i]
	}
	avgTime = avgTime / float64(len(resTime))
	return avgTime
}

func main() {
	items, err := getLogItems("logs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Avg = ", getAvgTimeURL(items))
	fmt.Println("Url = ", getMaxTimeURL(items))

}
func getLogItems(filiename string) ([]LogItem, error) {
	file, err := os.Open(filiename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	structArr := make([]LogItem, 0)
	for scanner.Scan() {

		row := strings.Split(scanner.Text(), "||")
		if len(row) != 5 {
			log.Printf("row has error : %s", row)
			continue
		}

		structArr = append(structArr, LogItem{
			date:         row[0],
			ip:           row[1],
			statusCode:   row[2],
			url:          row[3],
			responseTime: row[4],
		})

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return structArr, err
}
