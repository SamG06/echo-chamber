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
var getPost = regexp.MustCompile(`^/get-post$`)

func getHandler(w http.ResponseWriter, r *http.Request){
	switch {
	case index.MatchString(r.URL.Path):
		returnHomePage(w, r)
	case getPost.MatchString(r.URL.Path):
		getEchoPost(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404"))
	}
}


func returnHomePage (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the home route of the echo post API"))	
}

func getEchoPost (w http.ResponseWriter, r *http.Request){
	setupForJson(&w)

	x := Test{Post: "This is a real Echo Post", Minutes: 8}

	json.NewEncoder(w).Encode(x)
}

