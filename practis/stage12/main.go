package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

func GetUnicIp(arr []LogItem) []string {
	mapIp := map[string]struct{}{}
	for i := range arr {
		mapIp[arr[i].ip] = struct{}{}
	}
	ips := []string{}
	for k := range mapIp {
		if net.ParseIP(k) != nil {
			ips = append(ips, k)
		}

	}
	return ips
}

func main() {
	items, err := getLogItems("logs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(GetUnicIp(items))

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
