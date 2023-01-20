package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type Test struct {
	Post string `json:"post"`
	Minutes int `json:"minutes"`
}

func main() {
	http.HandleFunc("/", multiplexer)

	fmt.Print("Starting server at port 3000\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func setupForJson(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusCreated)
}

func multiplexer(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	}
}

var index = regexp.MustCompile(`^/$`)

func getHandler(w http.ResponseWriter, r *http.Request){
	switch {
	case index.MatchString(r.URL.Path):
		returnHomePage(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404"))
	}
}


func returnHomePage (w http.ResponseWriter, r *http.Request){
	setupForJson(&w)

	fmt.Printf("%s", http.MethodGet)

	x := Test{Post: "This is an Echo Post.", Minutes: 9}
	
	json.NewEncoder(w).Encode(x)
}