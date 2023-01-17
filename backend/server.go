package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		(w).Header().Set("Access-Control-Allow-Origin", "*")

		fmt.Fprint(w, "Uh hi...?")
	})

	fmt.Print("Starting server at port 3000\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}