package helper

import (
	"fmt"
)

type Set[T comparable] map[T]struct{}

// Adds an element to the set.
func (s Set[T]) Add(e T) {
	s[e] = struct{}{}

	// Why Use struct{}{}:
	// An empty struct consumes zero bytes of memory,
	// making it an efficient way to store elements in a map
	// when you only care about the presence of the key and not its value.
	// It's an ideal pattern when implementing a set,
	// where the goal is to track unique elements.
}

// Removes an element from the set.
func (s Set[T]) Remove(e T) {
	delete(s, e)
}

// Checks if an element is in the set.
func (s Set[T]) Contains(e T) bool {
	_, found := s[e]
	return found
}

// Returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// Returns a slice of all elements in the set.
func (s Set[T]) Elements() []T {
	elements := make([]T, 0, len(s))
	for e := range s {
		elements = append(elements, e)
	}
	return elements
}

// Prints the set.
func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Elements())
}
