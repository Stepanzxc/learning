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

// Функция должна принимать API адресс и ввозвращать result и err
func getJson(str string) map[string]interface{} {
	url, err := http.Get(str)
	if err != nil {
		//Функция не должна заканчиваться, если ты объявил что она что то возввращает то возвращай этот тип данных
		log.Fatalln(err)
	}
	//Переведи на struct а не на map[string]interface{}  вот ссылка с ответами https://stackoverflow.com/questions/47270595/how-to-parse-json-string-to-struct
	var result map[string]interface{}
	//Не обработал ошибку
	json.NewDecoder(url.Body).Decode(&result)
	//Зачем проерять на статус code decode логичнее это сделать до
	if url.StatusCode < 200 || url.StatusCode >= 300 {
		fmt.Println("status code: ", url.StatusCode)
		fmt.Println("Reason error: ", http.StatusText(url.StatusCode))
		result = nil
	}

	return result
}
