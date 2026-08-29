package ohsnap

import (
	"iter"
	"testing"
)

// Arbitrary is an interface for generating random values and shrinking them.
// Generate returns a lazy, effectively infinite sequence of random values.
// Shrink returns a lazy sequence of candidates simpler than the input value.
type Arbitrary[T any] interface {
	Generate() iter.Seq[T]
	Shrink(T) iter.Seq[T]
}

// Property is a function that takes a value of type T and returns a boolean.
type Property[T any] func(T) bool

// ShrinkOptions configures the shrinking phase of CheckWith.
type ShrinkOptions struct {
	// Budget limits the total number of property calls made during shrinking.
	// When the budget is exhausted, shrinking stops and returns the best
	// value found so far. Zero or negative means unlimited.
	Budget int
}

// Check runs property-based tests with random values and shrinking.
func Check[T any](t *testing.T, iterations int, arb Arbitrary[T], prop Property[T]) {
	CheckWith(t, iterations, arb, prop, ShrinkOptions{})
}

// CheckWith runs property-based tests with custom shrinking options.
func CheckWith[T any](t *testing.T, iterations int, arb Arbitrary[T], prop Property[T], opts ShrinkOptions) {
	value, shrunk := findSimplestBadCaseWith(iterations, arb, prop, opts)
	if value != nil {
		t.Errorf("Property failed for value: %v (shrunk: %v)", value, shrunk)
		return
	}
}

// findSimplestBadCase find simplest bad case of input value for property func
func findSimplestBadCase[T any](iterations int, arb Arbitrary[T], prop Property[T]) (*T, *T) {
	return findSimplestBadCaseWith(iterations, arb, prop, ShrinkOptions{})
}

func findSimplestBadCaseWith[T any](iterations int, arb Arbitrary[T], prop Property[T], opts ShrinkOptions) (*T, *T) {
	if iterations <= 0 {
		return nil, nil
	}

	i := 0
	for value := range arb.Generate() {
		if i >= iterations {
			break
		}
		i++

		if !prop(value) {
			var shrunk T
			if opts.Budget > 0 {
				budget := opts.Budget
				shrunk = shrinkValueBudgeted(arb, value, prop, &budget)
			} else {
				shrunk = shrinkValue(arb, value, prop)
			}
			return &value, &shrunk
		}
	}

	return nil, nil
}

// shrinkValue attempts to shrink a failing value to its minimal form.
func shrinkValue[T any](arb Arbitrary[T], value T, prop Property[T]) T {
	for smaller := range arb.Shrink(value) {
		if !prop(smaller) {
			return shrinkValue(arb, smaller, prop)
		}
	}

	return value
}

// shrinkValueBudgeted attempts to shrink a failing value with a limited
// number of property calls. When the budget is exhausted it returns the
// best value found so far.
func shrinkValueBudgeted[T any](arb Arbitrary[T], value T, prop Property[T], budget *int) T {
	for smaller := range arb.Shrink(value) {
		if *budget <= 0 {
			return value
		}
		*budget--

		if !prop(smaller) {
			return shrinkValueBudgeted(arb, smaller, prop, budget)
		}
	}

	return value
}
