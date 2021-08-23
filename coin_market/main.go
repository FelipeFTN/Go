package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cat struct {
	Name   string `json:"name"`
	Rank   string `json:"rank"`
	Supply string `json:"circulating_supply"`
}

func index(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://6123b9b0124d880017568442.mockapi.io/coin/data")
	if err != nil {
		print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	responseData := Cat{}
	json.Unmarshal(body, &responseData)

	fmt.Println(responseData)

	//json.NewEncoder(w).Encode([]Cat{
	//	{
	//		Rank:   responseData.Rank,
	//		Name:   responseData.Name,
	//		Supply: responseData.Supply,
	//	},
	//})
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
