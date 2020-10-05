package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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

func allPostArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Titlea", Desc: "Test Desca", Content: "Test contenta"},
		Article{Title: "Test Titleb", Desc: "Test Descb", Content: "Test contentb"},
	}

	fmt.Println("Endpont hit: all post articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func playground() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}

func concat(a string, b int) string {
	number := strconv.Itoa(b)
	final := a + number
	return final
}

func main() {
	// handleRequest()
	// playground()
	fmt.Fprintln(os.Stdout, concat("lol ", 3))
}
