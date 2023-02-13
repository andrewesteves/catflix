package transformers

import "context"

func WithDone[T any](ctx context.Context, in <-chan T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				select {
				case stream <- value:
				case <-ctx.Done():
				}
			}
		}
	}()

	return stream
}
