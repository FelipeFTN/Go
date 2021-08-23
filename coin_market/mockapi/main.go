package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main(){
	resp, err := http.Get("https://6123b9b0124d880017568442.mockapi.io/coin/data")

	if err != nil{
		fmt.Print("1 - Error: ", err.Error())
		return
	}

	if resp.StatusCode != 200{
		fmt.Print("Status code:",resp.StatusCode)
		return
	}

	type Database struct{
		ID int `json:"id"`
		Name string `json:,"name"`
		Symbol string `json:,"symbol"`
	}

	body, err := io.ReadAll(resp.Body)
	var response []Database

	err = json.Unmarshal(body, &response)

	if err != nil{
		fmt.Println("Error: ",err.Error())
		return
	}

	//fmt.Println(response)
	//fmt.Fprintf(resp, response)
}