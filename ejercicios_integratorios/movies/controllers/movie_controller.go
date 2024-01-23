package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/ejercicios_integratorios/movies/models"
	"github.com/labora/labora-golang/ejercicios_integratorios/movies/services"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie models.Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovieId, err := services.CreateMovie(newMovie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while creating the movie."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Movie created correcly with ID: %d\n", newMovieId)
	w.Write([]byte(response))
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(movieID)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Movie with ID %d deleted successfully.", id)
	w.Write([]byte(response))
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(movieID)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var updatedMovie models.Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	fmt.Println(updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateMovie(id, updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Movie with ID %d updated successfully.", id)
	w.Write([]byte(response))
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movieName := r.URL.Query().Get("name")
	if movieName != "" {
		GetMovies(w, r)
		return
	}
	movies, err := services.GetAllMovies()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error fetching the movies."))
		return
	}

	jsonMovies, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error converting movies to JSON."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := jsonMovies
	w.Write([]byte(response))
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	movieName := r.URL.Query().Get("name")
	if movieName != "" {
		movies, err := services.GetMoviesByName(movieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Devuelve la lista de películas filtradas por nombre
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if len(movies) <= 0 {
			response := fmt.Sprintf("There's no movie with the name=%s.", movieName)
			w.Write([]byte(response))
			return
		}
		json.NewEncoder(w).Encode(movies)
	} else {
		// Si no hay un parámetro "nombre", obtén detalles de una película específica por ID
		movieID := mux.Vars(r)["id"]
		id, err := strconv.Atoi(movieID)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		movie, err := services.GetMovieByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Devuelve los detalles de la película específica por ID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if movie == nil {
			response := fmt.Sprintf("There's no movie with ID=%s.", movieID)
			w.Write([]byte(response))
			return
		}
		json.NewEncoder(w).Encode(movie)
	}
}
