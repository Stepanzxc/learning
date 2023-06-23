package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type StructForProd struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
}

type StructUrl struct {
	Products []StructForProd `json:"products"`
}

func main() {
	info, err := getInfUrl("https://dummyjson.com/products")
	if err != nil {
		//Только в этом случае нужно выходить из прогрраммы
		log.Fatalln(err)
	}
	errs := WriteInFileCsv(info)
	if errs != nil {
		log.Fatalln(errs)
	}
}
func WriteInFileCsv(prod StructUrl) error {
	f, err := os.Create("products.csv")
	if err != nil {

		return err
	}

	result := [][]string{
		{"Id", "Title", "Description", "Price", "Brand", "Category"},
	}
	for i := range prod.Products {
		v := []string{strconv.Itoa(prod.Products[i].Id), prod.Products[i].Title, prod.Products[i].Description, strconv.Itoa(prod.Products[i].Price), prod.Products[i].Brand, prod.Products[i].Category}
		result = append(result, v)

	}
	w := csv.NewWriter(f)
	err = w.WriteAll(result)
	defer f.Close()
	return err
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
