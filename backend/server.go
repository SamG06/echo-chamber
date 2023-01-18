package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Test struct {
	Post string `json:"post"`
	Minutes int `json:"minutes"`
}


func writeHeader(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		writeHeader(&w)

		x := Test{Post: "Hello World", Minutes: 9}
		
		json.NewEncoder(w).Encode(x)
	})

	fmt.Print("Starting server at port 3000\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}