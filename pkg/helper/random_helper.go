package helper

import "math/rand"

// Sample returns a random sample of n elements from the provided slice.
// If n is greater than the length of the slice, the entire slice is returned.
//
// Type Parameters:
//   T: The type of elements in the slice.
//
// Parameters:
//   slice: The input slice from which to sample elements.
//   n: The number of elements to sample from the slice.
//
// Returns:
//   A slice containing n randomly sampled elements from the input slice.
func Sample[T any](slice []T, n int) []T {
	if n > len(slice) {
		return slice
	}

	random := rand.New(rand.NewSource((42)))
	indices := random.Perm(len(slice))

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = slice[indices[i]]
	}

	return result
}
