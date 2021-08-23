package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	response, err := http.Get("https://www.mercadobitcoin.net/api/BTC/ticker/")

	if err!=nil{
		fmt.Println("1 - error: ",err.Error())
		return
	}
	if response.StatusCode != 200{
		fmt.Println("2 - erro de status code |",response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)	

	type bitcoin struct{
		High string `json: "high"`
		Low string `json: "low"`
		Vol string `json: "vol"`
		Last string `json: "last"`
		Buy string `json: "buy"`
		Sell string `json: "sell"`
		Open string `json: "open"`
		Date int `json: "date"`
	}

	type ticker struct{
		Ticker bitcoin `json: "ticker"`
	}	

	var resp ticker

	err = json.Unmarshal(body,&resp)

	if err != nil{
		fmt.Println("3 - error:",err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>> ")
	UserInput, _ := reader.ReadString('\n')
	UserInput = strings.TrimSuffix(UserInput, "\n")

	switch UserInput{
	case "High": fmt.Println(resp.Ticker.High)
	case "Low": fmt.Println(resp.Ticker.Low)
	case "Vol": fmt.Println(resp.Ticker.Vol)
	case "Last": fmt.Println(resp.Ticker.Last)
	case "Buy": fmt.Println(resp.Ticker.Buy)
	case "Sell": fmt.Println(resp.Ticker.Sell)
	case "Open": fmt.Println(resp.Ticker.Open)
	case "Date": fmt.Println(resp.Ticker.Date)
	default: fmt.Println("Not found...")
	}


}