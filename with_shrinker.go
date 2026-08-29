package ohsnap

import (
	"iter"

	"github.com/okneniz/oh-snap/shrink"
)

// WithShrinker replaces the shrinking strategy of an arbitrary.
// Generate is passed through unchanged.
func WithShrinker[T any](arb Arbitrary[T], s shrink.Shrinker[T]) Arbitrary[T] {
	return &reshrinked[T]{
		arbitrary: arb,
		shrinker:  s,
	}
}

type reshrinked[T any] struct {
	arbitrary Arbitrary[T]
	shrinker  shrink.Shrinker[T]
}

func (r *reshrinked[T]) Generate() iter.Seq[T] {
	return r.arbitrary.Generate()
}

func (r *reshrinked[T]) Shrink(value T) iter.Seq[T] {
	return r.shrinker(value)
}
