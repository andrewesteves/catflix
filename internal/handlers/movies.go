package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrewesteves/catflix/internal/repositories"
	"github.com/andrewesteves/catflix/internal/transformers"
	"github.com/andrewesteves/catflix/internal/validations"
)

type Movies struct {
	movieStreamRepository repositories.MovieReader
}

func NewMovies(movieStreamRepository repositories.MovieReader) *Movies {
	return &Movies{
		movieStreamRepository: movieStreamRepository,
	}
}

func (m *Movies) GetMovies(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "no streaming support", http.StatusInternalServerError)
		return
	}
	flusher.Flush()

	for movie := range transformers.GetAge(
		r.Context(), transformers.GroupUppercase(r.Context(), m.movieStreamRepository.Read(r.Context())),
	) {
		m, err := json.Marshal(movie)
		validations.PanicOnError(err)

		fmt.Fprintf(w, "data: %v\n\n", string(m))
		flusher.Flush()
	}

	fmt.Fprintf(w, "data: %v\n\n", "closed")
	flusher.Flush()
}
