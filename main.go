package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// barcode to get data from api
	barcode := 74001755
	barCode := strconv.Itoa(int(barcode))
	url := "https://jsonmock.hackerrank.com/api/inventory?barcode=" + barCode

	client := &http.Client{}

	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// struct for bind data
	type Inventory struct {
		Barcode   string `json:"barcode"`
		Item      string `json:"item"`
		Category  string `json:"category"`
		Price     int32  `json:"price"`
		Discount  int32  `json:"discount"`
		Available int32  `json:"available"`
	}
	type Response struct {
		Page       int32       `json:"page"`
		PerPage    int32       `json:"per_page"`
		Total      int32       `json:"total"`
		TotalPages int32       `json:"total_pages"`
		Data       []Inventory `json:"data"`
	}

	var Body Response

	// read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// unmarshal api response to struct
	err = json.Unmarshal(body, &Body)
	if err != nil {
		fmt.Println(err)
	}

	// find discount price
	if len(Body.Data) != 0 {
		invent := Body.Data[0]
		discountedPrice := float64(invent.Price) * (float64(invent.Discount) / 100)
		fmt.Println("Item  :", invent.Item)
		fmt.Println("Price  :", invent.Price)
		fmt.Println("Discount %  :", invent.Discount)
		fmt.Println("Discount  :", discountedPrice)
		fmt.Println("Final price :", float64(invent.Price)-discountedPrice)
	} else {
		fmt.Printf("\nNo product on this ID %v\n", barCode)
	}

}
