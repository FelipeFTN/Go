package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Data struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Supply float64 `json:"circulating_supply"`
}
type GetData struct {
	Data []Data `json:"data"`
}

func index(w http.ResponseWriter, r *http.Request) {
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
	respBody, _ := ioutil.ReadAll(resp.Body)
	responseData := GetData{}

	json.Unmarshal(respBody, &responseData)

	for i := 0; i < len(responseData.Data); i++ {

		json.NewEncoder(w).Encode([]Data{
			{
				ID:     responseData.Data[i].ID,
				Name:   responseData.Data[i].Name,
				Supply: responseData.Data[i].Supply,
			},
		})
	}
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
