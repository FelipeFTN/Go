package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "99731e72-b235-475e-aedd-ae6699e0a859")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	type data struct {
		ID     int    `json:"id"`
		Name   string `json:,"name"`
		Symbol string `json:,"symbol"`
	}

	type DataWanted struct {
		Data []data `json:"data"`
	}

	respBody, _ := io.ReadAll(resp.Body)

	var response DataWanted

	err = json.Unmarshal(respBody, &response)

	for index, _ := range response.Data {

		var nome = response.Data[index].Name

		switch response.Data[index].Name {
		case "Bitcoin":
			fmt.Println("Moeda", index, ":", nome)
		case "Ethereum":
			fmt.Println("Moeda", index, ":", nome)
		case "Synthetix":
			fmt.Println("Moeda", index, ":", nome)
		case "Litecoin":
			fmt.Println("Moeda", index, ":", nome)
		}
	}

	//fmt.Println(response)

}
