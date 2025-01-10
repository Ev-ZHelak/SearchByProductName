package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ResponseData map[string]int

func main() {
	url := "https://raw.githubusercontent.com/Ev-ZHelak/SearchByProductName/refs/heads/main/db.json"

	data, err := downloadFile(url)
	if err != nil {
		panic("Ошибка загрузки файла: " + err.Error())
	} else {
		fmt.Println(`Файл "db.json" успешно загружен`)
	}

	sc := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Введите название товара: ")
		sc.Scan()
		input := sc.Text()

		if input == "0" {break}

		result, err := searchProduct(data, input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(strings.Join(result, "\n"))
		}
	}
	
}

func downloadFile(url string) (ResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return ResponseData{}, fmt.Errorf("ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ResponseData{}, fmt.Errorf("ошибка HTTP: статус код %d", resp.StatusCode)
	}

	var data ResponseData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return ResponseData{}, fmt.Errorf("ошибка декодирования JSON: %v", err)
	}

	return data, nil
}

func searchProduct(data ResponseData, x string) ([]string, error) {
	var listProduct []string
	for nameProduct, price := range data {
		if strings.Contains(strings.ToLower(nameProduct), strings.ToLower(x)) {
			listProduct = append(listProduct, fmt.Sprintf("%v: %v", nameProduct, price))
		}
	}

	if listProduct != nil {
		return listProduct, nil
	}

	return []string{}, fmt.Errorf("товар \"%v\" не найден", x)
}
