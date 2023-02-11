package validations

import "log"

func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
