package test

import (
	"github.com/LilMonk/go-tut/pkg/helper"
	"reflect"
	"testing"
)

type CustomStruct struct {
	ID   int
	Name string
}

// Convert a slice of any type to []interface{}
func toInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func TestSample(t *testing.T) {
	// Integers
	intSlice := []int{1, 2, 3, 4, 5}
	intResult := helper.Sample(intSlice, 3)
	if len(intResult) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(intResult))
	}

	// Floats
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	floatResult := helper.Sample(floatSlice, 4)
	if len(floatResult) != 4 {
		t.Errorf("Expected 4 elements, got %d", len(floatResult))
	}

	// Strings
	stringSlice := []string{"apple", "banana", "cherry", "date"}
	stringResult := helper.Sample(stringSlice, 2)
	if len(stringResult) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(stringResult))
	}

	// Custom Structs
	structSlice := []CustomStruct{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}}
	structResult := helper.Sample(structSlice, 2)
	if len(structResult) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(structResult))
	}

	// Array of Arrays
	arrayOfArrays := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	arrayResult := helper.Sample(arrayOfArrays, 3)
	if len(arrayResult) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(arrayResult))
	}

	// General Element Check: All sampled elements are from original slices
	for _, tc := range []struct {
		name   string
		slice  []interface{}
		result []interface{}
	}{
		{"Integers", toInterfaceSlice(intSlice), toInterfaceSlice(intResult)},
		{"Floats", toInterfaceSlice(floatSlice), toInterfaceSlice(floatResult)},
		{"Strings", toInterfaceSlice(stringSlice), toInterfaceSlice(stringResult)},
		{"Custom Structs", toInterfaceSlice(structSlice), toInterfaceSlice(structResult)},
		{"Array of Arrays", toInterfaceSlice(arrayOfArrays), toInterfaceSlice(arrayResult)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.result {
				found := false
				for _, orig := range tc.slice {
					if reflect.DeepEqual(v, orig) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Element %v not found in original %s slice", v, tc.name)
				}
			}
		})
	}
}
