// //Создать функциб которая сделает GET запрос на стороний API и поплучить резульат в формате JSON
// // https://ipapi.co/176.115.145.30/json/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

type IPInfo struct {
	Ip                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

func main() {
	//парсим IP адрес из строки и получаем тип net.IP
	ip := net.ParseIP("8.8.8.8")
	//Передаем IP адресс в функцию
	info, err := InfoByIP(ip)
	if err != nil {
		//Только в этом случае нужно выходить из прогрраммы
		log.Fatalln(err)
	}
	log.Printf("%+v", info)

	//fmt.Println(getJson("https://ipapi.co/176.115.145.30/json/"))
}

// InfoByIP получение информации по IP адресу клиента
func InfoByIP(ip net.IP) (IPInfo, error) {
	//Инизиализируем пусую структуру
	var result IPInfo
	//Осуществляем запрос к удаленному серверу
	response, err := http.Get(fmt.Sprintf("https://ipapi.co/%s/json/", ip.String()))
	if err != nil {
		return IPInfo{}, err //Если ввозникла оошибка, то возвращаем пустубю структуру и ошибку
	}
	//Если status code от серера не 200 то тоже возращаем пусутюб стктуру и кастомную ошибку
	if response.StatusCode != 200 {
		//Создаеим кастомную ошибку
		return IPInfo{}, fmt.Errorf("status code from service %d", response.StatusCode)
	}
	//Делаем JSON Decode
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		//Если не валидный JSON то воврввщаем ошибку
		return IPInfo{}, err
	}
	//В случае успешного результата возвращаем оезультат и пустую ошибку
	return result, nil
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
