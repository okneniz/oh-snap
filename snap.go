package ohsnap

import (
	"testing"
)

// Arbitrary is an interface for generating random values and shrinking them.
type Arbitrary[T any] interface {
	Generate() T
	Shrink(T) []T
}

// Property is a function that takes a value of type T and returns a boolean.
type Property[T any] func(T) bool

// Check runs property-based tests with random values and shrinking.
func Check[T any](t *testing.T, iterations int, arb Arbitrary[T], prop Property[T]) {
	for i := 0; i < iterations; i++ {
		value := arb.Generate()
		if !prop(value) {
			shrunk := shrinkValue(arb, value, prop)
			t.Errorf("Property failed for value: %v (shrunk: %v)", value, shrunk)
			return
		}
	}
}

// shrinkValue attempts to shrink a failing value to its minimal form.
func shrinkValue[T any](arb Arbitrary[T], value T, prop Property[T]) T {
	for _, smaller := range arb.Shrink(value) {
		if !prop(smaller) {
			return shrinkValue(arb, smaller, prop)
		}
	}

	return value
}
