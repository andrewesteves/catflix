package main

import (
	"fmt"
	"net/http"

	"github.com/andrewesteves/catflix/internal/handlers"
	"github.com/andrewesteves/catflix/internal/middlewares"
	"github.com/andrewesteves/catflix/internal/repositories"
	"github.com/andrewesteves/catflix/internal/validations"
)

func main() {

	movieStreamCSV := repositories.NewMovieStreamCSV()
	moviesHandler := handlers.NewMovies(movieStreamCSV)

	http.HandleFunc("/", handlers.GetHome)
	http.HandleFunc("/movies", middlewares.ServerSentEvent(moviesHandler.GetMovies))

	fmt.Println("API running")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		validations.PanicOnError(err)
	}
}
