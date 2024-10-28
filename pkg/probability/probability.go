package probability

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func CalcProb[T comparable](event mapset.Set[T], space mapset.Set[T]) Fraction {
	// The probability of an event, given a sample space:
	// the number of favorable cases divided by the number of all the cases possible."
	// P(E) = n(E) / n(S)
	// where n(E) is the number of favorable cases and n(S) is the number of all the cases possible.
	favourableCases := event.Intersect(space)
	return Fraction{
		Numerator:   favourableCases.Cardinality(),
		Denominator: space.Cardinality(),
	}
}

func Run() {
	event := mapset.NewSet[int]()
	space := mapset.NewSet[int]()
	event.Add(1)
	space.Add(1)
	space.Add(2)
	prob := CalcProb(event, space)
	println("The probability is", float64(prob.Numerator)/float64(prob.Denominator))
}
