package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StructForProd struct {
	Id                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}
type StructUrl struct {
	Products []StructForProd `json:"products"`
	Total    int             `json:"total"`
	Skip     int             `json:"skip"`
	Limit    int             `json:"limit"`
}

func main() {
	info, err := getInfUrl("https://dummyjson.com/products")
	if err != nil {
		//Только в этом случае нужно выходить из прогрраммы
		log.Fatalln(err)
	}
	log.Printf("%+v", info)

}

func getInfUrl(str string) (StructUrl, error) {
	var result StructUrl
	response, err := http.Get(str)
	if err != nil {
		return StructUrl{}, err //Если ввозникла ошибка, то возвращаем пустубю структуру и ошибку
	}
	//Если status code от серера не 200 то тоже возращаем пусутюб стктуру и кастомную ошибку
	if response.StatusCode != 200 {
		//Создаеим кастомную ошибку
		return StructUrl{}, fmt.Errorf("status code from service %d", response.StatusCode)
	}
	//Делаем JSON Decode
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		//Если не валидный JSON то воврввщаем ошибку
		return StructUrl{}, err
	}
	//В случае успешного результата возвращаем оезультат и пустую ошибку
	return result, nil
}
