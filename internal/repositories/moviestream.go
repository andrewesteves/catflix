package repositories

import (
	"context"

	"github.com/andrewesteves/catflix/internal/entities"
)

type MovieReader interface {
	Read(ctx context.Context) <-chan entities.Movie
}
