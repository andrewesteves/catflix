package transformers

import (
	"context"
	"strings"
	"time"

	"github.com/andrewesteves/catflix/internal/entities"
)

func GroupUppercase(ctx context.Context, movies <-chan entities.Movie) <-chan entities.Movie {
	stream := make(chan entities.Movie)

	go func() {
		defer close(stream)

		for movie := range WithDone(ctx, movies) {
			movie.Group = strings.ToUpper(movie.Group)
			stream <- movie
		}
	}()

	return stream
}

func GetAge(ctx context.Context, movies <-chan entities.Movie) <-chan entities.Movie {
	stream := make(chan entities.Movie)

	go func() {
		defer close(stream)
		currentYear := time.Now().Year()

		for movie := range WithDone(ctx, movies) {
			movie.Age = currentYear - movie.Year
			stream <- movie
		}
	}()

	return stream
}
