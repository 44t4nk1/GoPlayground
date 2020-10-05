package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Article is
type Article struct {
	Title   string  `json:"title"`
	Desc    string  `json:"desc"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}

//Author ...
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Articles ...
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := Articles{
		Article{Title: "Test Title1", Desc: "Test Desc1", Content: "Test content1", Author: &Author{Firstname: "John", Lastname: "Doe"}},
		Article{Title: "Test Title2", Desc: "Test Desc2", Content: "Test content2", Author: &Author{Firstname: "John", Lastname: "Cena"}},
	}

	fmt.Println("Endpont hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func allPostArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func oneArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := Articles{
		Article{Title: "TestTitle1", Desc: "Test Desc1", Content: "Test content1", Author: &Author{Firstname: "John", Lastname: "Doe"}},
		Article{Title: "TestTitle2", Desc: "Test Desc2", Content: "Test content2", Author: &Author{Firstname: "John", Lastname: "Cena"}},
	}
	params := mux.Vars(r)
	for _, item := range articles {
		if item.Title == params["title"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{title}", oneArticle).Methods("GET")
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
	handleRequest()
	// playground()
	// fmt.Println(concat("lol ", 3))
}
