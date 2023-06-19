// //Создать функциб которая сделает GET запрос на стороний API и поплучить резульат в формате JSON
// // https://ipapi.co/176.115.145.30/json/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(getJson("https://ipapi.co/176.115.145.30/json/"))
}
func getJson(str string) map[string]interface{} {
	url, err := http.Get(str)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(url.Body).Decode(&result)
	if url.StatusCode < 200 || url.StatusCode >= 300 {
		fmt.Println("status code: ", url.StatusCode)
		fmt.Println("Reason error: ", http.StatusText(url.StatusCode))
		result = nil
	}

	return result
}
