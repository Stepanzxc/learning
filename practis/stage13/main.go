package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type LogItem struct {
	date         string
	ip           string
	statusCode   string
	url          string
	responseTime string
}

// getUniqUrls написать функциию которая вернет уникальные url's из лога,
// Нужно получить чистый URL без "GET" && "HTTP/2.0"
func getUniqUrls(arr []LogItem) []string {
	mapIp := map[string]struct{}{}
	for i := range arr {
		mapIp[arr[i].url] = struct{}{}
	}
	ips := []string{}
	row := []string{}
	for k := range mapIp {
		ips = append(ips, k)

	}
	for i := 0; i < len(ips); i++ {
		v := strings.Split(ips[i], " ")
		if len(v) != 3 {
			fmt.Println("Defect url")
			continue
		}
		row = append(row, v[1])
	}
	return row
}

func main() {
	items, err := getLogItems("logs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(getUniqUrls(items))

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
