package repositories

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/andrewesteves/catflix/internal/entities"
	"github.com/andrewesteves/catflix/internal/validations"
)

type MovieStreamCSV struct{}

func NewMovieStreamCSV() MovieReader {
	return &MovieStreamCSV{}
}

func (m *MovieStreamCSV) Read(ctx context.Context) <-chan entities.Movie {
	movies := make(chan entities.Movie)

	go func() {
		defer close(movies)

		file, err := os.Open("netflix.csv")
		validations.PanicOnError(err)
		defer file.Close()

		reader := csv.NewReader(bufio.NewReader(file))
		_, err = reader.Read()
		validations.PanicOnError(err)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				row, err := reader.Read()
				if err != nil {
					if err != io.EOF {
						validations.PanicOnError(err)
					}
					break
				}
				year, err := strconv.Atoi(row[7])
				validations.PanicOnError(err)

				movies <- entities.Movie{
					Title: row[2],
					Group: row[1],
					Year:  year,
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return movies
}
