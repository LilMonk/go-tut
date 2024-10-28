package test

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/LilMonk/go-tut/pkg/probability"
	"testing"
)

func TestCalcProb(t *testing.T) {
	tests := []struct {
		name     string
		event    mapset.Set[int]
		space    mapset.Set[int]
		expected probability.Fraction
	}{
		{
			name:     "Single favorable case",
			event:    mapset.NewSet(1),
			space:    mapset.NewSet(1, 2),
			expected: probability.Fraction{Numerator: 1, Denominator: 2},
		},
		{
			name:     "No favorable cases",
			event:    mapset.NewSet(3),
			space:    mapset.NewSet(1, 2),
			expected: probability.Fraction{Numerator: 0, Denominator: 2},
		},
		{
			name:     "All favorable cases",
			event:    mapset.NewSet(1, 2),
			space:    mapset.NewSet(1, 2),
			expected: probability.Fraction{Numerator: 2, Denominator: 2},
		},
		{
			name:     "Empty event set",
			event:    mapset.NewSet[int](),
			space:    mapset.NewSet(1, 2),
			expected: probability.Fraction{Numerator: 0, Denominator: 2},
		},
		{
			name:     "Empty space set",
			event:    mapset.NewSet(1),
			space:    mapset.NewSet[int](),
			expected: probability.Fraction{Numerator: 0, Denominator: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := probability.CalcProb(tt.event, tt.space)
			if result != tt.expected {
				t.Errorf("calcProb() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
