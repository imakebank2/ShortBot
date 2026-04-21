package utils

import (
	"fmt"
	"math/rand/v2"
)

func GetRandomElement[T any](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("Slice is empty.")
	}

	return s[rand.IntN(len(s))], nil
}
