package transformers

import (
	"context"
	"testing"

	"github.com/andrewesteves/catflix/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestGroupUppercase(t *testing.T) {
	movieStream := make(chan entities.Movie)

	go func() {
		defer close(movieStream)
		for i := 0; i < 3; i++ {
			movieStream <- entities.Movie{Group: "movie"}
		}
	}()

	ctx := context.Background()
	var i int
	for movie := range GroupUppercase(ctx, movieStream) {
		require.Equal(t, "MOVIE", movie.Group)
		i++
	}
	require.Equal(t, 3, i)
}
