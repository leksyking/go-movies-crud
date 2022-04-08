package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: ""director`
}

type Director struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	//	fmt.Fprintf(w, "movies ar: %v", movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "23415", Title: "The Blacklist", Director: &Director{FirstName: "Felix", LastName: "Ogundipe"},
	})
	movies = append(movies, Movie{
		ID: "2", Isbn: "45272", Title: "Peaky Blinders", Director: &Director{FirstName: "John", LastName: "Doe"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Sever started on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
