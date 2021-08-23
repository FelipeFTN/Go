package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Coins struct {
	ID    int
	Name  string
	Value string
}
type Ticker struct {
	High string `json:"high"`
	Vol  string `json:"vol"`
	Buy  string `json:"buy"`
}
type Mercadobitcoin struct {
	Ticker Ticker `json:"ticker"`
}

func coinMarket(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://www.mercadobitcoin.net/api/BTC/ticker/")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	responseData := Mercadobitcoin{}
	err := json.Unmarshal(body, &responseData)
	if err != nil {
		print(err)
	}
	json.NewEncoder(w).Encode([]Coins{
		{
			ID:    1,
			Name:  "Bitcoin",
			Value: responseData.Ticker.Buy,
		},
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./static/index.html")

	json.NewEncoder(w).Encode([]Coins{
		{
			ID:   1,
			Name: "Bitcoin",
		},
		{
			ID:   2,
			Name: "Dodgecoin",
		},
	})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/coin-market", coinMarket)
	http.ListenAndServe(":3000", nil)
}
