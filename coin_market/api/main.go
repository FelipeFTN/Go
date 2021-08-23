package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/coinmarket", datareader)

	fmt.Println("Server ON")

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func datareader(res http.ResponseWriter, req *http.Request){

	type data struct {
		ID     int    `json:"id"`
		Name   string `json:,"name"`
		Symbol string `json:,"symbol"`
	}

	type DataWanted struct {
		Data []data `json:"data"`
	}

	res.Header().Set("content-type","application/json")

	/*json.NewEncoder(res).Encode(DataWanted{{

	}})*/

	

}