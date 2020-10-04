package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Article is
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Articles ...
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title1", Desc: "Test Desc1", Content: "Test content1"},
		Article{Title: "Test Title2", Desc: "Test Desc2", Content: "Test content2"},
	}

	fmt.Println("Endpont hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
