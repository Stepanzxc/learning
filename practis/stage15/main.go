// //Создать функциб которая сделает GET запрос на стороний API и поплучить резульат в формате JSON
// // https://ipapi.co/176.115.145.30/json/
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	getJson()
}
func getJson() {
	url, err := http.Get("https://ipapi.co/176.115.145.30/json/")
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(url.Body).Decode(&result)
	log.Println(result)
}
