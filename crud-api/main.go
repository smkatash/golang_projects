package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// *** List all movies ***
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// *** Delete movies ***
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// *** Get a particular movie ***
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// *** Create an item ***
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mv Movie
	_ = json.NewDecoder(r.Body).Decode(&mv)
	mv.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, mv)
	json.NewEncoder(w).Encode(mv)
}


// *** Update an item ***
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			var mv Movie
			_ = json.NewDecoder(r.Body).Decode(&mv)
			mv.ID = params["id"]
			movies = append(movies, mv)
			json.NewEncoder(w).Encode(mv)
		}
	}
}

// *** Main ***
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Harry Potter",
		Director: &Director{Firstname: "Kany", Lastname: "Tash"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45673", Title: "Home Alone",
		Director: &Director{Firstname: "Steve", Lastname: "Eve"}})
	r.HandleFunc("/movies", getMovies).Methods("GET") 
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") 
	r.HandleFunc("/movies", createMovie).Methods("POST") 
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}