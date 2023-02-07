package main

import (
	"echo_chamber/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type EchoPost struct {
	Content string `json:"content"`
	Date int `json:"date"`
}
var db = database.Connection()

func main() {
	database.SetupDatabase()

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
	setupForJson(&w)
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPost:
		postHandler(w, r)
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

var createPost = regexp.MustCompile(`^/create-post$`)

func postHandler(w http.ResponseWriter, r *http.Request){
	switch{
	case createPost.MatchString(r.URL.Path):
		createEchoPost(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404"))
	}
}
	



func returnHomePage (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the home route of the echo post API"))	
}

func getEchoPost (w http.ResponseWriter, r *http.Request){
	x := EchoPost{Content: "This is a real Echo Post", Date: 8}

	json.NewEncoder(w).Encode(x)
}


func executeSQL(sql_ string){
	
}

func createEchoPost (w http.ResponseWriter, r *http.Request){
	var post EchoPost
	
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	sqlMakePost := "INSERT INTO EchoPost (content) VALUES (?)"

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO echo_post (content) VALUES ($1)`)
    
fmt.Print(sqlMakePost)

	if err != nil {
		log.Fatal(err)
    }
    
	defer stmt.Close()

	_, err = stmt.Exec(post.Content)

	if err != nil {
		log.Fatal(err)
    }

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Post: %+v", post)
	json.NewEncoder(w).Encode(post)
}

