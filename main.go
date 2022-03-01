package main

import (
	"encoding/json" //encode data into json
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"Director"` // * defines Director struct
}

type Director struct {
	FName string `json:"FName"`
	LName string `json:"LName"`
}

var movies []Movie //movies slice of type Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	// return the remaining movies
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovieById(w http.ResponseWriter, r *http.Request) {

	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//loop over the movies, range
	for i, item := range movies {
		if item.Id == params["id"] {

			//delete the movie with the i.d that u've sent
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			//add a new movie - the movie that we sent in the body of postman
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "3434", Title: "The Batman", Director: &Director{FName: "Chris", LName: "Nolan"}})
	movies = append(movies, Movie{Id: "2", Isbn: "1212", Title: "The Batman Returns", Director: &Director{FName: "Chris", LName: "Nolan 2"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovieById).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")

	fmt.Printf("Starting the server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
