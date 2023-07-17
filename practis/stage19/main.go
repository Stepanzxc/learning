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

type User struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MaidenName string  `json:"maidenName"`
	Age        int     `json:"age"`
	Gender     string  `json:"gender"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	BirthDate  string  `json:"birthDate"`
	Image      string  `json:"image"`
	BloodGroup string  `json:"bloodGroup"`
	Height     int     `json:"height"`
	Weight     float64 `json:"weight"`
	EyeColor   string  `json:"eyeColor"`
	Hair       struct {
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"hair"`
	Domain  string `json:"domain"`
	IP      string `json:"ip"`
	Address struct {
		Address     string `json:"address"`
		City        string `json:"city"`
		Coordinates struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
		PostalCode string `json:"postalCode"`
		State      string `json:"state"`
	} `json:"address"`
	MacAddress string `json:"macAddress"`
	University string `json:"university"`
	Bank       struct {
		CardExpire string `json:"cardExpire"`
		CardNumber string `json:"cardNumber"`
		CardType   string `json:"cardType"`
		Currency   string `json:"currency"`
		Iban       string `json:"iban"`
	} `json:"bank"`
	Company struct {
		Address struct {
			Address     string `json:"address"`
			City        string `json:"city"`
			Coordinates struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"coordinates"`
			PostalCode string `json:"postalCode"`
			State      string `json:"state"`
		} `json:"address"`
		Department string `json:"department"`
		Name       string `json:"name"`
		Title      string `json:"title"`
	} `json:"company"`
	Ein       string `json:"ein"`
	Ssn       string `json:"ssn"`
	UserAgent string `json:"userAgent"`
}

func getInfUrl(str string) (User, error) {

	response, err := http.Get(str)
	if err != nil {
		return User{}, err //Если ввозникла ошибка, то возвращаем пустубю структуру и ошибку
	}
	//Если status code от серера не 200 то тоже возращаем пусутюб стктуру и кастомную ошибку
	if response.StatusCode != 200 {
		//Создаеим кастомную ошибку
		return User{}, fmt.Errorf("status code from service %d", response.StatusCode)
	}
	var result User
	//Делаем JSON Decode
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		//Если не валидный JSON то воврввщаем ошибку
		return User{}, err
	}
	//В случае успешного результата возвращаем оезультат и пустую ошибку
	return result, nil
}
func main() {
	info, err := getInfUrl("https://dummyjson.com/users/1")
	if err != nil {
		//Только в этом случае нужно выходить из прогрраммы
		log.Fatalln(err)
	}
	err = WriteInFileCsv(info)
	if err != nil {
		log.Fatalln(err)
	}
}
func WriteInFileCsv(user User) error {
	f, err := os.Create("someFromUser.csv")
	if err != nil {
		log.Println(err)
	}

	result := [][]string{
		{"firstName", "age", "email", "phone", "lat", "lng"},
	}
	v := []string{(user.FirstName), strconv.Itoa(user.Age), user.Email, user.Phone, fmt.Sprint(user.Address.Coordinates.Lat), fmt.Sprint(user.Address.Coordinates.Lng)}
	result = append(result, v)

	w := csv.NewWriter(f)
	err = w.WriteAll(result)
	defer f.Close()
	return err
}
