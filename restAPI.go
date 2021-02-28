package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json: "Desc"`
	Content string `json: "Content"`
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {
	articles := Article{
		Title:   "test",
		Desc:    "test de",
		Content: "test contet",
	}

	fmt.Println("end point hit")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.newRouter().StrinctSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/test", allArticle)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequests()
}
