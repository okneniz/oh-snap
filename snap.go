package ohsnap

import (
	"testing"
	"time"
)

// Arbitrary is an interface for generating random values and shrinking them.
type Arbitrary[T any] interface {
	Generate() T
	Shrink(T) []T
}

// Property is a function that takes a value of type T and returns a boolean.
type Property[T any] func(T) bool

// Config holds configuration for running property-based tests.
type Config struct {
	Iterations int
	Seed       int64
}

// DefaultConfig returns a default configuration.
func DefaultConfig() Config {
	return Config{
		Iterations: 1000,
		Seed:       0,
	}
}

// Check runs property-based tests with random values and shrinking.
func Check[T any](t *testing.T, arb Arbitrary[T], prop Property[T], cfg Config) {
	if cfg.Seed == 0 {
		cfg.Seed = time.Now().UnixNano()
	}

	for i := 0; i < cfg.Iterations; i++ {
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
